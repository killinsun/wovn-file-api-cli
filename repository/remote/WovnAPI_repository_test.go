package repository

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetTranslated(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

	var wovnAPIRepository IWovnAPIRepository
	wovnAPIRepository = &WovnAPIRepository{}

	t.Run("Should contain a correct word in got HTTP response body.", func(t *testing.T) {

		want := "<h1>It works!</h1>"

		ctx := context.Background()

		endpoint := srv.URL + "/v1/download_html"
		apiOptionQuery := url.Values{
			"url":              []string{"https://example.com"},
			"project_token":    []string{"hoge"},
			"target_lang_code": []string{"en"},
			"url_pattern":      []string{"path"},
		}

		got, _ := wovnAPIRepository.GetTranslated(ctx, endpoint, apiOptionQuery.Encode(), "key")

		if want != got {
			t.Errorf("want: %q, but got: %q, endpoint: %q, apiOptionQuery: %q ", want, got, endpoint, apiOptionQuery)
		}
	})
}

func TestSendRport(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

	var wovnAPIRepository IWovnAPIRepository
	wovnAPIRepository = &WovnAPIRepository{}

	t.Run("Should get HTTP 200 ", func(t *testing.T) {

		want := http.StatusOK

		endpoint := srv.URL + "/v1/upload_page"
		projectToken := "t3st!"
		apiKey := "AAAA"
		url := "https://example.com"

		got, _ := wovnAPIRepository.SendReport(endpoint, url, projectToken, apiKey)

		if want != got {
			t.Errorf("want: %v, but got: %v, endpoint: %q", want, got, endpoint)
		}
	})

}

func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/download_html", responseDownloadHTMLMock)
	handler.HandleFunc("/v1/upload_page", responseUploadHTMLMock)

	srv := httptest.NewServer(handler)

	return srv
}

func responseDownloadHTMLMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("<h1>It works!</h1>"))
}

func responseUploadHTMLMock(w http.ResponseWriter, r *http.Request) {
	contentLength, _ := strconv.Atoi(r.Header.Get("Content-Length"))

	body := make([]byte, contentLength)
	bodyLength, _ := r.Body.Read(body)

	type RequestBody struct {
		ProjectToken string `json:"project_token"`
		Url          string `json:"url"`
	}
	jsonBody := &RequestBody{}
	json.Unmarshal(body[:bodyLength], &jsonBody)

	if jsonBody.ProjectToken == "" || jsonBody.Url == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("<h1>It works!</h1>"))
}
