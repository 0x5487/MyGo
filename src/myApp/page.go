package main

import (
	"encoding/json"
	"github.com/martini-contrib/render"
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
		println("file not found")
		return
	}

	var pageJSON Page
	jsonParser := json.NewDecoder(pageFile)
	if err = jsonParser.Decode(&pageJSON); err != nil {
		println("json file parse error")
		return
	}

	r.HTML(200, pageJSON.Template, pageJSON)
}
