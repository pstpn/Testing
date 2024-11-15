package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"bee/config"
	"bee/internal/service"
	"bee/internal/service/dto"
	"bee/pkg/logger"
)

type BasicMiddleware struct {
	l           logger.Interface
	authService service.AuthService
	adminConfig *config.AdminConfig
}

func NewBasicMiddleware(l logger.Interface, authService service.AuthService, adminConfig *config.AdminConfig) *BasicMiddleware {
	return &BasicMiddleware{
		l:           l,
		authService: authService,
		adminConfig: adminConfig,
	}
}

func (a *BasicMiddleware) BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, password, hasAuth := c.Request.BasicAuth()
		if hasAuth {
			if email == a.adminConfig.Login && password == a.adminConfig.Password {
				c.Set("role", "admin")
				c.Next()
				return
			}

			login, err := a.authService.LoginUser(c.Request.Context(), &dto.LoginUserRequest{
				Email:    email,
				Password: password,
			})
			if err != nil {
				a.l.Errorf("incorrect request body: %s", err.Error())
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request body"})
				return
			}
			if login {
				c.Set("role", "user")
				c.Next()
				return
			}
		}

		a.l.Errorf("incorrect login or password")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Incorrect login or password"})
	}
}

func RoleMiddleware(requiredRoles map[string]struct{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if exists {
			if _, correct := requiredRoles[role.(string)]; correct {
				c.Next()
				return
			}
		}
		c.AbortWithStatus(http.StatusForbidden)
	}
}
