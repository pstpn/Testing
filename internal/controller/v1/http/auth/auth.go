package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"bee/internal/service"
	"bee/internal/service/dto"
	"bee/pkg/logger"
)

type Controller struct {
	l           logger.Interface
	authService service.AuthService
}

func NewAuthController(l logger.Interface, authService service.AuthService) *Controller {
	return &Controller{
		l:           l,
		authService: authService,
	}
}

type registerRequest struct {
	Email    string `json:"email" example:"beebee@bee.bee"`
	Password string `json:"password" example:"123"`
}

// Register godoc
//
//	@Summary		Регистрация пользователя
//	@Description	Метод для регистрации пользователя
//	@Tags			auth
//	@Param			registerRequest	body		registerRequest			true	"Регистрация пользователя"
//	@Success		201				{string}	string					"Пользователь успешно зарегистрирован"
//	@Failure		400				{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Router			/register [post]
func (a *Controller) Register(c *gin.Context) {
	var req registerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		a.l.Errorf("incorrect request body: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request body"})
		return
	}

	err := a.authService.RegisterUser(c.Request.Context(), &dto.RegisterUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		err = fmt.Errorf("can`t register user: %w", err)
		a.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to register employee"})
		return
	}

	c.JSON(http.StatusCreated, "OK")
}
