package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

// exists returns whether the given file or directory exists or not
func isExisting(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getStoreName(storeId int) string {
	return "jason"
}

func logError(message string) {
	log.Println("[Error] " + message)
}

func logInfo(message string) {
	log.Println("[Info] " + message)
}

func logDebug(message string) {
	log.Println("[Debug] " + message)
}

func toString(number int) string {
	return strconv.Itoa(number)
}
