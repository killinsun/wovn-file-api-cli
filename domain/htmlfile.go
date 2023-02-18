package domain

import (
	"log"

	"github.com/killinsun/wovn-file-api-cli/config"
	"github.com/killinsun/wovn-file-api-cli/repository/local"
	repository "github.com/killinsun/wovn-file-api-cli/repository/remote"
)

type HTMLFileStore struct {
	WovnAPIRepository *repository.IWovnAPIRepository
	Config            *config.WovnAPI
	UrlStore          *UrlStore
	TargetLangCode    string
	Path              string
	Content           string
}

func NewHTMLFileStore(
	urlStore *UrlStore,
	wovnAPI repository.IWovnAPIRepository,
	config config.WovnAPI,
	targetLangCode string,
) (htmlFileStore HTMLFileStore, err error) {
	isTest := false
	path, err := urlStore.GetFilePath(config.SrcLangCode, targetLangCode, isTest)
	if err != nil {
		log.Printf("getting file path by URL fail. url: %q, targetLangCode: %q", urlStore.Value, targetLangCode)
		log.Print(err)
		return htmlFileStore, err
	}

	htmlFileStore = HTMLFileStore{
		WovnAPIRepository: &wovnAPI,
		Config:            &config,
		UrlStore:          urlStore,
		TargetLangCode:    targetLangCode,
		Path:              path,
		Content:           "",
	}

	return htmlFileStore, nil
}

func (h *HTMLFileStore) GetTranslatedHTML(ch chan string) {
	content, err := h.UrlStore.GetTranslated(*h.WovnAPIRepository, *h.Config, h.TargetLangCode)

	if err != nil {
		log.Printf("path or content is null. path: %q, content: %q\n", h.Path, content)
		log.Print(err)
	}

	h.Content = content
	ch <- content
}

func (h *HTMLFileStore) Save(f local.IFsRepository, ch chan string) (err error) {
	content := <-ch
	err = f.Save(h.Path, content)

	if err != nil {
		log.Printf("Could not call Save(). %v ", err)
		return err
	}

	return nil

}
