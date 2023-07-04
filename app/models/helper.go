package models

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Get(url string) *http.Response {
	// Send a GET request to the endpoint
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Request failed:", err)
		return nil
	}

	return response
}

func Post(url string, contentType string, requestBody io.Reader) *http.Response {
	// Send a GET request to the endpoint
	response, err := http.Post(url, contentType, requestBody)

	if err != nil {
		fmt.Println("Request failed:", err, url, contentType, requestBody)
		return nil
	}

	return response
}

func Put(url string, contentType string, requestBody io.Reader) *http.Response {
	req, err := http.NewRequest("PUT", url, requestBody)

	if err != nil {
		fmt.Println("Request failed:", err)
		return nil
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}

	return response
}

func Del(url string) *http.Response {
	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		fmt.Println("Request failed:", err)
		return nil
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}

	return response
}

func GetBody(response *http.Response, strt any) any {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
	}

	err = json.Unmarshal(body, strt)

	if err != nil {
		fmt.Println("Error:", err)
	}

	return strt
}

func removeEmptyQueries(params url.Values) string {
	filteredParams := url.Values{}

	for key, values := range params {
		for _, value := range values {
			if value != "" {
				filteredParams.Add(key, value)
			}
		}
	}

	return filteredParams.Encode()
}
