package local

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestSave(t *testing.T) {
	path := "../../WovnTranslatedHtml_test/index.html" 
	want := "<html><body><h1>Hi!</h1></body></html>"


	repository := FsRepository{}
	repository.Save(path, want)

	bytes, _ := ioutil.ReadFile(path)

	if want != string(bytes) {
		t.Errorf("File was not put correctly. \n want: \n %q \n get: %q\n",want, string(bytes))
	}

	err := os.Remove(path)
	if err != nil {
		t.Errorf("File remove error. \n %q", err)
	}
}

func TestReadUrlList(t * testing.T){
	path := "./url_test.txt"
	want := []string{
		"https://example.com",
		"https://example.com/path",
		"https://example.com/path/to",
	}

	
	bytes := []byte(strings.Join(want, "\n"))
	err := ioutil.WriteFile(path, bytes, 0660 )
	if err != nil {
		t.Errorf("Create mock file error.")
	}

	repository := FsRepository{}
	got, _ := repository.ReadUrlList(path)

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("want: \n %q \n get: %q\n", want[i], got[i])
		}
	}

	err = os.Remove(path)
	if err != nil {
		t.Errorf("File remove error. \n %q", err)
	}
}