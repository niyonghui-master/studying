package main

import (
	"fmt"
	"os"
)

func existsPath(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err){
		return false
	}
	return true
}

func createPath(name string) string {
	if !existsPath(name) {
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			fmt.Println("err:", err)
		}
	}

	return name
}

func main() {
	fmt.Println(createPath("aa"))
}
