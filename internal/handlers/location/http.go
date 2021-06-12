package locationhandler

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type HTTPHandler struct {
	locationController ports.LocationController
}

func NewHTTPHandler(locationController ports.LocationController) *HTTPHandler {
	return &HTTPHandler{
		locationController: locationController,
	}
}

func (hdl *HTTPHandler) Get(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	location, err := hdl.locationController.Get(id)
	fmt.Println(err)
	if err != nil {
		c.JSON(500, "string")
		return
	}

	return c.JSON(200, location)
}

func (hdl *HTTPHandler) Create(c echo.Context) (err error) {
	body := domain.Location{}
	c.Bind(&body)

	location, err := hdl.locationController.Create(body.CompanyID, body.Name)
	if err != nil {
		c.JSON(500, "error")
		return
	}

	return c.JSON(200, location)
}
