package main

import (
	"bytes"
	"strings"
	"testing"
)

type mockAppService struct{}

func (m *mockAppService) DownloadHTMLFiles(configPath, urlListPath string) {}

func (m *mockAppService) SendReport(configPath, urlListPath string) {}

func TestRun(t *testing.T) {
	t.Run("Should response ExitCodeOK and print correctly when passed version option.", func(t *testing.T) {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		args := strings.Split("./wovn-file-api-cli -version", " ")
		want := versionOutput

		cli := &CLI{outStream: outStream, errStream: errStream}

		status := cli.Run(args)

		if status != ExitCodeOK {
			t.Errorf("want %d, but got %d", ExitCodeOK, status)
		}
		if !strings.Contains(outStream.String(), want) {
			t.Errorf("want %q, but got %q", want, outStream.String())
		}
	})
}

func TestShowVersion(t *testing.T) {
	t.Run("Should response ExitCodeOK and print correctly when passed version option.", func(t *testing.T) {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		want := versionOutput

		cli := &CLI{outStream: outStream, errStream: errStream}
		status := cli.ShowVersion()

		if status != ExitCodeOK {
			t.Errorf("want %d, but got %d", ExitCodeOK, status)
		}
		if !strings.Contains(outStream.String(), want) {
			t.Errorf("want %q, but got %q", want, outStream.String())
		}
	})
}

func TestDownloadHTMLFiles(t *testing.T) {
	t.Run("Should response ExitCodeOK.", func(t *testing.T) {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		want := ExitCodeOK

		cli := &CLI{outStream: outStream, errStream: errStream}

		mock := mockAppService{}
		configPath := "hoge"
		urlListPath := "fuga"
		got := cli.DownloadHTMLFiles(&mock, configPath, urlListPath)

		if want != got {
			t.Errorf("want %q, but got %q", want, outStream.String())
		}
	})
}

func TestSendReport(t *testing.T) {
	t.Run("Should response ExitCodeOK.", func(t *testing.T) {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		want := ExitCodeOK

		cli := &CLI{outStream: outStream, errStream: errStream}

		mock := mockAppService{}
		configPath := "hoge"
		urlListPath := "fuga"
		got := cli.SendReport(&mock, configPath, urlListPath)

		if want != got {
			t.Errorf("want %q, but got %q", want, outStream.String())
		}
	})
}
