package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"encoding/json"
	"log"
	"strings"
	"../controller"
)

var whois_servers []string

func queryDomain(w http.ResponseWriter, r *http.Request){

	// TO-DO: Check parameters and data validation

	// Getting domain from params in get
	domain := r.URL.Query().Get("domain")

	// More filters
	if strings.ContainsAny(domain, "+,|!\"£$%&/()=?^*ç°§;:_>][@"){
		// TO-DO: Raise error
		tro := "Errore"
	}

	// Retriving information from method
	information := controller.QueryDomain(domain)

	// Encoding in Json format
	json.NewEncoder(w).Encode(information)

}

func viewedDomains(w http.ResponseWriter, r *http.Request){

	// TO-DO: Check parameters and data validation

	// Retriving history from method
	history := controller.ViewedDomains()

	// Encoding in Json format
	json.NewEncoder(w).Encode(history)

}

func main() {

	r:= chi.NewRouter()

	// Setting two endpoints
	r.Get("/querydomain", queryDomain)
	r.Get("/vieweddomains", viewedDomains)

	// Starting server
	fmt.Println("Starting server on port :3000")
	err := http.ListenAndServe(":3000", r)

	if err != nil {
		fmt.Println("ListenAndServe:", err)
	} else {
		log.Fatal(err)
	}

}
