package downloads

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

var API = map[string]string{
	"fish": "https://acnhapi.com/v1a/fish",
	"bugs": "https://acnhapi.com/v1a/bugs",
	"sea":  "https://acnhapi.com/v1a/sea",
	// "fossils":   "https://acnhapi.com/v1a/fossils",
	// "villagers": "https://acnhapi.com/v1a/villagers",
	// "art":       "https://acnhapi.com/v1a/art",
}

var poolSize = 20
var directories = map[string]string{
	"data":   "./_source/data/",
	"images": "./_source/images/",
}

func dir(fs afero.Fs, filename string) {
	fileDir := filepath.Dir(filename)
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		fs.MkdirAll(fileDir, os.ModePerm)
	}
}

func dl(name string, url string, fs afero.Fs, filename string) (string, error) {
	// make the call
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// create the dir
	dir(fs, filename)
	// get the body of the call
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return "", e
	}
	// return file and error if failed to save
	return filename, afero.WriteFile(fs, filename, body, os.FileMode(int(0777)))
}
