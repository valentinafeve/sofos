package controller

import (
  "net"
  "math/rand"
  "strings"
  "io/ioutil"
  "time"
  "fmt"
  _ "github.com/lib/pq"
)

// Whois function send the domain to a random selected server. Returns information about it.
func whois(domain string) string{

  // Whois server according to the command whois
  whois_server := "199.71.0.46"

  // Setting the tcp connection
  address :=  whois_server+":43"
  connection, err := net.Dial("tcp",address)
   if err != nil {
       fmt.Println(err)
   }
  defer connection.Close()
  if err != nil {
      fmt.Println(err)
  }

  // The domain which we'll query
  domain_to_query := domain

  // Sending the message
  connection.Write([]byte("n + "+domain_to_query+ "\n"))
  result, err := ioutil.ReadAll(connection)
  if err != nil {
    fmt.Println("Error reading data from server, check your connection.")
  }

  s := string(result)

  // Select a random server from the list
  rand.Seed(time.Now().UnixNano())
  return s
}

func load_from_whois(s string, server *info.Server){
  // Reading results from whois
  data := strings.Split(s, "\n")
  for _, element := range data {
    // Getting country info from string
  	if strings.Contains(element, "Country:"){
  		country := strings.Split(element, ":")[1]
  		(*server).Country = country
  	}
    // Getting org info from string
  	if strings.Contains(element, "Organization:"){
  		org := strings.Split(element, ":")[1]
  		(*server).Owner = org
  	}
  }
}

// When a user asks for information about the domain, Query_domain executes both, web scrapping and API querying in order to return specific information.
func Query_domain(domain string) info.Info{

  // Creating a structure for saving information
  var information info.Info
  information.Servers = make([]info.Server, 0)

  // Multiple filters before querying for host
  if !strings.Contains(domain, ".") || len(domain) < 4 || strings.Contains(domain, " "){
    information.Status = "Invalid domain name."
    return information
  }

  // More filters
  if strings.ContainsAny(domain, "+,|!\"£$%&/()=?^*ç°§;:_>][@"){
    information.Status = "Invalid characters in domain name."
    return information
  }

  // Making lookup in order to check is host is down
  _, err := net.LookupHost(domain)
  if err != nil {
    information.Is_down = true
    return information
  }

  // Reading Json from the API
  println("Reading json...")
  jsonResponse := read_load.Read_json(domain)
  fmt.Println(jsonResponse)
  // Loading data into structure
  println("Loading data from json...")
  read_load.Load_from_json(jsonResponse, &information)

  // Reading page from web
  println("Reading page from web...")
  page := read_load.Read_webpage(domain)
  // Loading data into structure
  println("Loading data from web...")
  read_load.Load_from_web(page, &information, domain)

  // Calculating lowest SSL grade
  ssl_calc.Calc_lowest_grade(&information)

  // Whois for each server
  println("Loading information from each server...")
  var servers [](info.Server)
  for _, server := range information.Servers {
    s := whois(server.Address)
    if len(s) > 10 {
      load_from_whois(s, &server)
      servers = append(servers, server)
    } else{
    }
  }
  // Saving info into structure
  information.Servers = servers

  ssl_latest := cr_queries.Get_latest_grade(domain)
  information.Previous_SSL_grade = ssl_latest

  // Check if there are new servers
  println("Checking if servers have changed.")
  cr_queries.Check_if_changed(&information, domain)

  // Saving query into database
  println("Saving query in database.")
  cr_queries.Save_query(information, domain)

  // Returning obteined information
  return information
}
