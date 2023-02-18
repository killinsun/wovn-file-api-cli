package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type WovnAPIRepository struct {}

func (w *WovnAPIRepository) GetTranslated(ctx context.Context, endpoint, apiOptionQuery, apiKey string) (content string, err error) {

	req, err := http.NewRequest("GET", endpoint + "?" + apiOptionQuery, nil)
	req.WithContext(ctx)
	req.Header.Set("Authorization", "X-Wovn-Key: " + apiKey)

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("cannot get a request")
	}

	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("cannot read a HTML file")
	}
	content = string(byteArray)

	return	content, nil

}

func (w *WovnAPIRepository) SendReport(endpoint, url, projectToken, apiKey string) (status int, err error) {
	type RequestBody struct {
		ProjectToken string `json:"project_token"`
		Url string `json:"url"`
	}
	requestBody := &RequestBody{
		ProjectToken: projectToken,
		Url: url,
	}
	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(jsonBody))
	req.Header.Set("Authorization", "X-Wovn-Key: " + apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return http.StatusBadRequest, errors.New("cannot get a request")
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Could not read response body. url: %v", url)
	}

	if resp.StatusCode != 200 {
		log.Printf("url: %v, projectToken: %v\n",url, projectToken)
		log.Printf("message: \n %v", string(bodyBytes))
	}

	return	resp.StatusCode, nil
}