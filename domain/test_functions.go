package domain

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func assertCorrectStrings(t *testing.T, want, got string) {
	t.Helper()

	if got != want {
		t.Errorf("want: %q, \n got: %q", want, got)
	}
}

type mockWovnAPIRepository struct {} 

func (m *mockWovnAPIRepository) GetTranslated(ctx context.Context, endpoint, apiOptionQuery, apiKey string) (content string, err error) {
	if apiOptionQuery == "project_token=Zlq6ux&target_lang_code=en&url=https%3A%2F%2Fexample.com%2F&url_pattern=path" {
		content = "<h1>It works!</h1>"
	}

	if apiOptionQuery == "project_token=Zlq6ux&target_lang_code=fr&url=https%3A%2F%2Fexample.com%2F&url_pattern=path" {
		content = "<h1>Ça marche!</h1>"
	}

	if content == "" {
		content = "<h1>It works!</h1>"
	}
	return content, err
}

func (m *mockWovnAPIRepository) SendReport(endpoint, url, projectToken, apiKey string) (int, error) {
	return 200, nil
}

type SpyFsRepository struct {}

func (s *SpyFsRepository) Save(path, content string) error {
	fmt.Printf("Save() path: %q, content: %q \n", path, content)

	if path == "" {
		return errors.New(ErrSaveFile.Error())
	}

	if content == "" {
		return errors.New(ErrSaveFile.Error())
	}

	return nil
}

func (s *SpyFsRepository) ReadUrlList(path string) (urlList []string, err error) {
	urlList = []string {
		"https://example.com",
		"https://example.com/path",
		"https://example.com/path/to",
		"https://example.com/path/to/directory",
	}
	if path == "" {
		err = errors.New("path is an empty")
		return nil, err
	}
	return urlList, err
}