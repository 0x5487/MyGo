package main

import (
	//"encoding/json"
	"github.com/JasonSoft/render"
	"io/ioutil"
	"log"
	//"os"
	"path/filepath"
	"time"
)

type Page struct {
	Id           int64
	StoreId      int64  `xorm:"not null unique(page) index" form:"-" json:"-"`
	TemplateName string `xorm:"not null unique(page) index"`
	Name         string `xorm:"not null unique(page) index"`
	Title        string
	Description  string
	Content      string
	CreatedAt    time.Time
	UpdatedAt    time.Time `xorm:"index"`
}

func displayPage(r render.Render, myStore *Store, pageName string) {

	var page *Page

	for _, value := range *myStore.pages {
		if value.Name == pageName {
			page = &value
			break
		}
	}

	r.HTML(200, page.TemplateName, page)
}

func (page *Page) create() {
	_, err := _engine.Insert(page)
	if err != nil {
		panic(err)
	}
}

func getPage(pageName string) string {
	pagePath := filepath.Join(_appDir, "pages", pageName+".html")
	buf, err := ioutil.ReadFile(pagePath)
	if err != nil {
		panic(err)
	}
	return string(buf[:])
}

func getPages(storeId int64) *[]Page {
	log.Println("get pages from database")

	pages := make([]Page, 0)
	err := _engine.Where("StoreId = ?", storeId).Find(&pages)

	if err != nil {
		panic(err)
	}

	return &pages
}
