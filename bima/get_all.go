package bima

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"thxrhmn-bimatri-api/config"
	"thxrhmn-bimatri-api/dto"
)

func Get(e config.EndpointInfo, allReq dto.AllRequest) (map[string]interface{}, error) {
	dataRaw := fmt.Sprintf(`
	{
		"accessToken": "%s",
		"callPlan": "%s",
		"language": "%s",
		"msisdn": "%s",
		"secretKey": "%s",
		"subscriberType": "%s"
	}
	`,
		allReq.AccessToken,
		allReq.CallPlan,
		allReq.Language,
		allReq.Msisdn,
		allReq.SecretKey,
		allReq.SubscriberType,
	)

	client := &http.Client{}
	data := strings.NewReader(dataRaw)
	req, err := http.NewRequest(e.Method, e.URL, data)
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range config.CommonHeaders {
		req.Header.Set(key, value)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", allReq.AccessToken))
	req.Header.Set("X-MSISDN", allReq.Msisdn)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthorized: token tidak valid atau sudah kedaluwarsa ")
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf("internal server error: %s", string(bodyText))
	}

	// Parsing JSON ke map[string]interface{}
	var responseData map[string]interface{}
	err = json.Unmarshal(bodyText, &responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func GetRecommendation(e config.EndpointInfo, recReq dto.RecommendationRequest) (map[string]interface{}, error) {
	dataRaw := fmt.Sprintf(`
	{
		"accessToken": "%s",
		"callPlan": "%s",
		"language": "%s",
		"msisdn": "%s",
		"secretKey": "%s",
		"subscriberType": "%s"
	}
	`,
		recReq.AccessToken,
		recReq.CallPlan,
		recReq.Language,
		recReq.Msisdn,
		recReq.SecretKey,
		recReq.SubscriberType,
	)

	client := &http.Client{}
	data := strings.NewReader(dataRaw)
	req, err := http.NewRequest(e.Method, e.URL, data)
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range config.CommonHeaders {
		req.Header.Set(key, value)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", recReq.AccessToken))
	req.Header.Set("X-MSISDN", recReq.Msisdn)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("resp::", resp.StatusCode)
	fmt.Println("body:", string(bodyText))

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthorized: token tidak valid atau sudah kedaluwarsa ")
	}

	if resp.StatusCode == 500 {
		return nil, errors.New("oops something wrong happened")
	}

	// Parsing JSON ke map[string]interface{}
	var responseData map[string]interface{}
	err = json.Unmarshal(bodyText, &responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
