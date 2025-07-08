package service

import (
	"context"
	"customer-restful-api/model/web"
)

type CustomerService interface {
	FindAll(ctx context.Context) []web.CustomerResponse
	Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse
	FindById(ctx context.Context, customerId int) web.CustomerResponse
	Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse
	Delete(ctx context.Context, customerId int)
}
