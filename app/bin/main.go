package main

import (
  "database/sql"
  "fmt"
  _ "github.com/denisenkom/go-mssqldb"
)

const (
  host     = "db"
  user     = "sa"
  password = "z!oBx1ab"
  dbname   = "master"
)

func main() {
  sqlserverInfo := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30",
     user, password, host, dbname)

  db, err := sql.Open("sqlserver", sqlserverInfo)

  if err != nil {
    panic(err)
  }

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")

  db.Close()
}
