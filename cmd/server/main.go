package main

import "fmt"

// Run - is responsible for starting up the application
// and passing any errors to main
func Run() error {
	fmt.Println("Instantiation success")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Application run successfully")
}
