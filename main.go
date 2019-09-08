package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func fileWalker() error {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return err
			}
			if !strings.Contains(path, "git") {
				newFileName, err := match(path)
				if err != nil {
					fmt.Println(err)
				} else {
					//os.Rename(path, newFileName)
					fmt.Println(newFileName)
				}
			}
			return nil
		})
	return err
}

func main() {
	fmt.Println(fileWalker())
}

func match(fileName string) (string, error) {
	fileNameArr := strings.Split(fileName, ".")
	ext := fileNameArr[len(fileNameArr)-1]
	tmpFileName := strings.Join(fileNameArr[:len(fileNameArr)-1], ".")
	tmpFileNameArr := strings.Split(tmpFileName, "_")
	name := strings.Join(tmpFileNameArr[:len(tmpFileNameArr)-1], "_")
	number, err := strconv.Atoi(tmpFileNameArr[len(tmpFileNameArr)-1])
	if err != nil {
		return "", fmt.Errorf("%s, File Name did not match the pattern", fileName)
	}

	return fmt.Sprintf("%s - %d.%s", strings.Title(name), number, ext), nil
}
