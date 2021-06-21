package servicehandler

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/testd/cutlab/internal/core/ports"
)

type HTTPHandler struct {
	serviceController ports.ServiceController
}

func NewHTTPHandler(serviceController ports.ServiceController) *HTTPHandler {
	return &HTTPHandler{
		serviceController: serviceController,
	}
}

func (hdl *HTTPHandler) Get(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	service, err := hdl.serviceController.Get(id)
	fmt.Println(err)
	if err != nil {
		c.JSON(500, "string")
		return
	}

	return c.JSON(200, service)
}

func (hdl *HTTPHandler) Create(c echo.Context) (err error) {
	body := ports.CreateServiceView{}
	c.Bind(&body)

	service, err := hdl.serviceController.Create(body)
	if err != nil {
		c.JSON(500, "error")
		return
	}

	return c.JSON(200, service)
}
