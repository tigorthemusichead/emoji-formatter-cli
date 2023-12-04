package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"t3011/scanner"
)

func GetRootByUser() (string, error) {
	fmt.Println("Enter the path to a root folder")
	var path string
	var _ int
	var err error
	var i byte = 0
	for _, err = fmt.Scanf("%s", &path); err != nil && i < 3; {
		fmt.Println("Error while entering, try again")
		_, err = fmt.Scanf("%s", &path)
		i++
	}
	if i < 3 {
		return path, nil
	} else {
		return "", errors.New("too many attempts to specify the root directory")
	}

}

func GetRootByPWD() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	} else {
		return path, nil
	}
}

var ignoreList = []string{".git", ".gitignore"}

func main() {
	var manualMode bool
	var reversedMode bool
	flag.BoolVar(&manualMode, "manual", false, "Set root directory manually")
	flag.BoolVar(&reversedMode, "reverse", false, "Format the code back to readable")
	flag.Parse()
	var path string
	var err error
	if manualMode {
		path, err = GetRootByUser()
	} else {
		path, err = GetRootByPWD()
	}
	if err != nil {
		return
	}
	fmt.Println("Path: ", path)
	err = scanner.Scan(path, ignoreList, reversedMode)
}
