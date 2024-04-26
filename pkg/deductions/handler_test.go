package deductions

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/h3xry/assessment-tax/mocks"
	"github.com/h3xry/assessment-tax/pkg/models"
	"github.com/h3xry/assessment-tax/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandleSetKReceipt(t *testing.T) {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: utils.NewValidator()}
	mockUsecase := new(mocks.DeductionsUsecase)
	mockUsecase.On("Find", "kReceipt").Return(&models.Deductions{
		Name:   "kReceipt",
		Amount: 20000,
	}, nil).Once()
	mockUsecase.On("Update", &models.Deductions{
		Name:   "kReceipt",
		Amount: 10000,
	}).Return(nil).Once()

	handle := NewHandler(e.Group(""), mockUsecase)
	bodyReq := requestSetKReceipt{
		Amount: 10000,
	}
	bodyJson, err := json.Marshal(bodyReq)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(bodyJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = handle.setKReceipt()(c)

	resJsonExpected, err := json.Marshal(echo.Map{
		"kReceipt": bodyReq.Amount,
	})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, string(resJsonExpected), rec.Body.String())
}
