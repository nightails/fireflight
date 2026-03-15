package elgato

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// lights contains the array of LED light objects in a device (could has multiple zones)
type lights struct {
	Lights []light `json:"lightsJSON"`
}

// light contains the information of a LED light object
type light struct {
	On          int `json:"on"`
	Brightness  int `json:"brightness"`
	Temperature int `json:"temperature"`
}

func GetLightInfo(url string) (lights, error) {
	resp, err := http.Get(url)
	if err != nil {
		return lights{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return lights{}, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	var lightsArray lights
	if err := json.NewDecoder(resp.Body).Decode(&lightsArray); err != nil {
		return lights{}, fmt.Errorf("failed to decode response body: %v", err)
	}

	return lightsArray, nil
}
