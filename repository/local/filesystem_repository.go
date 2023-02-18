package local

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)


type FsRepository struct {}

func (f *FsRepository) Save(path, content string) error {

	f.makeDirAll(path)
	f.putFile(path,content)
	return nil
}

func (f *FsRepository) makeDirAll(path string) error {
	dir, _:= filepath.Split(path)

	os.MkdirAll(dir, 0750)

	return nil
}

func (f *FsRepository) putFile(filePath, content string) {

	bytes := []byte(content)
	err := ioutil.WriteFile(filePath, bytes, 0660 )

	if err != nil {
		fmt.Printf("Could not put a new file. filePath: %q", filePath)
		fmt.Printf("%v",err)
	}
}

func (f *FsRepository) ReadUrlList(path string) (urlList []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
		return urlList, nil
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		urlList = append(urlList, s.Text())
	}

	if s.Err() != nil {
		log.Print(s.Err())
		return urlList, s.Err()
	}

	return urlList, nil
}