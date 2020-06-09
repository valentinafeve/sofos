package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"encoding/json"
	"log"
	"strings"
	"../controller"
)

var whois_servers []string

func queryDomain(w http.ResponseWriter, r *http.Request) {

	// TO-DO: Check parameters and data validation

	// Getting domain from params in get
	domain := r.URL.Query().Get("domain")
	if (len(domain) < 3){
		http.Error(w, http.StatusText(400), 400)
		return
	}

	// More filters
	if strings.ContainsAny(domain, "+,|!\"£$%&/()=?^*ç°§;:_>][@ "){

		// TO-DO: Raise error
		http.Error(w, http.StatusText(400), 400)
		return
	}

	// Retriving information from method
	information, err := controller.QueryDomain(domain)

	if ( err != nil ){
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// Encoding in Json format
	json.NewEncoder(w).Encode(information)

}

func viewedDomains(w http.ResponseWriter, r *http.Request){

	// TO-DO: Check parameters and data validation

	// Retriving history from method
	history, err := controller.ViewedDomains()

	if ( err != nil ){
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// Encoding in Json format
	json.NewEncoder(w).Encode(history)

}

func main() {

	r:= chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"http://localhost:8081"},
    AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

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
