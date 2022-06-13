package incident

import (
	"encoding/json"
	"fmt"
	"os"
	"skillbox-diploma/internal/config"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // active or closed
}

func StatusIncident(url string) []IncidentData {
	result := make([]IncidentData, 0)

	data, err := os.ReadFile(config.GlobalConfig.IncidentFile)
	if err != nil {
		fmt.Println(err.Error())
		return []IncidentData{}
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return []IncidentData{}
	}

	return result
}
