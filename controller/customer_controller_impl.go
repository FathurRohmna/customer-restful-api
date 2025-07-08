package controller

import (
	"customer-restful-api/helper"
	"customer-restful-api/model/web"
	"customer-restful-api/service"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

type CustomerControllerImpl struct {
	CustomerService service.CustomerService
	Validate        *validator.Validate
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	validate := validator.New()

	return &CustomerControllerImpl{
		CustomerService: customerService,
		Validate:        validate,
	}
}

func (controller *CustomerControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customerResponse := controller.CustomerService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CustomerControllerImpl) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customerCreateRequest := web.CustomerCreateRequest{}
	helper.ReadFromRequestBody(r, &customerCreateRequest)

	customerResponse := controller.CustomerService.Create(r.Context(), customerCreateRequest)
	webResponse := web.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CustomerControllerImpl) FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get customerId queryParams value
	customerId := ps.ByName("customerId")

	err := controller.Validate.Var(customerId, "numeric")
	helper.PanicIfError(err)

	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerResponse := controller.CustomerService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CustomerControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customerUpdateRequest := web.CustomerUpdateRequest{}
	helper.ReadFromRequestBody(r, &customerUpdateRequest)

	customerId := ps.ByName("customerId")
	err := controller.Validate.Var(customerId, "numeric")
	helper.PanicIfError(err)

	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerUpdateRequest.Id = id

	customerResponse := controller.CustomerService.Update(r.Context(), customerUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CustomerControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get customerId queryParams value
	customerId := ps.ByName("customerId")
	err := controller.Validate.Var(customerId, "numeric")
	helper.PanicIfError(err)

	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	controller.CustomerService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   204,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}
