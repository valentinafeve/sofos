package models

import (
  "os"
  "database/sql"
  "fmt"
)

type History struct{
  Queries []Query
}

func GetQueries() []Query {

  queries := make([]Query, 0)  

  db, err := sql.Open("postgres","postgresql://"+os.Getenv("SOFOS_DATABASE")+"@"+os.Getenv("SOFOS_HOSTNAME")+":26257/sofos?sslmode=disable")
  if err != nil {
    println("Error connectiong to the database")
    println(err)
  }

  rows, err := db.Query("SELECT * FROM HistoryQueries ORDER BY latest_query DESC")
  if err != nil {
    fmt.Println(err)
  }

  defer rows.Close()
  for rows.Next() {
    var domain string
    var latest_query string
    if err := rows.Scan(&domain, &latest_query); err != nil {
        fmt.Println(err)
    }
    query := Query{
      Domain : domain,
      Time : latest_query,
    }

    queries = append(queries, query)

  }

  return queries

}
