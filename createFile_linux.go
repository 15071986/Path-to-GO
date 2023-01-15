//go:build linux
// +build linux

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {}

func CreateFile() {

	content, err := ioutil.ReadFile("C:\\Users\\Admin\\LessonGO\\level2\\lesson2\\lesson2.2\\createFile_linux.go")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents: %s", content)
}
