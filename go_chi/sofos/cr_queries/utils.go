package cr_queries

import (
  "database/sql"
  _ "github.com/lib/pq"
  "../info"
  "fmt"
)

func Get_latest_grade(domain string) string{
  db, err := sql.Open("postgres","postgresql://sofos_u@archievaldo:26257/sofos?sslmode=disable")
  ssl_grade_ := ""
  if err != nil {
    println("Error connectiong to the database")
    println(err)
  }
  // Print out the balances.
  rows, err := db.Query("SELECT SSL_grade FROM Domain_info WHERE Domain='"+domain+"'")
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
func Save_query(information info.Info, domain string){
  db, err := sql.Open("postgres","postgresql://sofos_u@archievaldo:26257/sofos?sslmode=disable")
  if err != nil {
    println("Error connectiong to the database")
    println(err)
  }

  ssl_grade := information.SSL_grade
  title := information.Title
  is_down := "false"
  if information.Is_down {
      is_down = "true"
    }

  query := `INSERT INTO Domain_info (SSL_grade, Title, Is_down, Domain) VALUES ('`
  query += ssl_grade+`','`+title+`','`+is_down+`','`+domain+`');`
  _, err = db.Exec(query);
  if err != nil {
    println(err)
  }
  query = `INSERT INTO History_queries (domain, latest_query) VALUES ('`
  query += domain
  query += `', CURRENT_TIMESTAMP);`
  _, err = db.Exec(query);
  if err != nil {
    println(err)
  }
  query = `DELETE FROM Related_servers WHERE domain='`
  query += domain
  query += `';`
  _, err = db.Exec(query);
  if err != nil {
    println(err)
  }
  for _, server := range information.Servers {
    query = `INSERT INTO Related_servers (domain, server) VALUES ('`
    query += domain+`','`+server.Address+`');`
    _, err = db.Exec(query);
    if err != nil {
      println(err)
    }

  }
}

// Check if servers changed querying database.
func Check_if_changed(information *info.Info, domain string){
  db, err := sql.Open("postgres","postgresql://sofos_u@archievaldo:26257/sofos?sslmode=disable")
  if err != nil {
    println("Error connectiong to the database")
    println(err)
  }
  // Print out the balances.
  rows, err := db.Query("SELECT Server FROM Related_servers WHERE Domain='"+domain+"'")
  if err != nil {
    fmt.Println(err)
  }
  defer rows.Close()
  for rows.Next() {
    var server string
    if err := rows.Scan(&server); err != nil {
        fmt.Println(err)
    }
    cont := 0
    olds := 0
    for _, old_server := range (*information).Servers {
      if old_server.Address == server {
        (*information).Servers[cont].Old=true
        olds ++;
        break
      }
      cont++
    }
    if len( (*information).Servers ) != olds {
      (*information).Servers_changed = true
    }
  }
}
