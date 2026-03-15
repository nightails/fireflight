package elgato

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Lights contains the array of LED light objects in a device (could has multiple zones)
type Lights struct {
	Lights []Light `json:"lightsJSON"`
}

// Light contains the information of a LED light object
type Light struct {
	On          int `json:"on"`
	Brightness  int `json:"brightness"`
	Temperature int `json:"temperature"`
}

// GetLightsInfo makes an http request to the provided URL and return Lights object with information of multiple lights
func GetLightsInfo(url string) (Lights, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Lights{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Lights{}, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	var lights Lights
	if err := json.NewDecoder(resp.Body).Decode(&lights); err != nil {
		return Lights{}, fmt.Errorf("failed to decode response body: %v", err)
	}

	return lights, nil
}
