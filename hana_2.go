package main

import (
    "database/sql"
    "log"
    _ "github.com/SAP/go-hdb/driver"
)

const driver (
    driverName = "hdb"
    hdbDsn     = "hdb://user:password@host:port"
)

func main() {
	db, err := sql.Open(driver.DriverName, driver.hdbDsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	stmt, err := db.Prepare("bulk insert into test values (?)") // Prepare bulk query.
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	
	for i := 0; i < 1000; i++ {
		if _, err := stmt.Exec(i); err != nil {
			log.Fatal(err)
		}
	}
	// Call final stmt.Exec().
	if _, err := stmt.Exec(); err != nil {
		log.Fatal(err)
	}
}