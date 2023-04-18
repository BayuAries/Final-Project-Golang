package main

import (
	"MyGram/database"
	"MyGram/router"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// load env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	database.StarDB()
	r := router.StarApp()

	port := os.Getenv("APP_PORT")
	r.Run(fmt.Sprintf(":%s", port))
}
