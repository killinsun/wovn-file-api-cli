package domain

import (
	"context"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/killinsun/wovn-file-api-cli/config"
	repository "github.com/killinsun/wovn-file-api-cli/repository/remote"
)

type UrlStore struct {
	Value string
}

func NewUrlStore(givenUrl string) *UrlStore {
	u := new(UrlStore)
	u.Value = givenUrl

	return u
}

func (u *UrlStore) GetTranslated(h repository.IWovnAPIRepository, config config.WovnAPI, taregetLangCode string) (content string, err error) {
	ctx := context.Background()

	endpoint := config.DownloadApiUri
	apiOptionQuery := url.Values{
		"url":              []string{u.Value},
		"project_token":    []string{config.ProjectToken},
		"target_lang_code": []string{taregetLangCode},
		"url_pattern":      []string{config.URLPattern},
	}

	content, err = h.GetTranslated(ctx, endpoint, apiOptionQuery.Encode(), config.APIToken)

	return content, err
}

func (u *UrlStore) GetRootDirectory(isTest bool) (rootDir string) {
	if isTest {
		rootDir = "./WovnTranslatedHtml_test"
	} else {
		rootDir = "./WovnTranslatedHtml"
	}

	return rootDir
}

func (u *UrlStore) GetFilePath(srcLangCode, targetLangCode string, isTest bool) (filePath string, err error) {

	parsedUrl, err := url.Parse(u.Value)
	if err != nil {
		return "", ErrUrlParse
	}

	pathSlice := strings.Split(parsedUrl.Path, "/")
	trailingSlashReg, err := regexp.Compile(`.*\/$`)
	if err != nil {
		return "", ErrRegexpCompile
	}

	hasFileName, err := regexp.Compile(`.*\/index.html$`)
	if err != nil {
		return "", ErrRegexpCompile
	}

	var filePathHead string
	if srcLangCode == targetLangCode {
		filePathHead = u.GetRootDirectory(isTest)
	} else {
		filePathHead = u.GetRootDirectory(isTest) + "/" + targetLangCode
	}

	var filePathBody string
	if trailingSlashReg.MatchString(parsedUrl.Path) {
		// has a slash at last
		filePathBody = strings.Join(pathSlice[:len(pathSlice)-1], "/") + "/index.html"
	} else {
		// doesn't have a slash at last
		if hasFileName.MatchString(parsedUrl.Path) {
			// has a index.html at last
			filePathBody = strings.Join(pathSlice, "/")
		} else {
			// doesn't have a index.html at last
			filePathBody = strings.Join(pathSlice, "/") + "/index.html"
		}
	}

	filePath = filePathHead + filePathBody

	return filePath, nil
}

func (u *UrlStore) SendReport(h repository.IWovnAPIRepository, config config.WovnAPI) error {
	if u.Value == "" || config.ProjectToken == "" || config.APIToken == "" {
		return ErrEmptyParams
	}
	endpoint := config.UploadApiUri
	result, err := h.SendReport(endpoint, u.Value, config.ProjectToken, config.APIToken)

	if result != 200 {
		log.Printf("Responsed HTTP status is not 200. %v", result)
	}

	return err
}
