package elgato

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// lightsJSON contains the array of single LED lights in a device
type lightsJSON struct {
	Lights []lightJSON `json:"lightsJSON"`
}

// lightJSON contains the information of a single LED light
type lightJSON struct {
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

	var lights lightsJSON
	if err := json.NewDecoder(resp.Body).Decode(&lights); err != nil {
		return "", fmt.Errorf("failed to decode response body: %v", err)
	}

	return fmt.Sprintf("%v", lights.Lights[0].Brightness), nil
}
