package database

import (
	"github.com/gobuffalo/pop/v6"
	"log"
)

var Connection *pop.Connection

func ConnectDb() {
	var err error
	Connection, err = pop.Connect("development")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
}
