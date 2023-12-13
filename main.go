package main

import "os"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := contollers.NewApplication(database.Product)
}
