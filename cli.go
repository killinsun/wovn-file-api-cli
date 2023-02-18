package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/killinsun/wovn-file-api-cli/service"
)

const (
	ExitCodeOK         = 0
	ExitCodeParseError = 1
	ExitCodeFlagError  = 2
)

var versionOutput = fmt.Sprintf("WOVN File translation tool version %s\n", Version)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	var showVersion, downloadHTMLFiles, sendReport bool
	var configPath string
	flags := flag.NewFlagSet("wovn-ft", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&showVersion, "version", false, "Print version information and quit")
	flags.BoolVar(&downloadHTMLFiles, "d", false, "Download Translated HTML Files by download_html api")
	flags.StringVar(&configPath, "c", "./settings.json", "Settings")
	flags.BoolVar(&sendReport, "u", false, "Send report by upload_api")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseError
	}

	if downloadHTMLFiles && sendReport {
		fmt.Println("You need to use either -d and -u flag.")
		return ExitCodeFlagError
	}

	if showVersion {
		return c.ShowVersion()
	}

	if downloadHTMLFiles {
		urlListPath := "./url.txt"

		appService := service.AppService{}
		return c.DownloadHTMLFiles(&appService, configPath, urlListPath)
	}

	if sendReport {
		urlListPath := "./url.txt"

		appService := service.AppService{}
		return c.SendReport(&appService, configPath, urlListPath)
	}

	return ExitCodeOK
}

func (c *CLI) ShowVersion() int {
	fmt.Fprintf(c.outStream, versionOutput)
	return ExitCodeOK
}

func (c *CLI) DownloadHTMLFiles(appService service.IApplicationService, configPath string, urlListPath string) int {
	appService.DownloadHTMLFiles(configPath, urlListPath)
	return ExitCodeOK
}

func (c *CLI) SendReport(appService service.IApplicationService, configPath string, urlListPath string) int {
	appService.SendReport(configPath, urlListPath)
	return ExitCodeOK
}
