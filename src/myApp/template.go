package main

import (
	"log"
	"time"
)

type Template struct {
	Id        int64
	StoreId   int64  `xorm:"not null unique(template)" form:"-" json:"-"`
	Name      string `xorm:"not null unique(template)"`
	Content   string `xorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time `xorm:"index"`
}

func (template *Template) create() {
	log.Println("creating template")

	_, err := _engine.Insert(template)
	if err != nil {
		panic(err)
	}
}

func getTemplates(storeId int64) *[]Template {
	log.Println("get templates: ")

	templates := make([]Template, 0)
	err := _engine.Where("StoreId = ?", storeId).Find(&templates)

	if err != nil {
		panic(err)
	}

	return &templates
}
