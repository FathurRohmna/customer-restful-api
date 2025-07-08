package main

import (
	"customer-restful-api/app"
	"customer-restful-api/controller"
	"customer-restful-api/exception"
	"customer-restful-api/helper"
	"customer-restful-api/repository"
	"customer-restful-api/service"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Load environment variables 
	err := godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db := app.NewDB()

	// Set up repository, service, and controller
	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db)
	customerController := controller.NewCustomerController(customerService)

	router := httprouter.New()

	router.GET("/api/customers", customerController.FindAll)
	router.POST("/api/customers", customerController.Create)
	router.GET("/api/customers/:customerId", customerController.FindById)
	router.PUT("/api/customers/:customerId", customerController.Update)
	router.DELETE("/api/customers/:customerId", customerController.Delete)

	// Set customize panic handler
	router.PanicHandler = exception.ErrorHandler

	// Wrap CORS middleware to handle cross requests
	handler := helper.CORSWrapper(router)

	http.ListenAndServe(":"+port, handler)
}
