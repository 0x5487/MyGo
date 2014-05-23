package main

type User struct {
	Id    int32
	Name  string `xorm:"varchar(25) not null unique 'usr_name'"`
	Email string
}
