// This script is to used to generate information about file name and size
package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	fileInfos := []map[string]interface{}{}

	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory", err)
		return
	}

	for _, file := range files {
		fileInfo := map[string]interface{}{
			"name": file.Name(),
			"size": getFileSize(file),
		}
		fileInfos = append(fileInfos, fileInfo)
	}

	fmt.Println(fileInfos)
}

func getFileSize(file fs.DirEntry) int64 {
	info, err := file.Info()
	if err != nil {
		fmt.Println("Error getting file info", err)
		return 0
	}
	return info.Size()
}
