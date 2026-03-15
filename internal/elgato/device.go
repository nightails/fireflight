package elgato

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Device metadata about an Elgato Key Light device. Credit to github.com/mdlayher/keylight
type Device struct {
	ProductName         string `json:"productName,omitempty"`
	HardwareBoardType   int    `json:"hardwareBoardType,omitempty"`
	FirmwareBuildNumber int    `json:"firmwareBuildNumber,omitempty"`
	FirmwareVersion     string `json:"firmwareVersion,omitempty"`
	SerialNumber        string `json:"serialNumber,omitempty"`
	DisplayName         string `json:"displayName,omitempty"`
}

// GetDeviceInfo makes an http get request to the provided URL and return a Device object with it's information
func GetDeviceInfo(url string) (Device, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Device{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Device{}, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	var device Device
	if err := json.NewDecoder(resp.Body).Decode(&device); err != nil {
		return Device{}, fmt.Errorf("failed to decode response body: %v", err)
	}

	return device, nil
}

// PutDeviceInfo makes an http put request to the provided URL
func PutDeviceInfo(url string) error {
	// TODO: implement logic to push new infomation to an elgato device, ex: DisplayName
	return nil
}
