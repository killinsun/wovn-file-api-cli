package domain

import (
	"log"
	"time"

	"github.com/killinsun/wovn-file-api-cli/config"
	"github.com/killinsun/wovn-file-api-cli/repository/local"
	repository "github.com/killinsun/wovn-file-api-cli/repository/remote"
	"github.com/killinsun/wovn-file-api-cli/utils"
)

type UrlListStore struct {
	UrlStore []UrlStore
}

func NewUrlListStore(localRepo local.IFsRepository, path string) *UrlListStore {

	urlListStore := &UrlListStore{
		UrlStore: []UrlStore{},
	}

	urlList, err := localRepo.ReadUrlList(path)
	if err != nil {
		log.Print("Read URL list error", err)
	}

	for _, v := range urlList {
		urlListStore.UrlStore = append(urlListStore.UrlStore, *NewUrlStore(v))
	}

	return urlListStore
}

func (u *UrlListStore) GetAllTranslatedHTML(
	localRepo local.IFsRepository,
	remoteRepo repository.IWovnAPIRepository,
	config config.WovnAPI,
) error {

	const concurrency = 4

	sem := make(chan struct{}, concurrency)
	defer close(sem)
	ticker := time.NewTicker(1500 * time.Millisecond)

	go utils.ClearChannelInterval(concurrency, ticker, &sem)

	for _, targetLangCode := range config.TargetLangCodes {
		for _, urlStore := range u.UrlStore {
			sem <- struct{}{}

			var store = HTMLFileStore{}
			store, err := NewHTMLFileStore(&urlStore, remoteRepo, config, targetLangCode)
			if err != nil {
				return err
			}
			ch := make(chan string)
			defer close(ch)
			go store.GetTranslatedHTML(ch)
			go store.Save(localRepo, ch)

			log.Println(store.Path)
		}
	}

	return nil
}

func (u *UrlListStore) SendReportByURLList(
	remoteRepo repository.IWovnAPIRepository,
	config config.WovnAPI,
) error {

	const concurrency = 4

	sem := make(chan struct{}, concurrency)
	defer close(sem)
	ticker := time.NewTicker(1500 * time.Millisecond)

	go utils.ClearChannelInterval(concurrency, ticker, &sem)
	for _, urlStore := range u.UrlStore {
		sem <- struct{}{}

		var store = HTMLFileStore{}
		store, err := NewHTMLFileStore(&urlStore, remoteRepo, config, config.SrcLangCode)
		if err != nil {
			return err
		}
		go urlStore.SendReport(remoteRepo, config)
		log.Println(store.Path)
	}

	return nil
}
