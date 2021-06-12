package customerhandler

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/testd/cutlab/internal/core/domain"
	"github.com/testd/cutlab/internal/core/ports"
)

type HTTPHandler struct {
	customerController ports.CustomerController
}

func NewHTTPHandler(customerController ports.CustomerController) *HTTPHandler {
	return &HTTPHandler{
		customerController: customerController,
	}
}

func (hdl *HTTPHandler) Get(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	customer, err := hdl.customerController.Get(id)
	fmt.Println(err)
	if err != nil {
		c.JSON(500, "string")
		return
	}

	return c.JSON(200, customer)
}

func (hdl *HTTPHandler) Create(c echo.Context) (err error) {
	body := domain.Customer{}
	c.Bind(&body)

	customer, err := hdl.customerController.Create(body.Lastname, body.Firstname, body.Email)
	if err != nil {
		c.JSON(500, "error")
		return
	}

	return c.JSON(200, customer)
}
