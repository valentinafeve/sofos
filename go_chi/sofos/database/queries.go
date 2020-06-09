package database

import (
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
  "os"
  "log"

  "../models"
)

func GetLatestGrade(domain string) string{
  db, err := sql.Open("postgres","postgresql://sofos_u@"+os.Getenv("SOFOS_HOSTNAME")+":26257/sofos?sslmode=disable")
  ssl_grade_ := ""
  if err != nil {
    println("Error connectiong to the database")
    println(err)
  }
  // Print out the balances.
  rows, err := db.Query("SELECT SSLGrade FROM DomainInformation WHERE Domain='"+domain+"'")
  if err != nil {
    fmt.Println(err)
  }
  defer rows.Close()
  for rows.Next() {
    var ssl_grade string
    if err := rows.Scan(&ssl_grade); err != nil {
        fmt.Println(err)
    }
    ssl_grade_ = ssl_grade
  }
  return ssl_grade_
}

// Save the query in the database
func SaveQuery(information *models.DomainInformation, domain string) (error){
  db, err := sql.Open("postgres","postgresql://sofos_u@"+os.Getenv("SOFOS_HOSTNAME")+":26257/sofos?sslmode=disable")
  if err != nil {
    println(err)
    println("Error when trying to establish connection to the database.")
    return err
  }

  ssl_grade := information.SSL_grade
  title := information.Title
  is_down := "false"
  if information.Is_down {
      is_down = "true"
    }

  query := `INSERT INTO DomainInformation (SSLGrade, Title, IsDown, Domain) VALUES ('`
  query += ssl_grade+`','`+title+`','`+is_down+`','`+domain+`');`
  _, err = db.Exec(query);
  if err != nil {
    println(err)
    print(query)
    println("Error when trying to insert information about the domain.")
    return err
  }

  query = `INSERT INTO HistoryQueries (domain, latestQuery) VALUES ('`
  query += domain
  query += `', CURRENT_TIMESTAMP);`
  _, err = db.Exec(query);
  if err != nil {
    println(err)
    println("Error when trying to insert information about the query.")
    return err
  }
  query = `DELETE FROM RelatedServers WHERE domain='`
  query += domain
  query += `';`
  _, err = db.Exec(query);
  if err != nil {
    println(err)
    println("Error when trying to delete information about the related servers.")
    return err
  }
  for _, server := range information.Servers {
    query = `INSERT INTO RelatedServers (domain, server) VALUES ('`
    query += domain+`','`+server.Address+`');`
    _, err = db.Exec(query);
    if err != nil {
      println(err)
      println("Error when trying to insert information about the related servers.")
      return err
    }

  }

  println("Database updated.")
  return nil
}

// Check if servers changed querying database.
func CheckIfChanged(information *models.DomainInformation, domain string){
  (*information).First = true
  db, err := sql.Open("postgres","postgresql://sofos_u@"+os.Getenv("SOFOS_HOSTNAME")+":26257/sofos?sslmode=disable")
  if err != nil {
    println("Error connectiong to the database")
    println(err)
  }
  // Print out the balances.
  rows, err := db.Query("SELECT Server FROM RelatedServers WHERE Domain='"+domain+"'")
  if err != nil {
    fmt.Println(err)
  }

  defer rows.Close()
  for rows.Next() {
    var server string
    if err := rows.Scan(&server); err != nil {
        fmt.Println(err)
    }
    olds := 0
    for i, old_server := range (*information).Servers {
    (*information).First = false
      if old_server.Address == server {
        (*information).Servers[i].Old=true
        olds ++;
        break
      }
    }
    if len( (*information).Servers ) != olds {
      (*information).Servers_changed = true
    }
  }
}

func GetQueries() ([]models.Query, error){

  queries := make([]models.Query, 0)

  log.Printf("Establishing connection between sofos and the database.")
  db, err := sql.Open("postgres","postgresql://sofos_u@"+os.Getenv("SOFOS_HOSTNAME")+":26257/sofos?sslmode=disable")
  if err != nil {
    log.Panic(err)
    return nil, err
  }
  log.Printf("Connection established.")

  rows, err := db.Query("SELECT * FROM HistoryQueries ORDER BY latestQuery DESC")
  if err != nil {
    log.Panic(err)
    return nil, err
  }

  log.Printf("Reading HistoryQueries...")
  defer rows.Close()
  for rows.Next() {
    var domain string
    var latest_query string
    if err := rows.Scan(&domain, &latest_query); err != nil {
      log.Panic(err)
      return nil, err
    }
    query := models.Query{
      Domain : domain,
      Time : latest_query,
    }
    queries = append(queries, query)
  }

  log.Printf("Data from HistoryQueries read.")
  return queries, nil
}
