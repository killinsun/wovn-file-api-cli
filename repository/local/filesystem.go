package local

type IFsRepository interface {
	Save(path, content string) error
	ReadUrlList(path string) (urlList []string, err error)
}