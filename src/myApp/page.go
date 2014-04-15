package main

import (
	"encoding/json"
	"github.com/martini-contrib/render"
	//"log"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Page struct {
	Id          int64
	StoreId     int64  `xorm:"not null unique(page) index" form:"-" json:"-"`
	TemplateId  int64  `xorm:"not null unique(page) index"`
	Name        string `xorm:"not null unique(page) index"`
	Title       string
	Description string
	Content     string `xorm:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time `xorm:"index"`
}

func displayPage(r render.Render, myStore *Store, pageName string) {
	pagePath := filepath.Join(myStore.storageRoot, "pages", pageName+".json")
	pageFile, err := os.Open(pagePath)
	if err != nil {
		panic(err)
	}

	var pageJSON Page
	jsonParser := json.NewDecoder(pageFile)
	if err = jsonParser.Decode(&pageJSON); err != nil {
		panic(err)
	}

	var page *Page

	for _, value := range *myStore.pages {
		if value.Name == pageName {
			page = &value
			break
		}
	}

	var templateName string

	for _, template := range *myStore.templates {
		if template.Id == page.TemplateId {
			templateName = template.Name
			break
		}
	}

	if len(templateName) <= 0 {
		panic("can't find template: " + templateName)
	}

	r.HTML(200, templateName, pageJSON)
}

func getPage(pageName string) string {
	pagePath := filepath.Join(_appDir, "pages", pageName+".html")
	buf, err := ioutil.ReadFile(pagePath)
	if err != nil {
		panic(err)
	}
	return string(buf[:])
}
