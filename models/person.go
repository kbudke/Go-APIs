package models

type Person struct {
	name      string   `json:"name"`
	gender    string   `json:"gender"`
	homeworld string   `json:"homeworld"`
	starships []string `json:"starships"`
}
