//go:build unit
// +build unit

package deductions

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/h3xry/assessment-tax/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUpdate(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing sql mock: %s", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("error initializing GORM: %s", err)
	}

	repo := &repository{DB: gormDB}

	model := &models.Deductions{Name: "kReceipt", Amount: 50000}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "deductions"`)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err := repo.Update(model); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unfulfilled expectations: %s", err)
	}
}
