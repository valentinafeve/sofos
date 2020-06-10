package database

import (
  "database/sql"
  _ "github.com/lib/pq"
  "os"
  "log"

  "../models"
)

func GetLatestGrade(domain string) (string, error){

  db, err := sql.Open("postgres","postgresql://sofos_u@"+os.Getenv("SOFOS_HOSTNAME")+":26257/sofos?sslmode=disable")
  ssl_grade_ := ""
  if err != nil {
    log.Panic(err)
    log.Panic("Error when trying to establlish a connection to the database")
    return "", err
  }

  // Print out the balances.
  rows, err := db.Query("SELECT SSLGrade FROM DomainInformation WHERE Domain='"+domain+"'")
  if err != nil {
    log.Panic(err)
    log.Panic("Error when trying to query data from the database")
    return "", err
  }
  defer rows.Close()
  for rows.Next() {
    var ssl_grade string
    if err := rows.Scan(&ssl_grade); err != nil {
      log.Panic(err)
      log.Panic("Error when trying to read data from the database")
      return "", err
    }
    ssl_grade_ = ssl_grade
  }
  return ssl_grade_, nil
}

// Save the query in the database
func SaveQuery(information *models.DomainInformation, domain string) (error){

  log.Printf("Saving information in the database...")

  db, err := sql.Open("postgres","postgresql://sofos_u@"+os.Getenv("SOFOS_HOSTNAME")+":26257/sofos?sslmode=disable")
  if err != nil {
    log.Panic(err)
    log.Panic("Error when trying to establlish a connection to the database")
    return err
  }

  sslGrade := information.SSLGrade
  title := information.Title
  isDown := "false"
  if information.Is_down {
      isDown = "true"
    }

  log.Printf("Saving information into DomainInformation...")
  query := `INSERT INTO DomainInformation (SSLGrade, Title, IsDown, Domain) VALUES ('`
  query += sslGrade+`','`+title+`','`+isDown+`','`+domain+`');`
  _, err = db.Exec(query);
  if err != nil {
    log.Panic(err)
    log.Panic("Error when trying to insert information about the domain.")
    return err
  }

  log.Printf("Saving information into HistoryQueries...")
  query = `INSERT INTO HistoryQueries (domain, latestQuery) VALUES ('`
  query += domain
  query += `', CURRENT_TIMESTAMP);`
  _, err = db.Exec(query);
  if err != nil {
    log.Panic(err)
    log.Panic("Error when trying to insert information about the query.")
    return err
  }
  query = `DELETE FROM RelatedServers WHERE domain='`
  query += domain
  query += `';`
  _, err = db.Exec(query);
  if err != nil {
    log.Panic(err)
    log.Panic("Error when trying to delete information about the related servers.")
    return err
  }

  log.Printf("Saving information into RelatedServers...")
  for _, server := range information.Servers {
    query = `INSERT INTO RelatedServers (domain, server) VALUES ('`
    query += domain+`','`+server.Address+`');`
    _, err = db.Exec(query);
    if err != nil {
      log.Panic(err)
      log.Panic("Error when trying to insert information about the related servers.")
      return err
    }

  }

  log.Printf("Database updated.")
  return nil
}

// Check if servers changed querying database.
func CheckIfChanged(information *models.DomainInformation, domain string) (error){
  (*information).First = true
  db, err := sql.Open("postgres","postgresql://sofos_u@"+os.Getenv("SOFOS_HOSTNAME")+":26257/sofos?sslmode=disable")
  if err != nil {
    log.Panic(err)
    log.Panic("Error when trying to insert information about the domain.")
    return err
  }

  // Print out the balances.
  rows, err := db.Query("SELECT Server FROM RelatedServers WHERE Domain='"+domain+"'")
  if err != nil {
    log.Panic(err)
  }

  defer rows.Close()
  for rows.Next() {
    var server string
    if err := rows.Scan(&server); err != nil {
      log.Panic(err)
      log.Panic("Error when trying to read data from RelatedServers.")
      return err
    }
    olds := 0
    for i, oldServer := range (*information).Servers {
    (*information).First = false
      if oldServer.Address == server {
        (*information).Servers[i].Old=true
        olds ++;
        break
      }
    }
    if len( (*information).Servers ) != olds {
      (*information).ServersChanged = true
    }
  }

  return nil
}

func GetQueries() ([]models.Query, error){

  queries := make([]models.Query, 0)

  log.Printf("Establishing connection between sofos and the database.")
  db, err := sql.Open("postgres","postgresql://sofos_u@"+os.Getenv("SOFOS_HOSTNAME")+":26257/sofos?sslmode=disable")
  if err != nil {
    log.Panic(err)
    log.Panic("Error when trying to establlish a connection to the database")
    return nil, err
  }
  log.Printf("Connection established.")

  rows, err := db.Query("SELECT * FROM HistoryQueries ORDER BY latestQuery DESC")
  if err != nil {
    log.Panic(err)
    log.Panic("Error when trying to read data from HistoryQueries.")
    return nil, err
  }

  log.Printf("Reading HistoryQueries...")
  defer rows.Close()
  for rows.Next() {
    var domain string
    var latest_query string
    if err := rows.Scan(&domain, &latest_query); err != nil {
      log.Panic("Error when trying to read data from HistoryQueries.")
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
