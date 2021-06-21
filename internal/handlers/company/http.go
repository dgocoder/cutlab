package companyhandler

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/testd/cutlab/internal/core/ports"
)

type HTTPHandler struct {
	companyController ports.CompanyController
}

func NewHTTPHandler(companyController ports.CompanyController) *HTTPHandler {
	return &HTTPHandler{
		companyController: companyController,
	}
}

func (hdl *HTTPHandler) Get(c echo.Context) (err error) {
	id, _ := uuid.Parse(c.Param("id"))
	company, err := hdl.companyController.Get(id)
	fmt.Println(err)
	if err != nil {
		c.JSON(500, "string")
		return
	}

	return c.JSON(200, company)
}

func (hdl *HTTPHandler) Create(c echo.Context) (err error) {
	body := ports.CreateCompanyController{}
	c.Bind(&body)

	company, err := hdl.companyController.Create(body)
	if err != nil {
		c.JSON(500, "error")
		return
	}

	return c.JSON(200, company)
}
