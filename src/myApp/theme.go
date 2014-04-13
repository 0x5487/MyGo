package main

import (
	"github.com/martini-contrib/binding"
	"log"
	"net/http"
	"time"
)

type Theme struct {
	Id        int64  `xorm:"index"`
	Name      string `xorm:"varchar(25) not null unique"`
	IsDefault bool
	TimeStamp string    `form:"-" json:"-"`
	Created   time.Time `xorm:"index"`
	StoreId   int64     `xorm:"not null unique" form:"-" json:"-"`
}

func (theme Theme) Validate(errors *binding.Errors, req *http.Request) {

	if len(theme.Name) < 4 {
		errors.Fields["Name"] = "Too short; minimum 4 characters"
	} else if len(theme.Name) > 120 {
		errors.Fields["Name"] = "Too long; maximum 120 characters"
	}
}

func (theme *Theme) Create() {
	log.Println("create theme")

	//insert to database
	_, err := _engine.Insert(theme)
	if err != nil {
		println(err.Error())
	}

	//create theme folder
}
