package downloads

import (
	"download/apistructures"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gammazero/workerpool"
	"github.com/spf13/afero"
)

func findImagesFromData(fs afero.Fs) map[string]string {
	var images map[string]string = map[string]string{}

	pattern := directories["data"] + "*.json"
	files, _ := afero.Glob(fs, pattern)

	for _, file := range files {
		itemType := strings.ReplaceAll(filepath.Base(file), ".json", "")
		fmt.Printf("  > Loading [%s] from file: [%s]\n", itemType, file)

		content, _ := afero.ReadFile(fs, file)
		items := make([]apistructures.Item, 0)
		// unmarshal the file
		json.Unmarshal(content, &items)
		// loop over each item
		for _, item := range items {
			if len(item.IconURI) > 0 {
				key := fmt.Sprintf("%s/%s/%s", itemType, "thumb", item.FileName)
				images[key] = item.IconURI
			}
			if len(item.ImageURI) > 0 {
				key := fmt.Sprintf("%s/%s/%s", itemType, "main", item.FileName)
				images[key] = item.ImageURI
			}
		}

	}
	return images
}

func Images(fs afero.Fs) error {
	var _error error
	images := findImagesFromData(fs)
	fmt.Printf("Found [%d] images to download\n", len(images))
	wp := workerpool.New(poolSize)
	dir := directories["images"]

	for key, url := range images {
		d := dir
		n := key
		u := url
		wp.Submit(func() {
			filename := d + n + ".png"
			file, e := dl(n, u, fs, filename)
			if e != nil {
				_error = e
				fmt.Printf("[Error] Failed to download [%s]:\n[%s]", u, e)
			} else {
				fmt.Printf("  > [%s] to [%s]\n", u, file)
			}
		})
	}
	wp.StopWait()

	return _error
}
