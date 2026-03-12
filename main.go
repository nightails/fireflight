package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// test IP from .env
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")

	client := &http.Client{Timeout: time.Second * 5}

	resp, err := client.Get(fmt.Sprintf("http://%s:%s", ip, port))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
