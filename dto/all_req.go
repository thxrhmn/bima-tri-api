package dto

type AllRequest struct {
	AccessToken    string `json:"accessToken"`
	CallPlan       string `json:"callPlan"`
	Language       string `json:"language"`
	Msisdn         string `json:"msisdn"`
	SecretKey      string `json:"secretKey"`
	SubscriberType string `json:"subscriberType"`
}
