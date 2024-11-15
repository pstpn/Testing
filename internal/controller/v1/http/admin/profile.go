package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"bee/internal/service"
	"bee/internal/service/dto"
	"bee/pkg/logger"
)

type ProfileController struct {
	l              logger.Interface
	profileService service.ProfileService
}

func NewProfileController(l logger.Interface, profileService service.ProfileService) *ProfileController {
	return &ProfileController{
		l:              l,
		profileService: profileService,
	}
}

type createProfileRequest struct {
	Name    string `json:"name" example:"Bee"`
	Surname string `json:"surname" example:"Beevich"`
	City    string `json:"city" example:"Beeland"`
	Email   string `json:"email" example:"beebee@bee.bee"`
}

// CreateProfile godoc
//
//	@Summary		Создание профиля пользователя
//	@Description	Метод для создания профиля пользователя
//	@Tags			admin
//	@Param			createProfileRequest	body		createProfileRequest	true	"Создание профиля пользователя"
//	@Success		201						{string}	string					"Профиль успешно создан"
//	@Failure		400						{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Failure		401						{object}	http.StatusUnauthorized	"Авторизация неуспешна"
//	@Failure		403						{object}	http.StatusForbidden	"Отсутствуют права доступа"
//	@Security		BasicAuth
//	@Router			/profiles [post]
func (p *ProfileController) CreateProfile(c *gin.Context) {
	var req createProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		p.l.Errorf("incorrect request body: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request body"})
		return
	}

	err := p.profileService.CreateProfile(c.Request.Context(), &dto.CreateProfileRequest{
		Name:    req.Name,
		Surname: req.Surname,
		City:    req.City,
		Email:   req.Email,
	})
	if err != nil {
		p.l.Errorf("failed to create profile: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to create profile"})
		return
	}

	c.JSON(http.StatusCreated, "OK")
}

type updateProfileRequest struct {
	Name    string `json:"name" example:"Bee"`
	Surname string `json:"surname" example:"Beevich"`
	City    string `json:"city" example:"NewBeeland"`
}

// UpdateProfile godoc
//
//	@Summary		Изменение профиля пользователя
//	@Description	Метод для изменения профиля пользователя
//	@Tags			admin
//	@Param			email					path		string					true	"Email пользователя"
//	@Param			updateProfileRequest	body		updateProfileRequest	true	"Изменение профиля пользователя"
//	@Success		200						{string}	string					"Профиль успешно изменен"
//	@Failure		400						{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Failure		401						{object}	http.StatusUnauthorized	"Авторизация неуспешна"
//	@Failure		403						{object}	http.StatusForbidden	"Отсутствуют права доступа"
//	@Security		BasicAuth
//	@Router			/profiles/{email} [put]
func (p *ProfileController) UpdateProfile(c *gin.Context) {
	var req updateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		p.l.Errorf("incorrect request body: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request body"})
		return
	}

	err := p.profileService.UpdateProfile(c.Request.Context(), &dto.UpdateProfileRequest{
		Name:    req.Name,
		Surname: req.Surname,
		City:    req.City,
		Email:   c.Param("email"),
	})
	if err != nil {
		p.l.Errorf("failed to update profile: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, "OK")
}

// DeleteProfile godoc
//
//	@Summary		Удаление профиля пользователя
//	@Description	Метод для удаления профиля пользователя
//	@Tags			admin
//	@Param			email	path		string					true	"Email пользователя"
//	@Success		200		{string}	string					"Профиль успешно удален"
//	@Failure		400		{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Failure		401		{object}	http.StatusUnauthorized	"Авторизация неуспешна"
//	@Failure		403		{object}	http.StatusForbidden	"Отсутствуют права доступа"
//	@Security		BasicAuth
//	@Router			/profiles/{email} [delete]
func (p *ProfileController) DeleteProfile(c *gin.Context) {
	err := p.profileService.DeleteProfile(c.Request.Context(), &dto.DeleteProfileRequest{
		Email: c.Param("email"),
	})
	if err != nil {
		p.l.Errorf("failed to update profile: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, "OK")
}
