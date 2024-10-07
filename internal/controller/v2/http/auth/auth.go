package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"course/internal/model"
	"course/internal/service"
	"course/internal/service/dto"
	"course/pkg/logger"
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
	PhoneNumber string `json:"phoneNumber" example:"+71234567890"`
	Name        string `json:"name"  example:"Степа"`
	Surname     string `json:"surname"  example:"Степик"`
	CompanyID   int64  `json:"companyID"  example:"1"`
	Post        string `json:"post"  example:"Сотрудник"`
	Password    string `json:"password"  example:"123"`
	DateOfBirth string `json:"dateOfBirth"  example:"31.03.2004"`
}

type registerResponse struct {
	AccessToken  string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2IiwiaWF0IjoxNzI1NzE4NTc2fQ.RdcJHE8TULJKW-mVyn-0OBL_O_kmISrFNuK6E8FeRSs"`
	RefreshToken string `json:"refreshToken" example:"c8edc98acee4d6243add4e59f8fd46d6a7f150831d74f0feb3a10144cbe0c032"`
	IsAdmin      bool   `json:"isAdmin" example:"false" format:"bool"`
}

// Register godoc
//
//	@Summary		Регистрация пользователя
//	@Description	Метод для регистрации пользователя
//	@Tags			auth
//	@Param			registerRequest	body		registerRequest					true	"Регистрация пользователя"
//	@Success		200				{object}	registerResponse				"Пользователь успешно зарегистрирован"
//	@Failure		400				{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		500				{object}	http.StatusInternalServerError	"Внутренняя ошибка регистрации пользователя"
//	@Router			/register [post]
func (a *Controller) Register(c *gin.Context) {
	var req registerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		a.l.Errorf("incorrect request body: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request body"})
		return
	}

	birthDate, err := time.Parse("02.01.2006", req.DateOfBirth)
	if err != nil {
		a.l.Errorf("incorrect request body date of birth: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request body date of birth"})
		return
	}

	tokens, err := a.authService.RegisterEmployee(c.Request.Context(), &dto.RegisterEmployeeRequest{
		PhoneNumber:    req.PhoneNumber,
		FullName:       req.Name + " " + req.Surname,
		CompanyID:      req.CompanyID,
		Post:           model.ToPostTypeFromString(req.Post).Int(),
		Password:       req.Password,
		RefreshToken:   "",
		TokenExpiredAt: nil,
		DateOfBirth:    &birthDate,
	})
	if err != nil {
		err = fmt.Errorf("can`t register employee: %w", err)
		a.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to register employee"})
		return
	}

	c.JSON(http.StatusOK, registerResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		IsAdmin:      tokens.IsAdmin,
	})
}

type loginRequest struct {
	PhoneNumber string `json:"phoneNumber" example:"+71234567890"`
	Password    string `json:"password" example:"123"`
}

type loginResponse struct {
	AccessToken  string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2IiwiaWF0IjoxNzI1NzE4NTc2fQ.RdcJHE8TULJKW-mVyn-0OBL_O_kmISrFNuK6E8FeRSs"`
	RefreshToken string `json:"refreshToken" example:"c8edc98acee4d6243add4e59f8fd46d6a7f150831d74f0feb3a10144cbe0c032"`
	IsAdmin      bool   `json:"isAdmin" example:"false" format:"bool"`
}

// Login godoc
//
//	@Summary		Вход в аккаунт пользователя
//	@Description	Метод для входа в аккаунт пользователя
//	@Tags			auth
//	@Param			loginRequest	body		loginRequest			true	"Вход пользователя"
//	@Success		200				{object}	loginResponse			"Пользователь успешно авторизовался"
//	@Failure		400				{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Failure		401				{object}	http.StatusUnauthorized	"Вход неуспешен"
//	@Router			/login [post]
func (a *Controller) Login(c *gin.Context) {
	var req loginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		a.l.Errorf("incorrect request body: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request body"})
		return
	}

	tokens, err := a.authService.LoginEmployee(c.Request.Context(), &dto.LoginEmployeeRequest{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})
	if err != nil {
		a.l.Errorf("can`t login employee: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Can`t login employee"})
		return
	}

	c.JSON(http.StatusOK, loginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		IsAdmin:      tokens.IsAdmin,
	})
}

type refreshTokensRequest struct {
	AccessToken  string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2IiwiaWF0IjoxNzI1NzE4NTc2fQ.RdcJHE8TULJKW-mVyn-0OBL_O_kmISrFNuK6E8FeRSs"`
	RefreshToken string `json:"refreshToken" example:"c8edc98acee4d6243add4e59f8fd46d6a7f150831d74f0feb3a10144cbe0c032"`
}

type refreshTokensResponse struct {
	AccessToken  string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2IiwiaWF0IjoxNzI1NzE4NTc2fQ.RdcJHE8TULJKW-mVyn-0OBL_O_kmISrFNuK6E8FeRSs"`
	RefreshToken string `json:"refreshToken" example:"c8edc98acee4d6243add4e59f8fd46d6a7f150831d74f0feb3a10144cbe0c032"`
	IsAdmin      bool   `json:"isAdmin" example:"false" format:"bool"`
}

// RefreshTokens godoc
//
//	@Summary		Обновление токенов доступа
//	@Description	Метод для обновления токенов доступа пользователя
//	@Tags			auth
//	@Param			refreshTokensRequest	body		refreshTokensRequest	true	"Обновление токенов доступа"
//	@Success		200						{object}	refreshTokensResponse	"Токены успешно обновлены"
//	@Failure		400						{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Failure		401						{object}	http.StatusUnauthorized	"Вход неуспешен"
//	@Router			/refresh [post]
func (a *Controller) RefreshTokens(c *gin.Context) {
	var req refreshTokensRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		a.l.Errorf("incorrect request body: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request body"})
		return
	}

	payload, err := a.authService.VerifyEmployeeAccessToken(c.Request.Context(), &dto.VerifyEmployeeAccessTokenRequest{AccessToken: req.AccessToken})
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) && !errors.Is(err, jwt.ErrTokenNotValidYet) {
		a.l.Errorf("failed to verify token: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}
	infoCardID, err := payload.GetInfoCardID()
	if err != nil {
		a.l.Errorf("failed to parse infoCard id from payload: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	tokens, err := a.authService.RefreshTokens(c.Request.Context(), &dto.RefreshEmployeeTokensRequest{
		InfoCardID:   infoCardID,
		RefreshToken: req.RefreshToken,
	})
	if errors.Is(err, jwt.ErrTokenExpired) {
		a.l.Warnf("expired refresh token: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Expired refresh token"})
		return
	}
	if err != nil {
		a.l.Errorf("refresh tokens for employee: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can`t refresh tokens"})
		return
	}

	c.JSON(http.StatusOK, refreshTokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		IsAdmin:      tokens.IsAdmin,
	})
}
