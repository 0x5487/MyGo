package main

import (
	"encoding/json"
	"io/ioutil"
	//"log"
	//"os"
	"fmt"
	"github.com/go-martini/martini"
	"path/filepath"
)

type appError struct {
	Ex      error
	Message string
	Code    int
}

type User struct {
	Id    int32
	Name  string `xorm:"varchar(25) not null unique 'usr_name'"`
	Email string
}

type myClassic struct {
	*martini.Martini
	martini.Router
}

func (e *appError) Error() string { return e.Message }

func displayPrivate(fileName string) string {
	filePath := filepath.Join(_appDir, "private", fileName)
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(buf[:])
}

func withoutLogging() *myClassic {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Recovery())
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return &myClassic{m, r}
}

func MarshalToType(source interface{}, target interface{}) bool {
	var result bool = false

	j, err := json.Marshal(source)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(j, target)
	if err != nil {
		fmt.Println(err)
	}
	result = true
	return result
}

func AppendIfMissing(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
