package tax

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/h3xry/assessment-tax/mocks"
	"github.com/h3xry/assessment-tax/pkg/domain"
	"github.com/h3xry/assessment-tax/pkg/models"
	"github.com/h3xry/assessment-tax/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandleCalculation(t *testing.T) {
	bodyReqInterface := requestCalculation{
		TotalIncome: 500000.0,
		Wht:         0.0,
		Allowances: []domain.TaxAllowance{
			{
				AllowanceType: "donation",
				Amount:        0.0,
			},
		},
	}
	bodyReqJson, err := json.Marshal(bodyReqInterface)
	assert.NoError(t, err)

	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: utils.NewValidator()}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(bodyReqJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	useCase := NewUseCase()

	deductionUsecase := new(mocks.DeductionsUsecase)
	deductionUsecase.On("Find", "personalDeduction").Return(&models.Deductions{
		Name:   "personalDeduction",
		Amount: 60000,
	}, nil).Once()

	handler := NewHandler(e.Group(""), useCase, deductionUsecase)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = handler.handleCalculation()(c)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	resJsonExpected, err := json.Marshal(echo.Map{
		"tax": 29000.0,
	})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, string(resJsonExpected), rec.Body.String())
}
