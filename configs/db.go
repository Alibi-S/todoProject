package config

import (
	"log"
	"os"

	controllers "github.com/Alibi-S/todoProject/controllers"
	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	options := &pg.Options{
		User:     "postgres",
		Password: "1234",
		Addr:     "localhost:5432",
		Database: "ToDoDB",
	}

	var db *pg.DB = pg.Connect(options)

	if db == nil {
		log.Printf("Failed")
		os.Exit(100)
	}

	log.Printf("Connected to database")
	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)

	return db
}
