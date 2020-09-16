package main

import (
    "database/sql"
    "log"
    _ "github.com/SAP/go-hdb/driver"
)

const (
    driverName = "hdb"
    hdbDsn     = "hdb://DBADMIN:Asd123456789@618751f0-4770-4fa7-ab7e-77ed58df994b.hana.trial-us10.hanacloud.ondemand.com:443"
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