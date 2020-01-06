package sofos

import(
  "fmt"
  // "net/http"
  "./info"
  "database/sql"
)

func Viewed_domains() info.History{
  var history info.History
  var queries []info.Query
  db, err := sql.Open("postgres","postgresql://sofos_u@archievaldo:26257/sofos?sslmode=disable")
  if err != nil {
    println("Error connectiong to the database")
    println(err)
  }
  rows, err := db.Query("SELECT * FROM History_queries ORDER BY latest_query DESC")
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
    query := info.Query{
      Domain : domain,
      Time : latest_query,
    }
    queries = append(queries, query)
  }
  history.Queries = queries
  return history
}
