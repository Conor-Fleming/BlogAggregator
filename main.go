package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

}
