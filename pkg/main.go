package main

import (
	"fmt"
	"os"
)

func findFileDfs(reqFile string, currentFile string) (resPaths []string) {

	dirFiles, err := os.ReadDir(currentFile)
	if err != nil {
		return resPaths
	}

	for _, file := range dirFiles {
		if reqFile == file.Name() {
			resPaths = append(resPaths, currentFile+"/"+file.Name())
		} else if file.IsDir() {
			resPaths = append(resPaths, findFileDfs(reqFile, currentFile+"/"+file.Name())...)
		}
	}

	return resPaths
}

func main() {
	arg := os.Args
	arg = arg[1:]

	ans := findFileDfs(arg[0], ".")

	for _, i := range ans {
		fmt.Println(i)
	}
}
