package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"../models"
	"encoding/json"
)

var whois_servers []string

func query_domain(w http.ResponseWriter, r *http.Request){
	// Getting domain from params in get
	domain := r.URL.Query().Get("domain")

	// Retriving information from method
	information := sofos.Query_domain(domain)

	// Encoding in Json format
	json.NewEncoder(w).Encode(information)

}

func viewed_domains(w http.ResponseWriter, r *http.Request){

	// Retriving history from method
	history := sofos.Viewed_domains()
	// Encoding in Json format
	json.NewEncoder(w).Encode(history)

}

func main() {

	r:= chi.NewRouter()

	// Setting two endpoints
	r.Get("/query_domain", query_domain)
	r.Get("/viewed_domains", viewed_domains)

	// Starting server
	fmt.Println("Starting server on port :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}


}
