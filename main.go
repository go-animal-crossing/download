package main

import (
	"download/downloads"

	"github.com/spf13/afero"
)

func main() {

	fs := afero.NewOsFs()

	downloads.Data(fs)
	downloads.Images(fs)
}
