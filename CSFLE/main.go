package main

import (
	"GoLandWorkSpace/CSFLE/services"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", services.GetAll).Methods("GET")
	r.HandleFunc("/sort", services.SortByParams).Methods("GET")
	r.HandleFunc("/get", services.GetAllWithEachParam).Methods("GET")
	r.HandleFunc("/login", services.LogIn).Methods("POST")
	r.HandleFunc("/add", services.AddEmployeeDetails).Methods("POST")
	r.HandleFunc("/update", services.UpdateEmployeeDetails).Methods("PUT")
	r.HandleFunc("/delete", services.DeleteEmployeeDetails).Methods("DELETE")
	//r.HandleFunc("/token", services.ReqHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}
