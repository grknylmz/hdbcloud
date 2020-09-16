package main

import (
    "database/sql"
    "log"
    _ "github.com/SAP/go-hdb/driver"
)

const (
    driverName = "hdb"
    hdbDsn     = "hdb://user:password@host:port"
)

func main() {
	db, err := sql.Open(DriverName, TestDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	table := RandomIdentifier("testNamedArg_")
	if _, err := db.Exec(fmt.Sprintf("create table %s.%s (i integer, j integer)", TestSchema, table)); err != nil {
		log.Fatal(err)
	}
	
	var i = 0
	if err := db.QueryRow(fmt.Sprintf("select count(*) from %s.%s where i = :1 and j = :1", TestSchema, table), 1).Scan(&i); err != nil {
		log.Fatal(err)
	}
	
	if err := db.QueryRow(fmt.Sprintf("select count(*) from %s.%s where i = ? and j = :3", TestSchema, table), 1, "soso", 2).Scan(&i); err != nil {
		log.Fatal(err)
	}
	
	fmt.Print(i)
}