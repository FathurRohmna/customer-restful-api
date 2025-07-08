package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CustomerController interface {
	FindAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
