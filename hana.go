package main

import (
    "database/sql"
    "log"
    _ "github.com/SAP/go-hdb/driver"
)

const (
    driverName = "hdb"
    hdbDsn     = ""
)

func main() {
    db, err := sql.Open(driverName, hdbDsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
}
