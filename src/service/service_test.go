package service

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"sample-go-with-gorm.com/src/db"
)

func TestSelectAllItems(t *testing.T) {
	mockDB, mock := db.NewMockDB()

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "items" WHERE "items"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{}))

	service := &Service{}
	service.SelectAllItems(mockDB)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find User: %v", err)
	}
}

func TestSelectAllOrders(t *testing.T) {
	mockDB, mock := db.NewMockDB()

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "orders" WHERE "orders"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{}))

	service := &Service{}
	service.SelectAllOrders(mockDB)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find User: %v", err)
	}
}
