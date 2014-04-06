package main

import {
	"github.com/go-martini/martini"
	"fmt"
}


func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	
	m.Run()
}
