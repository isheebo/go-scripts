package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// often times, the images downloaded from twitter are in the format
// .jpg_large. Since very few image viewers can readily read them, this
// rename program comes in handy.

// This program does bulk renaming of files from
// aabb_cc.jpg_large to aabb_cc.jpg

// it walks the file system looking for files whose extensions is
// .jpg_large and simply replaces them with .jpg
//

var (
	// count keeps track of the number of files that have been successfully renamed
	count    = 0
	homePath = os.Getenv("HOME")
	// dirPath  = flag.String("p", "", "path to folder with .jpg_large files")
)

const (
	jpgFormat      = ".jpg"
	jpgLargeFormat = ".jpg_large"
)

func main() {
	dirPath := homePath + "/downloads"

	if err := filepath.Walk(dirPath, renamer); err != nil {
		fmt.Fprintf(os.Stdout, "error: %v\n", err)
	}

	fmt.Printf("%d files have been renamed\n", count)

}

func renamer(path string, info os.FileInfo, err error) error {
	if filepath.Ext(path) == jpgLargeFormat {
		newPath := strings.Join([]string{strings.Trim(path, filepath.Ext(path)), jpgFormat}, "")

		if err := os.Rename(path, newPath); err != nil {
			return err
		}
		count++
	}

	return nil
}
