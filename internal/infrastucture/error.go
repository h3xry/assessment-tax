package infrastucture

import (
	"fmt"
	"net/http"

	"github.com/h3xry/assessment-tax/pkg/domain"
	"github.com/h3xry/assessment-tax/pkg/utils"
	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}
	if customErr, ok := err.(domain.Error); ok {
		code = customErr.HttpCode
		message = customErr.Message
	}

	if customErr, ok := err.(utils.ValidatorError); ok {
		c.JSON(http.StatusBadRequest, echo.Map{
			"message": customErr.Message,
			"fields":  customErr.Fields,
		})
		return
	}

	if code == http.StatusInternalServerError {
		fmt.Println(err.Error())
	}
	c.JSON(code, echo.Map{
		"message": message,
	})
}
