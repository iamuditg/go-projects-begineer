package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/iamuditg/services"
	"log"
	"net/http"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/api/v1/fetchNifty", GetNifty50Data).Methods("GET")
	route.HandleFunc("/api/v1/fetchNiftyCompany", GetNifty50CompanyData).Methods("GET")

	http.ListenAndServe(":8080", route)

}

func GetNifty50CompanyData(writer http.ResponseWriter, request *http.Request) {
	nifty := services.Nifty{}
	data, err := nifty.FetchNiftyCompaniesData()
	if err != nil {
		log.Fatal(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(data)
}

func GetNifty50Data(writer http.ResponseWriter, request *http.Request) {
	nifty := services.Nifty{}
	data, err := nifty.FetchNiftyData()
	if err != nil {
		log.Fatal(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(data)
}
