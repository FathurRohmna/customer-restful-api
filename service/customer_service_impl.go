package service

import (
	"context"
	"customer-restful-api/exception"
	"customer-restful-api/helper"
	"customer-restful-api/model/domain"
	"customer-restful-api/model/web"
	"customer-restful-api/repository"
	"database/sql"

	"github.com/go-playground/validator"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCustomerService(customerRepository repository.CustomerRepository, DB *sql.DB) *CustomerServiceImpl {
	validate := validator.New()
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CustomerServiceImpl) FindAll(ctx context.Context) []web.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customers := service.CustomerRepository.FindAll(ctx, tx)

	return helper.ToCustomerResponses(customers)
}

func (service *CustomerServiceImpl) Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse {
	// validate struct 
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer := domain.Customer{
		Name:  request.Name,
		Email: request.Email,
		Phone: request.Phone,
	}

	customer = service.CustomerRepository.Save(ctx, tx, customer)

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) FindById(ctx context.Context, customerId int) web.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, customerId)
	if err != nil {
		// case if user not found
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse {
	// validate struct 
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		// case if user not found
		panic(exception.NewNotFoundError(err.Error()))
	}

	customer.Name = request.Name
	customer.Email = request.Email
	customer.Phone = request.Phone

	customer = service.CustomerRepository.Update(ctx, tx, customer)

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) Delete(ctx context.Context, customerId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, customerId)
	if err != nil {
		// case if user not found
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CustomerRepository.Delete(ctx, tx, customer)
}
