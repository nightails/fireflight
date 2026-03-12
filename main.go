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

	url := fmt.Sprintf("http://%s:%s/elgato/accessory-info", ip, port)
	resp, err := GetLightInfo(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}

func GetLightInfo(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// TODO: parse response body

	return resp.Status, nil
}
