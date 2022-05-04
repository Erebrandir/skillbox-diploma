package sms

import (
	"bufio"
	"fmt"
	"os"
	"skillbox-diploma/pkg/check"
	"strings"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func StatusSMS(csvFile string) []SMSData {
	smsData := make([]SMSData, 0)

	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err.Error() + `: ` + csvFile)
		return []SMSData{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		data := strings.Split(line, ";")

		if len(data) != 4 {
			continue
		}

		if check.CheckCountry(data[0]) && check.CheckBandwidth(data[1]) && check.CheckResponseTime(data[2]) && check.CheckProvider(data[3]) {
			elem := SMSData{Country: data[0], Bandwidth: data[1], ResponseTime: data[2], Provider: data[3]}
			smsData = append(smsData, elem)
		}
	}

	return smsData
}
