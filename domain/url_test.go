package domain

import (
	"context"
	"fmt"
	"testing"

	"github.com/killinsun/wovn-file-api-cli/mock"
)

type SpyHtmlFileRepository struct{}

func (s *SpyHtmlFileRepository) GetTranslated(ctx context.Context, endpoint, apiConfigQuery, key string) (content string, err error) {
	fmt.Println(apiConfigQuery)
	if apiConfigQuery == "project_token=Zlq6ux&target_lang_code=en&url=https%3A%2F%2Fexample.com%2Fpath%2Fto%2Fdirectory%2Findex.html&url_pattern=path" {
		content = "<h1>It works!</h1>"
	}

	if apiConfigQuery == "project_token=Zlq6ux&target_lang_code=fr&url=https%3A%2F%2Fexample.com%2Fpath%2Fto%2Fdirectory%2Findex.html&url_pattern=path" {
		content = "<h1>Ça marche!</h1>"
	}

	return content, err
}

func (s *SpyHtmlFileRepository) SendReport(endpoint, url, projectToken, apiKey string) (int, error) {
	return 200, nil
}

func TestGetTranslated(t *testing.T) {
	t.Run("should response an english content when given targetLangCode is en", func(t *testing.T) {
		url := "https://example.com/path/to/directory/index.html"
		targetLangCode := "en"
		want := "<h1>It works!</h1>"

		urlStore := NewUrlStore(url)
		spy := &SpyHtmlFileRepository{}

		got, _ := urlStore.GetTranslated(spy, mock.Config, targetLangCode)

		assertCorrectStrings(t, want, got)
	})

	t.Run("should response a french content when given targetLangCode is fr", func(t *testing.T) {
		url := "https://example.com/path/to/directory/index.html"
		targetLangCode := "fr"
		want := "<h1>Ça marche!</h1>"

		urlStore := NewUrlStore(url)
		spy := &SpyHtmlFileRepository{}

		got, _ := urlStore.GetTranslated(spy, mock.Config, targetLangCode)

		assertCorrectStrings(t, want, got)
	})
}

func TestGetFilePath(t *testing.T) {
	srcLangCode := "ja"
	cases := map[string]struct {
		url  string
		lang string
		want string
	}{
		"with / at last 1":                          {"https://example.com/", "en", "./WovnTranslatedHtml_test/en/index.html"},
		"with / at last 2":                          {"https://example.com/path/", "en", "./WovnTranslatedHtml_test/en/path/index.html"},
		"without / at last 1":                       {"https://example.com/path/to", "en", "./WovnTranslatedHtml_test/en/path/to/index.html"},
		"without / at last 2":                       {"https://example.com/path/to/directory", "en", "./WovnTranslatedHtml_test/en/path/to/directory/index.html"},
		"with /index.html at last 1":                {"https://example.com/index.html", "en", "./WovnTranslatedHtml_test/en/index.html"},
		"with /index.html at last 2":                {"https://example.com/path/index.html", "en", "./WovnTranslatedHtml_test/en/path/index.html"},
		"with / at last 1 in french":                {"https://example.com/", "fr", "./WovnTranslatedHtml_test/fr/index.html"},
		"with / at last 2 in french":                {"https://example.com/path/", "fr", "./WovnTranslatedHtml_test/fr/path/index.html"},
		"without / at last 1 in french":             {"https://example.com/path/to", "fr", "./WovnTranslatedHtml_test/fr/path/to/index.html"},
		"without / at last 2 in french":             {"https://example.com/path/to/directory", "fr", "./WovnTranslatedHtml_test/fr/path/to/directory/index.html"},
		"with /index.html at last 1 in french":      {"https://example.com/index.html", "fr", "./WovnTranslatedHtml_test/fr/index.html"},
		"with /index.html at last 2 in french":      {"https://example.com/path/index.html", "fr", "./WovnTranslatedHtml_test/fr/path/index.html"},
		"same as srcLangCode with targetLangCode 1": {"https://example.com/path/index.html", "ja", "./WovnTranslatedHtml_test/path/index.html"},
		"same as srcLangCode with targetLangCode 2": {"https://example.com/index.html", "ja", "./WovnTranslatedHtml_test/index.html"},
		"same as srcLangCode with targetLangCode 3": {"https://example.com/", "ja", "./WovnTranslatedHtml_test/index.html"},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			url := tt.url
			targetLangCode := tt.lang
			want := tt.want

			urlStore := NewUrlStore(url)
			isTest := true
			got, _ := urlStore.GetFilePath(srcLangCode, targetLangCode, isTest)

			if got != want {
				t.Errorf("\n%q \n want: %q, \n got: %q", name, want, got)
			}
		})
	}
}

func TestSendReport(t *testing.T) {
	spy := &SpyHtmlFileRepository{}

	t.Run("Should run repository function correctly.", func(t *testing.T) {
		url := "https://example.com"
		urlStore := NewUrlStore(url)

		got := urlStore.SendReport(spy, mock.Config)

		if got != nil {
			t.Errorf("want: nil, \n got: %q", got)
		}
	})

	t.Run("Should return an error when projectToken is empty.", func(t *testing.T) {
		want := ErrEmptyParams

		url := "https://example.com"
		urlStore := NewUrlStore(url)

		mock.Config.ProjectToken = ""

		got := urlStore.SendReport(spy, mock.Config)

		if got != want {
			t.Errorf("want: %q, \n got: %q", want, got)
		}
	})

	t.Run("Should return an error when API token is empty.", func(t *testing.T) {
		want := ErrEmptyParams

		url := "https://example.com"
		urlStore := NewUrlStore(url)

		mock.Config.APIToken = ""

		got := urlStore.SendReport(spy, mock.Config)

		if got != want {
			t.Errorf("want: %q, \n got: %q", want, got)
		}
	})

	t.Run("Should return an error when url is empty.", func(t *testing.T) {
		want := ErrEmptyParams

		url := ""
		urlStore := NewUrlStore(url)

		got := urlStore.SendReport(spy, mock.Config)

		if got != want {
			t.Errorf("want: %q, \n got: %q", want, got)
		}
	})
}
