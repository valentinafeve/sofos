package controller

import (
  "net"
  "fmt"
  "../models"
  "./utils"
)

// When a user asks for information about the domain, QueryDomain executes both, web scrapping and API querying in order to return specific information.
func QueryDomain(domain string) models.DomainInformation{

  // Creating a structure for saving information
  var domainInformation models.DomainInformation
  domainInformation.Servers = make([]models.Server, 0)

  // Making lookup in order to check is host is down
  _, err := net.LookupHost(domain)
  if err != nil {
    domainInformation.Is_down = true
    return domainInformation
  }

  // Reading Json from the API
  println("Reading json...")
  jsonResponse := utils.ReadJson(domain)
  fmt.Println(jsonResponse)

  // Loading data into structure
  println("Loading data from json...")

  // load_from_json gets the response given after the query to the API and saves the domainInformation in the global info variable.
  for _, element := range jsonResponse.Endpoints{
    server := models.Server{
      Address : element.IpAddress,
      SSL_grade : element.Grade,
    }

    domainInformation.Servers = append(domainInformation.Servers, server)

  }


  // Reading page from web
  println("Reading page from web...")
  page := utils.ReadWebpage(domain)

  // Loading data into structure
  println("Loading data from web...")
  utils.LoadFromWeb(page, domain)

  // Calculating lowest SSL grade
  utils.CalcLowestGrade(&domainInformation)

  // Whois for each server
  println("Loading information from each server...")
  var servers [](models.Server)
  for _, server := range domainInformation.Servers {
    s := utils.WhoIs(server.Address)
    if len(s) > 10 {
      serverCountry, serverOwner := utils.LoadFromWhoIs(s)
      server.Country = serverCountry
      server.Owner = serverOwner
      servers = append(servers, server)
    } else{
    }
  }
  // Saving info into structure
  domainInformation.Servers = servers
  //
  // ssl_latest := utils.GetLatestGrade(domain)
  // information.Previous_SSL_grade = ssl_latest
  //
  // // Check if there are new servers
  // println("Checking if servers have changed.")
  // cr_queries.Check_if_changed(&information, domain)
  //
  // // Saving query into database
  // println("Saving query in database.")
  // cr_queries.Save_query(information, domain)

  // Returning obteined information
  return domainInformation
}
