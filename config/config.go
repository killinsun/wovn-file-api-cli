package config

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var Config WovnAPI

type CustomLangAliasesJSON struct {
	Ar    string `json:"ar"`
	My    string `json:"my"`
	Nl    string `json:"nl"`
	Id    string `json:"id"`
	It    string `json:"it"`
	Ms    string `json:"ms"`
	Pt    string `json:"pt"`
	Es    string `json:"es"`
	Th    string `json:"th"`
	Vi    string `json:"vi"`
	Bn    string `json:"bn"`
	ZhCHS string `json:"zh-CHS"`
	De    string `json:"de"`
	Fr    string `json:"fr"`
}
type Options struct {
		CustomLangAliasesJSON CustomLangAliasesJSON `json:"custom_lang_aliases_json"`
		CustomDomainLangsJSON string `json:"custom_domain_langs_json"`
		NoIndexLangsJSON      string `json:"no_index_langs_json"`
		SitePrefixPath        string `json:"site_prefix_path"`
}

type WovnAPI struct {
	UploadApiUri   string `json:"uploadApiUri"`
	DownloadApiUri string `json:"downloadApiUri"`
	ProjectToken   string `json:"projectToken"`
	URLPattern     string `json:"urlPattern"`
	APIToken       string `json:"apiToken"`
	Options 			 Options `json:"options"`
	TargetLangCodes []string `json:"TargetLangCodes"`
	SrcLangCode		string	`json:"srcLangCode"`
}


func (w *WovnAPI) Load(path string) *WovnAPI{
	v := viper.New()
	var (
		_, b, _, _ = runtime.Caller(0)
			basepath   = filepath.Dir(b)
	)

	v.AddConfigPath(basepath)
	v.AddConfigPath("./")
	v.AddConfigPath(path)
	v.SetConfigName("settings")
	v.SetConfigType("json")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("cannot read settings.json file: %w \n", err))
	}

	err = v.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("unmarshal error : %w \n", err))
	}

	return &Config
}