package mock

import "github.com/killinsun/wovn-file-api-cli/config"

var Config = config.WovnAPI{
	UploadApiUri:   "https://api.wovn.io/v1/upload_page",
	DownloadApiUri: "https://api.wovn.io/v1/download_html",
	ProjectToken:   "Zlq6ux",
	URLPattern:     "path",
	APIToken:       "eyJhbGciOiJIUzI1NiJ9.eyJwcm9qZWN0X3Rva2VuIjoiWmxxNnV4IiwidG9rZW5fdXVpZCI6IjQyMGMwOGU0LTFlZWUtNDJhYS05NDJmLTQ0NWU1Njg1M2EwMyJ9.y0DgphMTvUwgCRR-7iJGpZ8QmlPwV7hdULjqp8p7hP4",
	Options: config.Options{
		CustomLangAliasesJSON: config.CustomLangAliasesJSON{
			Ar:    "ara",
			My:    "mya",
			Nl:    "dut",
			Id:    "ind",
			It:    "ita",
			Ms:    "msa",
			Pt:    "por",
			Es:    "spa",
			Th:    "tha",
			Vi:    "vie",
			Bn:    "ben",
			ZhCHS: "zho",
			De:    "deu",
			Fr:    "fra",
		},
	},
	TargetLangCodes: []string{"ja", "en"},
	SrcLangCode:     "ja",
}
