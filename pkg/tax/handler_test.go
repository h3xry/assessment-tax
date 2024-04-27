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

func setupHandleCalculation(body []byte) (*httptest.ResponseRecorder, error) {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: utils.NewValidator()}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(body)))
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
	err := handler.handleCalculation()(c)
	return rec, err
}

type handleCalculationTestCase struct {
	name             string
	bodyReqInterface requestCalculation
	responseExpected string
}

func TestHandleCalculation(t *testing.T) {
	testCases := []handleCalculationTestCase{
		{
			name: "story-1-success",
			bodyReqInterface: requestCalculation{
				TotalIncome: 500000.0,
				Wht:         0.0,
				Allowances: []domain.TaxAllowance{
					{
						AllowanceType: "donation",
						Amount:        0.0,
					},
				},
			},
			responseExpected: `{"tax":29000.0}`,
		},
		{
			name: "story-2-success",
			bodyReqInterface: requestCalculation{
				TotalIncome: 500000.0,
				Wht:         25000.0,
				Allowances: []domain.TaxAllowance{
					{
						AllowanceType: "donation",
						Amount:        0.0,
					},
				},
			},
			responseExpected: `{"tax":4000.0}`,
		},
		{
			name: "story-3-success",
			bodyReqInterface: requestCalculation{
				TotalIncome: 500000.0,
				Wht:         0.0,
				Allowances: []domain.TaxAllowance{
					{
						AllowanceType: "donation",
						Amount:        200000.0,
					},
				},
			},
			responseExpected: `{"tax":19000.0}`,
		},
		{
			name: "income-600000-no-allowance",
			bodyReqInterface: requestCalculation{
				TotalIncome: 600000.0,
				Wht:         0.0,
				Allowances:  []domain.TaxAllowance{},
			},
			responseExpected: `{"tax":41000.0}`,
		},
		{
			name: "story-4-success",
			bodyReqInterface: requestCalculation{
				TotalIncome: 500000.0,
				Wht:         0.0,
				Allowances: []domain.TaxAllowance{
					{
						AllowanceType: "donation",
						Amount:        200000.0,
					},
				},
			},
			responseExpected: `{"tax":19000,"taxLevel":[{"level":"0-150,000","tax":0},{"level":"150,001-500,000","tax":19000},{"level":"500,001-1,000,000","tax":0},{"level":"1,000,001-2,000,000","tax":0},{"level":"2,000,001 ขึ้นไป","tax":0}]}`,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			bodyReqJson, err := json.Marshal(tt.bodyReqInterface)
			rec, err := setupHandleCalculation(bodyReqJson)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.responseExpected, rec.Body.String())
		})
	}
}
