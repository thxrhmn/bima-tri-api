package dto

type LoginResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type LoginWithOtpResponse struct {
	AccessToken     string      `json:"accessToken"`
	AppsflyerMsisdn string      `json:"appsflyerMsisdn"`
	Balance         string      `json:"balance"`
	CallPlan        string      `json:"callPlan"`
	CreditLimit     string      `json:"creditLimit"`
	FirstTime       bool        `json:"firstTime"`
	Language        string      `json:"language"`
	Msisdn          string      `json:"msisdn"`
	ProfileColor    string      `json:"profileColor"`
	ProfileTime     int         `json:"profileTime"`
	SecretKey       string      `json:"secretKey"`
	Status          bool        `json:"status"`
	SubscriberType  string      `json:"subscriberType"`
	Xlocation       interface{} `json:"xlocation"`
}

type LoginWithOtpErrorResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
