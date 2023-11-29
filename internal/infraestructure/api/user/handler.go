package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maoudev/todo/internal/pkg/domain"
	"github.com/maoudev/todo/internal/pkg/ports"
)

type userHandler struct {
	userService ports.UserService
}

func newHandler(userService ports.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (u *userHandler) CreateUser(c *gin.Context) {
	user := &domain.User{}
	if err := c.BindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := u.userService.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (u *userHandler) Login(c *gin.Context) {
	credentias := &domain.DefaultCredentials{}

	if err := c.BindJSON(credentias); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	jwtToken, err := u.userService.Login(credentias)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
	})
}
