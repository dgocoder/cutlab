package userhandler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/testd/cutlab/internal/core/ports"
)

type HTTPHandler struct {
	userController ports.UserController
}

func NewHTTPHandler(UserController ports.UserController) *HTTPHandler {
	return &HTTPHandler{
		userController: UserController,
	}
}

type UserView struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (hdl *HTTPHandler) Get(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	service, err := hdl.userController.Get(id)
	if err != nil {
		c.JSON(500, err)
		return
	}

	return c.JSON(200, service)
}

func (hdl *HTTPHandler) Create(c echo.Context) (err error) {
	body := UserView{}
	c.Bind(&body)

	user, err := hdl.userController.Create(body.Name, body.Email, body.Password, nil)
	if err != nil {
		c.JSON(500, "error")
		return
	}

	return c.JSON(200, user)
}
