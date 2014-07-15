package main

import (
	"github.com/go-martini/martini"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

var _appDir string
var _engine *xorm.Engine
var _hostApp map[string]*myClassic

func init() {
	//define global variables
	var err error
	_appDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	databasePath := "./database/test.db"
	if _, err := os.Stat(databasePath); err == nil {
		os.Remove(databasePath)
		log.Println("database file was removed")
	}
	//define global variables *

	_engine, err = xorm.NewEngine("sqlite3", "./database/test.db")
	// ToDo: we need to close the database connection
	//defer _engine.Close()
	_engine.ShowSQL = false
	_engine.SetMapper(core.SameMapper{})
	_engine.Sync(new(Host), new(Store), new(User), new(Theme), new(Template), new(Page), new(Image), new(Collection), new(Product), new(CustomField), new(Variation), new(collection_product), new(image_any))

	createFakeData()
	getHostApp()
	getHostMappings()

	runtime.GOMAXPROCS(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
}

func main() {
	m := withoutLogging()
	martini.Env = martini.Prod

	//m.Use(render.Renderer())

	//public folder steup
	publicOption := martini.StaticOptions{SkipLogging: true}
	m.Use(martini.Static("public", publicOption))

	m.Use(func(res http.ResponseWriter, req *http.Request, c martini.Context) {

		var isHostMatch = false

		for key, value := range getHostApp() {
			if req.Host == key {
				//log.Println("Matched: " + req.Host)
				isHostMatch = true
				value.ServeHTTP(res, req)
				break
			}
		}

		if isHostMatch == false {
			c.Next()
		}
	})

	m.Get("/", func() string {
		return "Hello Root!!"
	})

	m.Run()
}
