package main

import (
	"encoding/json"
	"github.com/martini-contrib/render"
	//"log"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Page struct {
	Name        string
	Title       string
	Description string
	Content     string
	Template    string
}

func displayPage(r render.Render, myStore *Store, pageName string) {
	pagePath := filepath.Join(myStore.StorageRoot, "pages", pageName+".json")
	pageFile, err := os.Open(pagePath)
	if err != nil {
		panic(err)
	}

	var pageJSON Page
	jsonParser := json.NewDecoder(pageFile)
	if err = jsonParser.Decode(&pageJSON); err != nil {
		panic(err)
	}

	r.HTML(200, pageJSON.Template, pageJSON)
}

func getPage(pageName string) string {
	pagePath := filepath.Join(_appDir, "pages", pageName+".html")
	buf, err := ioutil.ReadFile(pagePath)
	if err != nil {
		panic(err)
	}
	return string(buf[:])
}
