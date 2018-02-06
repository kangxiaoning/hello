package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// get the directory of the currently running file
	// method 1
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	// method 2, as of go 1.8
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	// method 3, old version
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	// get parent directory of dir
	parentPath1, err := filepath.Abs(filepath.Join(pwd,".."))
	if err != nil {
		fmt.Println(err)
	}
	parentPath2, err := filepath.Abs(filepath.Join(pwd,"../../.."))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parentPath1)
	fmt.Println(parentPath2)
}
