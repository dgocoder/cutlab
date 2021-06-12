package eventhandler

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type HTTPHandler struct {
	eventController ports.EventController
}

func NewHTTPHandler(eventController ports.EventController) *HTTPHandler {
	return &HTTPHandler{
		eventController: eventController,
	}
}

func (hdl *HTTPHandler) Get(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	event, err := hdl.eventController.Get(id)
	fmt.Println(err)
	if err != nil {
		c.JSON(500, "string")
		return
	}

	return c.JSON(200, event)
}

func (hdl *HTTPHandler) Create(c echo.Context) (err error) {
	body := domain.Event{}
	c.Bind(&body)

	event, err := hdl.eventController.Create(body.LocationID, body.ServiceID, body.ResourceID, body.CustomerID, body.StartAt, body.EndAt, body.Type)
	if err != nil {
		c.JSON(500, "error")
		return
	}

	return c.JSON(200, event)
}
