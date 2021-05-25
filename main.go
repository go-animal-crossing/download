package main

import (
	"download/downloads"
	"fmt"

	"github.com/spf13/afero"
)

func main() {
	fmt.Printf("Starting download..")
	fs := afero.NewOsFs()

	downloads.Data(fs)
	downloads.Images(fs)
	fmt.Printf("Ending download..")
}
