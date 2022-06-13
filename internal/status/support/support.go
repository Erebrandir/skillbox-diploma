package support

import (
	"encoding/json"
	"fmt"
	"os"
	"skillbox-diploma/internal/config"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func StatusSupport(url string) []SupportData {
	result := make([]SupportData, 0)

	data, err := os.ReadFile(config.GlobalConfig.SupportFile)
	if err != nil {
		fmt.Println(err.Error())
		return []SupportData{}
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return []SupportData{}
	}

	return result
}
