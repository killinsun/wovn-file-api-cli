package domain

import (
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/killinsun/wovn-file-api-cli/mock"
	"github.com/killinsun/wovn-file-api-cli/utils"
)

const concurrency = 17

func TestNewHTMLStore(t *testing.T) {
	t.Run("Should response HTMLstore in English.", func(t *testing.T) {
		wovnAPIRepository := &mockWovnAPIRepository{}
		targetLangCode := "en"

		want := HTMLFileStore{
			Path:           "./WovnTranslatedHtml/en/index.html",
			TargetLangCode: targetLangCode,
		}

		url := "https://example.com/"
		urlStore := NewUrlStore(url)

		config := mock.Config

		got, _ := NewHTMLFileStore(urlStore, wovnAPIRepository, config, targetLangCode)
		assertCorrectStrings(t, want.Path, got.Path)
		assertCorrectStrings(t, want.TargetLangCode, got.TargetLangCode)
	})

	t.Run("Should response HTMLstore in French.", func(t *testing.T) {
		wovnAPIRepository := &mockWovnAPIRepository{}
		targetLangCode := "fr"

		want := HTMLFileStore{
			Path:           "./WovnTranslatedHtml/fr/index.html",
			TargetLangCode: targetLangCode,
		}

		url := "https://example.com/"
		urlStore := NewUrlStore(url)

		config := mock.Config

		got, _ := NewHTMLFileStore(urlStore, wovnAPIRepository, config, targetLangCode)
		assertCorrectStrings(t, want.Path, got.Path)
		assertCorrectStrings(t, want.TargetLangCode, got.TargetLangCode)
	})
}

func TestGetTranslatedHTML(t *testing.T) {
	t.Run("Should get HTML data into store", func(t *testing.T) {
		wovnAPIRepository := &mockWovnAPIRepository{}

		want := HTMLFileStore{
			Path:    "./WovnTranslatedHtml/fr/index.html",
			Content: "<h1>Ça marche!</h1>",
		}

		url := "https://example.com/"
		urlStore := NewUrlStore(url)

		config := mock.Config
		targetLangCode := "fr"

		var wg sync.WaitGroup
		var counter = 1
		wg.Add(counter)
		htmlFileStore, _ := NewHTMLFileStore(urlStore, wovnAPIRepository, config, targetLangCode)
		ch := make(chan string)
		go htmlFileStore.GetTranslatedHTML(ch)
		go func() { <-ch }()

		wg.Wait()

		assertCorrectStrings(t, want.Content, htmlFileStore.Content)

	})
}

func BenchmarkGetTranslatedHTML(b *testing.B) {
	wovnAPIRepository := &mockWovnAPIRepository{}

	config := mock.Config
	targetLangCode := "en"

	b.Run("groutine test", func(b *testing.B) {
		var counter = 1000

		sem := make(chan struct{}, concurrency)
		ticker := time.NewTicker(1 * time.Second)
		go utils.ClearChannelInterval(concurrency, ticker, &sem)

		var wg sync.WaitGroup
		for i := 0; i < counter; i++ {
			sem <- struct{}{}
			url := "https://example.com/" + strconv.Itoa(i)
			urlStore := NewUrlStore(url)
			wg.Add(1)
			htmlFileStore, _ := NewHTMLFileStore(urlStore, wovnAPIRepository, config, targetLangCode)
			ch := make(chan string)
			go htmlFileStore.GetTranslatedHTML(ch)
			go func() { <-ch }()
		}
		wg.Wait()

	})
}

func TestSave(t *testing.T) {

	t.Run("Call Save() correctly", func(t *testing.T) {
		spy := &SpyFsRepository{}
		html := HTMLFileStore{
			UrlStore:       NewUrlStore("https://example.com"),
			TargetLangCode: "en",
			Path:           "../WovnTranslatedHtml_test/path/to/directory/index.html",
			Content:        "<h1>It works!</h1>",
		}

		ch := make(chan string)
		go func() { ch <- html.Content }()
		err := html.Save(spy, ch)

		if err != nil {
			t.Errorf("Save() was not called collectly.")
		}
	})

	t.Run("Call Save() returns ErrSaveFile error when filePath is blank.", func(t *testing.T) {
		spy := &SpyFsRepository{}
		html := HTMLFileStore{
			UrlStore:       NewUrlStore("https://example.com"),
			TargetLangCode: "en",
			Path:           "",
			Content:        "<h1>It works!</h1>",
		}

		ch := make(chan string)
		go func() { ch <- html.Content }()
		err := html.Save(spy, ch)

		if err == nil {
			t.Errorf("Save() was not called collectly.")
		}
	})

	t.Run("Call Save() returns ErrSaveFile error when content is blank.", func(t *testing.T) {
		spy := &SpyFsRepository{}
		html := HTMLFileStore{
			UrlStore:       NewUrlStore("https://example.com"),
			TargetLangCode: "en",
			Path:           "../WovnTranslatedHtml_test/index.html",
			Content:        "",
		}

		ch := make(chan string)
		go func() { ch <- html.Content }()
		err := html.Save(spy, ch)

		if err == nil {
			t.Errorf("Save() was not called collectly.")
		}
	})
}
