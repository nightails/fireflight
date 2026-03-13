package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	url := fmt.Sprintf("http://%s:%s%s", ip, port, lightsURL)
	resp, err := GetLightInfo(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}

// Device metadata about an Elgato Key Light device. Credit to github.com/mdlayher/keylight
type Device struct {
	ProductName         string `json:"productName,omitempty"`
	HardwareBoardType   int    `json:"hardwareBoardType,omitempty"`
	FirmwareBuildNumber int    `json:"firmwareBuildNumber,omitempty"`
	FirmwareVersion     string `json:"firmwareVersion,omitempty"`
	SerialNumber        string `json:"serialNumber,omitempty"`
	DisplayName         string `json:"displayName,omitempty"`
}

func GetDeviceInfo(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	var device Device
	if err := json.NewDecoder(resp.Body).Decode(&device); err != nil {
		return "", fmt.Errorf("failed to decode response body: %v", err)
	}

	return device.SerialNumber, nil
}

type Lights struct {
	Lights []Light `json:"lights"`
}
type Light struct {
	On          int `json:"on"`
	Brightness  int `json:"brightness"`
	Temperature int `json:"temperature"`
}

func GetLightInfo(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	var lights Lights
	if err := json.NewDecoder(resp.Body).Decode(&lights); err != nil {
		return "", fmt.Errorf("failed to decode response body: %v", err)
	}

	return fmt.Sprintf("%v", lights.Lights[0].Brightness), nil
}
