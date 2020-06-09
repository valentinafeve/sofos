package controller

import(
  "log"

  "../models"
  "../database"
)

func ViewedDomains() (models.History, error){

  var history models.History
  var queries []models.Query

  queries, err := database.GetQueries()

  if (err == nil){
    history.Queries = queries
  } else {
    log.Printf("Data from HistoryQueries wasn't read succesfully.")
    return history, err
  }

  return history, nil
}
