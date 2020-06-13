package controller

import (
	"../database"
	"../models"
	"./utils"
	"log"
	"net"
)

// When a user asks for information about the domain, QueryDomain executes both, web scrapping and API querying in order to return specific information.
func QueryDomain(domain string) (models.DomainInformation, error) {

	// Creating a structure for saving information
	var domainInformation models.DomainInformation
	domainInformation.Servers = make([]models.Server, 0)

	// Making lookup in order to check is host is down
	_, err := net.LookupHost(domain)
	if err != nil {
		domainInformation.IsDown = true
		return domainInformation, nil
	}

	// Reading Json from the API
	log.Printf("Reading json...")
	jsonResponse := utils.ReadJson(domain)

	// Loading data into structure
	log.Printf("Loading data from json...")

	// load_from_json gets the response given after the query to the API and saves the domainInformation in the global info variable.
	for _, element := range jsonResponse.Endpoints {
		server := models.Server{
			Address:  element.IpAddress,
			SSLGrade: element.Grade,
		}

		domainInformation.Servers = append(domainInformation.Servers, server)

	}

	// Reading page from web
	log.Printf("Reading page from web...")
	page := utils.ReadWebpage(domain)

	// Loading data into structure
	log.Printf("Loading data from web...")
	logo, title := utils.LoadFromWeb(page, domain)
	domainInformation.Logo = logo
	domainInformation.Title = title

	// Calculating lowest SSL grade
	utils.CalcLowestGrade(&domainInformation)

	// Whois for each server
	log.Printf("Loading information from each server...")
	var servers [](models.Server)
	for _, server := range domainInformation.Servers {
		s := utils.WhoIs(server.Address)
		if len(s) > 10 {
			serverCountry, serverOwner := utils.LoadFromWhoIs(s)
			server.Country = serverCountry
			server.Owner = serverOwner
			servers = append(servers, server)
		} else {
		}
	}
	// Saving info into structure
	domainInformation.Servers = servers

	// Getting latest SSL Grade from the database
	log.Printf("Getting latest SSL grade")
	sslLatest, err := database.GetLatestGrade(domain)
	if err != nil {
		log.Panic(err)
		log.Panic("Error when trying to read latest SSL status.")
	}
	domainInformation.PreviousSSLGrade = sslLatest

	// Check if there are new servers
	log.Printf("Checking if servers have changed.")
	database.CheckIfChanged(&domainInformation, domain)

	// Saving query into database
	log.Printf("Saving query in database.")
	database.SaveQuery(&domainInformation, domain)

	// Returning obteined information
	return domainInformation, nil
}
