package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	infoURL   = "/elgato/accessory-info"
	lightsURL = "/elgato/lights"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// test IP from .env
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")

	// url example
	_ = fmt.Sprintf("http://%s:%s%s", ip, port, lightsURL)
}
