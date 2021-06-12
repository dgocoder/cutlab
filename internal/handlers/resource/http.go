package resourcehandler

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type HTTPHandler struct {
	resourceController ports.ResourceController
}

func NewHTTPHandler(resourceController ports.ResourceController) *HTTPHandler {
	return &HTTPHandler{
		resourceController: resourceController,
	}
}

func (hdl *HTTPHandler) Get(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	resource, err := hdl.resourceController.Get(id)
	fmt.Println(err)
	if err != nil {
		c.JSON(500, "string")
		return
	}

	return c.JSON(200, resource)
}

func (hdl *HTTPHandler) Create(c echo.Context) (err error) {
	body := domain.Resource{}
	c.Bind(&body)

	resource, err := hdl.resourceController.Create(body.Name, body.LocationID, body.CompanyID, body.Email)
	if err != nil {
		c.JSON(500, "error")
		return
	}

	return c.JSON(200, resource)
}
