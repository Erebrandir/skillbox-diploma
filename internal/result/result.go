package result

import (
	"skillbox-diploma/internal/config"
	"skillbox-diploma/internal/status/billing"
	"skillbox-diploma/internal/status/check"
	"skillbox-diploma/internal/status/email"
	"skillbox-diploma/internal/status/incident"
	"skillbox-diploma/internal/status/mms"
	"skillbox-diploma/internal/status/sms"
	"skillbox-diploma/internal/status/support"
	"skillbox-diploma/internal/status/voicecall"
	"sort"
	"time"
)

type ResultT struct {
	Status bool `json:"status"` // true, сбор данных прошел успешно,
	// false в остальных случаях

	Data ResultSetT `json:"data"` // заполнен, сбор данных прошел успешно,
	// nil в остальных случаях

	Error string `json:"error"` // пустая строка если сбор данных прошел успешно,
	// в случае ошибки - текст ошибки
}

type ResultSetT struct {
	SMS       [][]sms.SMSData                `json:"sms"`
	MMS       [][]mms.MMSData                `json:"mms"`
	VoiceCall []voicecall.VoiceCallData      `json:"voice_call"`
	Email     map[string][][]email.EmailData `json:"email"`
	Billing   billing.BillingData            `json:"billing"`
	Support   []int                          `json:"support"`
	Incidents []incident.IncidentData        `json:"incident"`
}

var (
	smsChan       = make(chan [][]sms.SMSData, 0)
	mmsChan       = make(chan [][]mms.MMSData, 0)
	voiceCallChan = make(chan []voicecall.VoiceCallData, 0)
	emailChan     = make(chan map[string][][]email.EmailData, 0)
	billingChan   = make(chan billing.BillingData, 0)
	supportChan   = make(chan []int, 0)
	incidentsChan = make(chan []incident.IncidentData, 0)
)

func getSMSStat() {
	smsDataProviderSort := sms.StatusSMS(config.GlobalConfig.SMSFile)
	smsDataCountrySort := smsDataProviderSort
	smsDataProviderSort = sms.SMSChangeCodeToCountry(smsDataProviderSort)
	smsDataCountrySort = sms.SMSChangeCodeToCountry(smsDataCountrySort)

	sort.SliceStable(smsDataProviderSort, func(i, j int) bool {
		return smsDataProviderSort[i].Provider < smsDataProviderSort[j].Provider
	})

	sort.SliceStable(smsDataCountrySort, func(i, j int) bool {
		return smsDataCountrySort[i].Country < smsDataCountrySort[j].Country
	})

	var result [][]sms.SMSData
	result = append(result, smsDataProviderSort)
	result = append(result, smsDataCountrySort)

	smsChan <- result
}

func getMMSStat() {
	mmsDataProviderSort := mms.StatusMMS(config.GlobalConfig.MMSAddr)
	mmsDataCountrySort := mmsDataProviderSort
	mmsDataProviderSort = mms.MMSChangeCodeToCountry(mmsDataProviderSort)
	mmsDataCountrySort = mms.MMSChangeCodeToCountry(mmsDataCountrySort)

	sort.SliceStable(mmsDataProviderSort, func(i, j int) bool {
		return mmsDataProviderSort[i].Provider < mmsDataProviderSort[j].Provider
	})

	sort.SliceStable(mmsDataCountrySort, func(i, j int) bool {
		return mmsDataCountrySort[i].Country < mmsDataCountrySort[j].Country
	})

	var result [][]mms.MMSData
	result = append(result, mmsDataProviderSort)
	result = append(result, mmsDataCountrySort)
	mmsChan <- result
}

func getVoiceCallStat() {
	result := voicecall.StatusVoiceCall(config.GlobalConfig.VoiceCallFile)
	voiceCallChan <- result
}

func getEmailStat() {
	result := make(map[string][][]email.EmailData, 0)
	data := email.StatusEmail(config.GlobalConfig.EmailFile)
	countries := make(map[string]int)
	for _, elem := range data {
		country := check.GetCountryForCode(elem.Country) //
		elem.Country = country                           //
		countries[elem.Country]++
	}

	for countryCode, _ := range countries {
		var emailDataItem [][]email.EmailData
		emailDataItem = append(emailDataItem, email.Get3MinDeliveryTimeByCountry(data, countryCode))
		emailDataItem = append(emailDataItem, email.Get3MaxDeliveryTimeByCountry(data, countryCode))
		result[countryCode] = emailDataItem
	}

	emailChan <- result
}

func getBillingStat() {
	result := billing.StatusBilling(config.GlobalConfig.BillingFile)
	billingChan <- result
}

func getSupportStat() {
	result := make([]int, 0)
	data := support.StatusSupport(config.GlobalConfig.SupportAddr)
	if len(data) == 0 {
		supportChan <- result
		return
	}

	amountActiveTickets := 0
	for _, elem := range data {
		amountActiveTickets += elem.ActiveTickets
	}

	loading := 0
	switch {
	case amountActiveTickets < 9:
		loading = 1
	case amountActiveTickets >= 9 && amountActiveTickets <= 16:
		loading = 2
	case amountActiveTickets > 16:
		loading = 3
	}
	result = append(result, loading)

	averageTime := amountActiveTickets * 60 / 18
	result = append(result, averageTime)
	supportChan <- result
}

func getIncidentsStat() {
	result := incident.StatusIncident(config.GlobalConfig.IncidentAddr)
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Status < result[j].Status
	})
	incidentsChan <- result
}

func GetResultData() ResultSetT {
	res, ok := GetFromCache()
	if ok {
		return res
	}

	go getSMSStat()
	go getMMSStat()
	go getVoiceCallStat()
	go getEmailStat()
	go getBillingStat()
	go getSupportStat()
	go getIncidentsStat()

	res.SMS = <-smsChan
	res.MMS = <-mmsChan
	res.VoiceCall = <-voiceCallChan
	res.Email = <-emailChan
	res.Billing = <-billingChan
	res.Support = <-supportChan
	res.Incidents = <-incidentsChan

	SetToCache(res, time.Now())
	return res
}

func CheckResult(r ResultSetT) bool {
	if len(r.MMS[0]) == 0 && len(r.MMS[1]) == 0 {
		return false
	}

	if len(r.Support) == 0 {
		return false
	}

	if len(r.Incidents) == 0 {
		return false
	}

	return true
}
