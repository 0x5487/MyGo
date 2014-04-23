package main

import (
	//"log"
	"time"
)

type Template struct {
	Id        int64
	StoreId   int64  `xorm:"not null unique(template)" form:"-" json:"-"`
	ThemeId   int64  `xorm:"not null unique(template)"`
	Name      string `xorm:"not null unique(template)"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (template *Template) create() {
	_, err := _engine.Insert(template)
	if err != nil {
		panic(err)
	}
}

func getTemplates(storeId int64) *[]Template {
	templates := make([]Template, 0)
	err := _engine.Where("StoreId = ?", storeId).Find(&templates)

	if err != nil {
		panic(err)
	}

	return &templates
}
