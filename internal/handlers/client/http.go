package clienthandler

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type HTTPHandler struct {
	clientController ports.ClientController
}

func NewHTTPHandler(clientController ports.ClientController) *HTTPHandler {
	return &HTTPHandler{
		clientController: clientController,
	}
}

func (hdl *HTTPHandler) Get(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	client, err := hdl.clientController.Get(id)
	fmt.Println(err)
	if err != nil {
		c.JSON(500, "string")
		return
	}

	return c.JSON(200, client)
}

func (hdl *HTTPHandler) Create(c echo.Context) (err error) {
	body := domain.Client{}
	c.Bind(&body)

	client, err := hdl.clientController.Create(body.Lastname, body.Firstname, body.Email)
	if err != nil {
		c.JSON(500, "error")
		return
	}

	return c.JSON(200, client)
}
