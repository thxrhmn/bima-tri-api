package bima

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"thxrhmn-bimatri-api/config"
)

func LoginOtpRequest(e config.EndpointInfo, msisdn string) map[string]interface{} {
	dataRaw := fmt.Sprintf(`{"imei":"WebSelfcare", "msisdn": "%s"}`, msisdn)

	endpointInfo := config.EndpointInfo{
		URL:    config.Endpoints.LoginOtpRequestRequest,
		Method: "POST",
	}

	client := &http.Client{}
	data := strings.NewReader(dataRaw)
	req, err := http.NewRequest(endpointInfo.Method, endpointInfo.URL, data)
	if err != nil {
		log.Fatal(err)
	}

	// Set headers
	for key, value := range config.CommonHeaders {
		req.Header.Set(key, value)
	}
	req.Header.Set("X-MSISDN", msisdn)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Parsing JSON ke map[string]interface{}
	var responseData map[string]interface{}
	err = json.Unmarshal(bodyText, &responseData)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
