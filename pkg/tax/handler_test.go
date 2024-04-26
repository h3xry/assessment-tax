package tax

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/h3xry/assessment-tax/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandleCalculation(t *testing.T) {
	bodyReqInterface := echo.Map{
		"totalIncome": 500000.0,
		"wht":         0.0,
		"allowances": []echo.Map{
			{
				"allowanceType": "donation",
				"amount":        0.0,
			},
		},
	}
	bodyReqJson, err := json.Marshal(bodyReqInterface)
	assert.NoError(t, err)

	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: utils.NewValidator()}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(bodyReqJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	handler := NewHandler(e.Group(""))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = handler.handleCalculation()(c)

	resJsonExpected, err := json.Marshal(echo.Map{
		"tax": "29000.0",
	})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, string(resJsonExpected), rec.Body.String())
}
