package user

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

type getProfileResponse struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	City    string `json:"city"`
	Email   string `json:"email"`
}

// GetProfile godoc
//
//	@Summary		Просмотр профиля пользователя
//	@Description	Метод для просмотра профиля пользователя
//	@Tags			user
//	@Param			email	path		string					true	"Email пользователя"
//	@Success		200		{object}	getProfileResponse		"Профиль успешно получен"
//	@Failure		400		{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Failure		401		{object}	http.StatusUnauthorized	"Авторизация неуспешна"
//	@Failure		403		{object}	http.StatusForbidden	"Отсутствуют права доступа"
//	@Security		BasicAuth
//	@Router			/profiles/{email} [get]
func (p *ProfileController) GetProfile(c *gin.Context) {
	profile, err := p.profileService.GetProfile(c.Request.Context(), &dto.GetProfileRequest{
		Email: c.Param("email"),
	})
	if err != nil {
		p.l.Errorf("failed to get profile: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Failed to get profile"})
		return
	}

	c.JSON(http.StatusOK, getProfileResponse{
		Name:    profile.Name,
		Surname: profile.Surname,
		City:    profile.City,
		Email:   profile.Email,
	})
}
