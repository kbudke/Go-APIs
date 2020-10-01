package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	hostname      = "localhost"
	host_port     = 1234
	username      = "starwars_user"
	password      = "starwars"
	database_name = "StarWars"
)

func main() {
	pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, database_name)

	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
