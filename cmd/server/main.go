package main

import (
	"context"
	"fmt"

	"github.com/srcct3/go-rest-comment-api/internal/db"
)

// Run - is responsible for starting up the application
// and passing any errors to main
func Run() error {
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		fmt.Println("failed to ping database")
		return err
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
