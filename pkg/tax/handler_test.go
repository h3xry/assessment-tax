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
	bodyReqByte := `{
		"totalIncome": 500000.0,
		"wht": 0.0,
		"allowances": [
		  {
			"allowanceType": "donation",
			"amount": 0.0
		  }
		]
	  }`
	bodyReqJson, err := json.Marshal(bodyReqByte)
	assert.NoError(t, err)

	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: utils.NewValidator()}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(bodyReqJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	NewHandler(e.Group(""))
}
