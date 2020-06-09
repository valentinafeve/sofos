package controller

import(
  "../models"
)

func ViewedDomains() models.History{

  var history models.History
  var queries []models.Query

  queries = models.GetQueries()
  history.Queries = queries

  return history
}
