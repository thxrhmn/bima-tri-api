package config

import "github.com/spf13/viper"

type EndpointInfo struct {
	URL    string
	Method string
}

type BimaCreds struct {
	AccessToken string
	SecretKey   string
	Msisdn      string
}

var CommonHeaders map[string]string

type EndpointURLs struct {
	ProfileAccount             string
	LoginOtpRequestRequest     string
	LoginWithOtp               string
	Promo                      string
	HistoryPaket               string
	Recommendation             string
	Notification               string
	TransactionHistoryPulsa    string
	TransactionHistoryPaket    string
	TransactionHistoryTransfer string
	TransactionHistoryHiburan  string
	TransactionHistoryGames    string
}

var Endpoints EndpointURLs

func Init() {
	viper.AutomaticEnv()

	baseURL := viper.GetString("BASE_URL")

	CommonHeaders = map[string]string{
		"Accept":                       "application/json, text/plain, */*",
		"Accept-Language":              "en-US,en;q=0.9",
		"Access-Control-Allow-Headers": "Content-Type",
		"Access-Control-Allow-Methods": "GET,POST,OPTIONS",
		"Access-Control-Allow-Origin":  "*",
		"Connection":                   "keep-alive",
		"Content-Type":                 "application/json",
		"Host":                         baseURL,
		"Origin":                       "https://" + baseURL,
		"Referer":                      "https://" + baseURL + "/",
		"Sec-Fetch-Dest":               "empty",
		"Sec-Fetch-Mode":               "cors",
		"Sec-Fetch-Site":               "same-origin",
		"User-Agent":                   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",
		"app-version":                  "selfcare",
		"sec-ch-ua":                    `"Chromium";v="123", "Not:A-Brand";v="8"`,
		"sec-ch-ua-mobile":             "?0",
		"sec-ch-ua-platform":           `"Windows"`,
	}

	Endpoints = EndpointURLs{
		ProfileAccount:             viper.GetString("PROFILE_ACCOUNT"),
		LoginOtpRequestRequest:     viper.GetString("LOGIN_OTP_REQUEST"),
		LoginWithOtp:               viper.GetString("LOGIN_WITH_OTP"),
		Promo:                      viper.GetString("PROMO"),
		Recommendation:             viper.GetString("RECOMMENDATION"),
		Notification:               viper.GetString("NOTIFICATION"),
		TransactionHistoryPulsa:    viper.GetString("TRANSACTION_HISTORY_PULSA"),
		TransactionHistoryPaket:    viper.GetString("TRANSACTION_HISTORY_PAKET"),
		TransactionHistoryTransfer: viper.GetString("TRANSACTION_HISTORY_TRANSFER"),
		TransactionHistoryHiburan:  viper.GetString("TRANSACTION_HISTORY_HIBURAN"),
		TransactionHistoryGames:    viper.GetString("TRANSACTION_HISTORY_GAMES"),
	}
}
