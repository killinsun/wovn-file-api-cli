package service

import (
	"fmt"

	"github.com/killinsun/wovn-file-api-cli/config"
	"github.com/killinsun/wovn-file-api-cli/domain"
	"github.com/killinsun/wovn-file-api-cli/repository/local"
	repository "github.com/killinsun/wovn-file-api-cli/repository/remote"
)

func (a *AppService) SendReport(configPath, urlListPath string) {
	fsRepository := local.FsRepository{}
	wovnApiRepository := repository.WovnAPIRepository{}

	urlListStore := domain.NewUrlListStore(&fsRepository, urlListPath)
	err := urlListStore.SendReportByURLList(&wovnApiRepository, *config.Config.Load(configPath))
	if err != nil {
		fmt.Println(err)
		panic("error...")
	}
}
