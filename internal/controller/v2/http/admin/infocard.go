package admin

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	httputils "course/internal/controller/v2/http/utils"
	"course/internal/model"
	"course/internal/service"
	"course/internal/service/dto"
	"course/pkg/logger"
	_ "course/pkg/server/http"
	"course/pkg/storage/postgres"
)

type InfoCardController struct {
	l                 logger.Interface
	infoCardService   service.InfoCardService
	documentService   service.DocumentService
	fieldService      service.FieldService
	checkpointService service.CheckpointService
	authService       service.AuthService
}

func NewInfoCardController(
	l logger.Interface,
	infoCardService service.InfoCardService,
	documentService service.DocumentService,
	fieldService service.FieldService,
	checkpointService service.CheckpointService,
	authService service.AuthService,
) *InfoCardController {
	return &InfoCardController{
		l:                 l,
		infoCardService:   infoCardService,
		documentService:   documentService,
		fieldService:      fieldService,
		checkpointService: checkpointService,
		authService:       authService,
	}
}

type listFullInfoCardsResponse struct {
	InfoCards []*model.FullInfoCard `json:"infoCards"`
}

// ListFullInfoCards godoc
//
//	@Summary		Получение коллекции информационных карточек
//	@Description	Метод для получения коллекции информационных карточек
//	@Tags			admin
//	@Param			pattern	query		string							false	"Значение для фильтрации"
//	@Param			field	query		string							false	"Поле для фильтрации и сортировки"
//	@Param			sort	query		string							false	"Направление сортировки"
//	@Success		200		{object}	listFullInfoCardsResponse		"Информация о карточках успешно получена"
//	@Failure		400		{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		401		{object}	http.StatusUnauthorized			"Авторизация неуспешна"
//	@Failure		500		{object}	http.StatusInternalServerError	"Внутренняя ошибка получения карточек пользователей"
//	@Security		BearerAuth
//	@Router			/infocards [get]
func (i *InfoCardController) ListFullInfoCards(c *gin.Context) {
	_, err := httputils.VerifyAccessToken(c, i.l, i.authService)
	if err != nil {
		return
	}

	fullInfoCards, err := i.infoCardService.ListInfoCards(c.Request.Context(), &dto.ListInfoCardsRequest{
		Pagination: &postgres.Pagination{
			PageNumber: -1,
			PageSize:   -1,
			Filter: postgres.FilterOptions{
				Pattern: c.Query("pattern"),
				Column:  c.Query("field"),
			},
			Sort: postgres.SortOptions{
				Direction: postgres.SortDirectionFromString(c.Query("sort")),
				Columns:   []string{c.Query("field")},
			},
		},
	})
	if err != nil {
		i.l.Errorf("failed to list fullInfoCards: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list info cards"})
		return
	}

	c.JSON(http.StatusOK, listFullInfoCardsResponse{InfoCards: fullInfoCards})
}

type getFullInfoCardResponse struct {
	Document *model.FullDocument   `json:"document"`
	Passages []*model.ShortPassage `json:"passages"`
}

// GetFullInfoCard godoc
//
//	@Summary		Получение элемента коллекции информационных карточек
//	@Description	Метод для получения элемента коллекции информационных карточек
//	@Tags			admin
//	@Param			id	path		string							true	"Идентификатор информационной карточки"
//	@Success		200	{object}	getFullInfoCardResponse			"Информация о карточке успешно получена"
//	@Failure		400	{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized			"Авторизация неуспешна"
//	@Failure		404	{object}	http.StatusNotFound				"Карточка не найдена"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка получения карточки пользователя"
//	@Security		BearerAuth
//	@Router			/infocards/{id} [get]
func (i *InfoCardController) GetFullInfoCard(c *gin.Context) {
	_, err := httputils.VerifyAccessToken(c, i.l, i.authService)
	if err != nil {
		return
	}

	infoCardID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		i.l.Errorf("failed to parse infoCard ID from query args: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to get info card ID"})
		return
	}

	document, err := i.documentService.GetDocumentByInfoCard(c.Request.Context(), &dto.GetDocumentByInfoCardIDRequest{
		InfoCardID: infoCardID,
	})
	if err != nil {
		i.l.Errorf("failed to get document by infoCard ID: %s", err.Error())

		status := http.StatusInternalServerError
		if errors.Is(err, pgx.ErrNoRows) {
			status = http.StatusNotFound
		}
		c.AbortWithStatusJSON(status, gin.H{"error": "Failed to get info card document"})
		return
	}

	documentFields, err := i.fieldService.ListDocumentFields(c.Request.Context(), &dto.ListDocumentFieldsRequest{
		DocumentID: document.ID.Int(),
	})
	if err != nil {
		i.l.Errorf("failed to list document fields: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list document fields"})
		return
	}

	passages, err := i.checkpointService.ListPassages(c.Request.Context(), &dto.ListPassagesRequest{DocumentID: document.ID.Int()})
	if err != nil {
		i.l.Errorf("failed to list passages: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list passages"})
		return
	}

	c.JSON(http.StatusOK, getFullInfoCardResponse{
		Document: &model.FullDocument{
			Data: &model.DataDocument{
				DocumentType: document.Type.String(),
				SerialNumber: document.SerialNumber,
			},
			Fields: model.ModelToKeyValue(documentFields),
		},
		Passages: model.ModelToShortPassages(passages),
	})
}

// ConfirmEmployeeInfoCard godoc
//
//	@Summary		Подтверждение информационной карточки сотрудника
//	@Description	Метод для подтверждения информационной карточки сотрудника
//	@Tags			admin
//	@Param			id	path		string					true	"Идентификатор информационной карточки"
//	@Success		200	{string}	string					"Информация о карточке успешно подтверждена"
//	@Failure		400	{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized	"Авторизация неуспешна"
//	@Security		BearerAuth
//	@Router			/infocards/{id} [patch]
func (i *InfoCardController) ConfirmEmployeeInfoCard(c *gin.Context) {
	_, err := httputils.VerifyAccessToken(c, i.l, i.authService)
	if err != nil {
		return
	}

	infoCardID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		i.l.Errorf("failed to parse infoCard ID from query args: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to get info card ID"})
		return
	}

	err = i.infoCardService.ValidateInfoCard(c.Request.Context(), &dto.ValidateInfoCardRequest{
		InfoCardID:  infoCardID,
		IsConfirmed: true,
	})
	if err != nil {
		i.l.Errorf("failed to validate infoCard: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to validate info card"})
		return
	}

	c.JSON(http.StatusOK, "OK")
}
