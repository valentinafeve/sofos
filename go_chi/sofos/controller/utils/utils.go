package utils

import (
  "net"
  "math/rand"
  "strings"
  "io/ioutil"
  "time"
  "fmt"
  "net/http"
  "encoding/json"
  "log"
)

// Whois function send the domain to a random selected server. Returns information about it.
func WhoIs(domain string) string{

  // Whois server according to the command whois
  whoIsServer := "199.71.0.46"

  // Setting the tcp connection
  address :=  whoIsServer+":43"
  connection, err := net.Dial("tcp",address)
   if err != nil {
       log.Panic(err)
   }

  defer connection.Close()
  if err != nil {
      log.Panic(err)
  }

  // The domain which we'll query
  domainToQuery := domain

  // Sending the message
  connection.Write([]byte("n + "+domainToQuery+ "\n"))
  result, err := ioutil.ReadAll(connection)

  if err != nil {
    log.Panic(err)
    log.Panicf("Error reading data from server, check your connection.")
  }

  s := string(result)

  // Select a random server from the list
  rand.Seed(time.Now().UnixNano())
  return s
}

func LoadFromWhoIs(s string) (string, string) {

  // Reading results from whois
  data := strings.Split(s, "\n")

  var serverCountry string
  var serverOwner string

  for _, element := range data {
    // Getting country info from string
  	if strings.Contains(element, "Country:"){
  		country := strings.Split(element, ":")[1]
  		serverCountry = country
  	}
    // Getting org info from string
  	if strings.Contains(element, "Organization:"){
  		org := strings.Split(element, ":")[1]
  		serverOwner = org
  	}
  }

  return serverCountry, serverOwner
}

type EndPoint struct {
  IpAddress string
  Grade string
}

type JsonResponse struct {
  Endpoints []EndPoint
}

func ReadWebpage(domain string) string{
  responding_web := true

  // Trying with SSL and without it
  url := "https://"+domain
  resp, err := http.Get(url)
  if err != nil {
    url = "http://"+domain
    resp_, err_ := http.Get(url)
    resp = resp_
    println(err)
    if err_ != nil {
      println("Error while resolving "+url)
      responding_web = false
      println(err_)
    }
  }

  if (responding_web){
    // Reading response
    bytes, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
      print(err)
    }
    resp.Body.Close()
    // Returning page
    return string(bytes)
  }

  // Returning void string
  return ""
}

func ReadJson(domain string) JsonResponse{
  var jsonResponse JsonResponse
  cont := 0
  info_obtained := false
  for {
    if cont > 3 || info_obtained {
      break;
    }
    var url = "https://api.ssllabs.com/api/v3/analyze?host="+domain+"&all=done&grade"
    resp, _ := http.Get(url)
    bytes, _ := ioutil.ReadAll(resp.Body)
    json.Unmarshal([]byte(string(bytes)), &jsonResponse)
    if len(jsonResponse.Endpoints) > 0{
      for _, endpoint := range jsonResponse.Endpoints {
        if (endpoint.Grade != ""){
          info_obtained = true;
        }
      }
    }
    cont++;
  }
  return jsonResponse
}

// load_from_web gets the code given by the web scrapper and recolects useful information in order to save it in the global info variable.
func LoadFromWeb(page string, domain string) (string, string){

  var logo string
  var title string

  if len(page) < 10{
    fmt.Println("Invalid page, omitting data loading...")
  }

  if strings.Contains(page, "<head>"){
    head := strings.Split(page, "</head>")[0]
    head = strings.Split(head, "<head>")[1]
    if strings.Contains(head, `<link rel="icon"`){
      icon := strings.Split(head, `<link rel="icon"`)[1]
      icon = strings.Split(icon, `href="`)[1]
      icon = strings.Split(icon, `"`)[0]
      logo = icon
    } else {
      if strings.Contains(head, `<link rel="shortcut icon"`){
        icon := strings.Split(head, `<link rel="shortcut icon"`)[1]
        icon = strings.Split(icon, `href="`)[1]
        icon = strings.Split(icon, `"`)[0]
        logo = icon
      } else {
        logo = domain+"/favicon.ico"
      }
    }
    if strings.Contains(page, "<title>"){
      title = strings.Split(head, "</title>")[0]
      title = strings.Split(title, "<title")[1]
      title = strings.Split(title, ">")[1]
      title = title
    } else{
      if strings.Contains(page, `<meta property="og:title"`){
        title = strings.Split(head, `<meta property="og:title"`)[1]
        title = strings.Split(title, `">`)[0]
        title = strings.Split(title, `content=`)[1]
        title = title
      }
    }
  }

  return logo, title
}
