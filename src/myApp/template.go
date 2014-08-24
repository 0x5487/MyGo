package main

import (
	//"log"
	"strings"
	"time"
)

type Template struct {
	Id        int    `xorm:"SERIAL index"`
	StoreId   int    `xorm:"INT not null unique(template)" form:"-" json:"-"`
	ThemeId   int    `xorm:"INT not null unique(template)" binding:"required"`
	Name      string `xorm:"not null unique(template)" binding:"required"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TemplateContext struct {
	Page  *Page
	Store *Store
}

func (context TemplateContext) Collections() []Collection {
	collections, _ := GetCollections(context.Store.Id)
	return collections
}

func (template *Template) create() error {
	template.CreatedAt = time.Now().UTC()
	template.UpdatedAt = time.Now().UTC()
	_, err := _engine.Insert(template)

	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "UNIQUE constraint failed:") {
			myErr := appError{Ex: err, Message: "template name was already existing.", Code: 4001}
			return &myErr
		}
		return err
	} else {
		return nil
	}
}

func getTemplates(storeId int) *[]Template {
	templates := make([]Template, 0)
	err := _engine.Where("\"StoreId\" = ?", storeId).Find(&templates)

	PanicIf(err)

	return &templates
}
