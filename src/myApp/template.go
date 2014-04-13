package main

type Template struct {
	Id      int64
	Name    string
	Content string
	StoreId int64 `xorm:"not null unique" form:"-" json:"-"`
}
