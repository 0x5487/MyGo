package main

import (
	"log"
	"time"
)

type Template struct {
	Id        int64
	StoreId   int64     `xorm:"not null" form:"-" json:"-"`
	ThemeId   int64     `xorm:"not null" form:"-" json:"-"`
	Name      string    `xorm:"not null unique index"`
	Content   string    `xorm:"-"`
	CreatedAt time.Time `xorm:"index"`
	UpdatedAt time.Time `xorm:"index"`
}

func (template *Template) create() {
	log.Println("create template")

	//insert to database
	_, err := _engine.Insert(template)
	if err != nil {
		panic(err)
	}

	//create theme folder
}

func getTemplates(themeName string) *[]Template {
	log.Println("get templates: " + themeName)

	templates := make([]Template, 0)
	err := _engine.Find(&templates)

	if err != nil {
		panic(err)
	}

	return &templates
}
