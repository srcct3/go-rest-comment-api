package main

import (
	"fmt"

	"github.com/srcct3/go-rest-comment-api/internal/db"
)

// Run - is responsible for starting up the application
// and passing any errors to main
func Run() error {
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database:", err)
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfuly connected and ping the database")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Application run successfully")
}
