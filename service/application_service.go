package service

type IApplicationService interface {
	DownloadHTMLFiles(configPath, urlListPath string)
	SendReport(configPath, urlListPath string)
}

type AppService struct{} 