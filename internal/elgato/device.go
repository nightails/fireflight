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
