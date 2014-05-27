package main

import (
	//"encoding/json"
	"io/ioutil"
	//"log"
	//"os"
	"path/filepath"
)

type appError struct {
	Ex      error
	Message string
	Code    int
}

func (e *appError) Error() string { return e.Message }

func displayPrivate(fileName string) string {
	filePath := filepath.Join(_appDir, "private", fileName)
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(buf[:])
}
