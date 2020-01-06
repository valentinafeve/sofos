package ssl_calc

import (
  "../info"
)

// Calc SSl grade given a letter
func calc_SSL_val (letter string) int{
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
func Calc_lowest_grade(information *info.Info){
  min_val := 10000
  var min_grade string
  for _, element := range (*information).Servers {
    ssl_grade := element.SSL_grade
    if len(ssl_grade) < 1 {
      continue
    }
    ssl_val := calc_SSL_val(ssl_grade)
    if ssl_val < min_val{
      min_val = ssl_val
      min_grade = ssl_grade
    }
  }
  (*information).SSL_grade = min_grade
}
