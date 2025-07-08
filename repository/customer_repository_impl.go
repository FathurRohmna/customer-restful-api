package repository

import (
	"context"
	"customer-restful-api/helper"
	"customer-restful-api/model/domain"
	"database/sql"
	"errors"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (r *CustomerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer {
	query := `
		SELECT * FROM customers
	`
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	var customers []domain.Customer
	for rows.Next() {
		customer := domain.Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Phone, &customer.CreatedAt, &customer.UpdatedAt)
		helper.PanicIfError(err)
		customers = append(customers, customer)
	}

	return customers
}

func (r *CustomerRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	query := `
		INSERT INTO customers (name, email, phone) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at
	`

	var savedCustomer domain.Customer
	err := tx.QueryRowContext(ctx, query, customer.Name, customer.Email, customer.Phone).Scan(&savedCustomer.Id, &savedCustomer.CreatedAt, &savedCustomer.UpdatedAt)
	helper.PanicIfError(err)

	customer.Id = int(savedCustomer.Id)
	customer.CreatedAt = savedCustomer.CreatedAt
	customer.UpdatedAt = savedCustomer.UpdatedAt

	return customer
}

func (r *CustomerRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, customerId int) (domain.Customer, error) {
	query := `
		SELECT id, name, email, phone, created_at, updated_at
		FROM customers
		WHERE id = $1
	`
	row := tx.QueryRowContext(ctx, query, customerId)

	customer := domain.Customer{}
	err := row.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Phone, &customer.CreatedAt, &customer.UpdatedAt)

	if err != nil {
		// return error if customer not found
		if err == sql.ErrNoRows {
			return customer, errors.New("Customer not found")
		}
		return customer, err
	}
	return customer, nil
}

func (r *CustomerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	query := `UPDATE customers 
			  SET name = $1, email = $2, phone = $3, updated_at = NOW()
			  WHERE id = $4`
	_, err := tx.ExecContext(ctx, query, customer.Name, customer.Email, customer.Phone, customer.Id)
	helper.PanicIfError(err)

	customer, err = r.FindById(ctx, tx, customer.Id)
	helper.PanicIfError(err)

	return customer
}

func (r *CustomerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customer domain.Customer) {
	query := `DELETE FROM customers WHERE id = $1`
	_, err := tx.ExecContext(ctx, query, customer.Id)
	helper.PanicIfError(err)
}
