package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	dirname := "/Users/martin/Downloads"
	dirEntries, err := os.ReadDir(dirname)
	if err != nil {
		panic(err)
	}

	fileInfos := make([]fs.FileInfo, 0, len(dirEntries))
	for _, entry := range dirEntries {
		info, err := entry.Info()
		if err != nil {
			panic(err)
		}
		fileInfos = append(fileInfos, info)
	}

	for _, fileInfo := range fileInfos {
		filePath := filepath.Join(dirname, fileInfo.Name())
		if fileInfo.ModTime().Day() == 2 &&
			fileInfo.ModTime().Month() == 7 &&
			fileInfo.ModTime().Hour() == 11 &&
			fileInfo.ModTime().Minute() == 13 {
			info, err := os.Stat(filePath)
			if err != nil {
				panic(err)
			}
			if info.IsDir() {
				err := os.RemoveAll(filePath)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
