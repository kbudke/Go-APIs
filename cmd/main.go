package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

const (
	hostname     = "localhost"
	host_port    = 1234
	username     = "starwars_user"
	password     = "1234"
	databasename = "StarWars"
)

func main() {
	pg_con_string := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostname, host_port, username, password, databasename)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// We can also ping our connection which will let us know if our connection is correct /// or not then we put an error-handling code right after that.
	error = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("You are Successfully connected!")
}
