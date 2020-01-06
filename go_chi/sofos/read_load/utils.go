package read_load

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "../info"
  "strings"
  "fmt"
)

type EndPoint struct {
  IpAddress string
  Grade string
}

type JsonResponse struct {
  Endpoints []EndPoint
}

func Read_webpage(domain string) string{
  responding_web := true

  // Trying with SSL and without it
  url := "https://"+domain
  resp, err := http.Get(url)
  if err != nil {
    println("Error while resolving "+url)
    url = "http://"+domain
    resp, err = http.Get(url)
    if err != nil {
      println("Error while resolving "+url)
      responding_web = false
    }
  }
  defer func(){
    println("Host is not responding web requests...")
  }()

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

func Read_json(domain string) JsonResponse{
  var jsonResponse JsonResponse
  cont := 0
  for {
    if cont >= 5 {
      break
    }
    var url = "https://api.ssllabs.com/api/v3/analyze?host="+domain+"&all=on&grade"
    resp, _ := http.Get(url)
    bytes, _ := ioutil.ReadAll(resp.Body)
    json.Unmarshal([]byte(string(bytes)), &jsonResponse)
    fmt.Println(jsonResponse)
    if len(jsonResponse.Endpoints) > 0{
      break
    }
    cont++
  }
  return jsonResponse
}

// load_from_web gets the code given by the web scrapper and recolects useful information in order to save it in the global info variable.
func Load_from_web(page string, information *info.Info){
  if len(page) < 10{
    fmt.Println("Invalid page, omitting data loading...")
    return
  }
  if strings.Contains(page, "<head>"){
    head := strings.Split(page, "</head>")[0]
    head = strings.Split(head, "<head>")[1]
    if strings.Contains(page, `<link rel="icon"`){
      icon := strings.Split(head, `<link rel="icon"`)[1]
      icon = strings.Split(icon, `href="`)[1]
      icon = strings.Split(icon, `"`)[0]
      (*information).Logo=icon
    }
    if strings.Contains(page, "<title>"){
      title := strings.Split(head, "</title>")[0]
      title = strings.Split(title, "<title")[1]
      title = strings.Split(title, ">")[1]
      (*information).Title=title
    } else{
      if strings.Contains(page, `<meta property="og:title"`){
        title := strings.Split(head, `<meta property="og:title"`)[1]
        title = strings.Split(title, `">`)[0]
        title = strings.Split(title, `content=`)[1]
        (*information).Title=title
      }
    }
  }
}

// load_from_json gets the response given after the query to the API and saves the information in the global info variable.
func Load_from_json(jsonResponse JsonResponse, information *info.Info ){
  for _, element := range jsonResponse.Endpoints{
    server := info.Server{
      Address : element.IpAddress,
      SSL_grade : element.Grade,
    }
    (*information).Servers = append(information.Servers, server)
  }
}
