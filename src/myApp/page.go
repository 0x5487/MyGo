package main

import (
	//"encoding/json"
	"github.com/JasonSoft/render"
	//"log"
	//"os"
	"strings"
	"time"
)

type Page struct {
	Id           int    `xorm:"SERIAL index"`
	StoreId      int    `xorm:"INT not null unique(page) index" form:"-" json:"-"`
	TemplateName string `xorm:"not null unique(page) index" binding:"required"`
	Name         string `xorm:"not null unique(page) index" binding:"required"`
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

	if page != nil {
		context := TemplateContext{Page: page, Store: myStore}
		r.HTML(200, page.TemplateName, context)
	}
}

func (page *Page) create() error {
	page.CreatedAt = time.Now().UTC()
	page.UpdatedAt = time.Now().UTC()
	_, err := _engine.Insert(page)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "UNIQUE constraint failed:") {
			myErr := appError{Ex: err, Message: "page name was already existing.", Code: 4001}
			return &myErr
		}
		return err
	} else {
		return nil
	}
}

func getPages(storeId int) *[]Page {
	pages := make([]Page, 0)
	err := _engine.Where("\"StoreId\" = ?", storeId).Find(&pages)
	PanicIf(err)
	return &pages
}
