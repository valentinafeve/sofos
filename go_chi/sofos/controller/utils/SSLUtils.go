package utils

import(
  "../../models"
)

// Calc SSl grade given a letter
func calcSSLVal (letter string) int{
  fc := letter[0]
  val := 1000-int(fc)*3
  if len(letter) > 1{
    sign := string(letter[1])
    switch sign {
    case "+":{
      val++;
    }
    case "-":{
      val--;
    }
    }
  }
  return val
}

// Calc and set lowest SSL grade given all the informatio
func CalcLowestGrade(information *models.DomainInformation){
  minVal := 10000
  var minGrade string
  for _, element := range (*information).Servers {
    sslGrade := element.SSLGrade
    if len(sslGrade) < 1 {
      continue
    }
    sslVal := calcSSLVal(sslGrade)
    if sslVal < minVal{
      minVal = sslVal
      minGrade = sslGrade
    }
  }
  (*information).SSLGrade = minGrade
}
