package downloads

import (
	"fmt"

	"github.com/gammazero/workerpool"
	"github.com/spf13/afero"
)

func Data(fs afero.Fs) error {
	var _error error
	wp := workerpool.New(poolSize)
	dir := directories["data"]

	for name, url := range API {
		d := dir
		n := name
		u := url
		wp.Submit(func() {
			filename := d + n + ".json"
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
