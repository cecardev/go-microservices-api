package app

import (
	"log"
	"net/http"

	"github.com/cecardev/go-microservices-api/banking/domain"
	"github.com/cecardev/go-microservices-api/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	mux := mux.NewRouter()
	//wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//routes
	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	//start the server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
