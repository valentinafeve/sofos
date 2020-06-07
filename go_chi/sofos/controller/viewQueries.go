package controller

import(
  "fmt"
  "../models"
)

func ViewedDomains() models.History{

  var history models.History
  var queries []models.Query

  queries = models.GetQueries()
  
  return history
}
