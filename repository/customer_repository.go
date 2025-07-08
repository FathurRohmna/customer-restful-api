package repository

import (
	"context"
	"customer-restful-api/model/domain"
	"database/sql"
)

type CustomerRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer
	Save(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer
	FindById(ctx context.Context, tx *sql.Tx, customerId int) (domain.Customer, error)
	Update(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer
	Delete(ctx context.Context, tx *sql.Tx, customer domain.Customer)
}
