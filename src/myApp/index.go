package main

import (
	"github.com/go-martini/martini"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var _appDir string
var _engine *xorm.Engine
var _hostApp map[string]*myClassic
var _defaultDatabaseTime string = "1900-1-1"

func init() {
	//begin defining  global variables
	var err error
	_appDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	//end defining global variables

	_engine, err = xorm.NewEngine("postgres", "host=localhost dbname=ecStore user=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	_engine.ShowSQL = true
	_engine.TZLocation = time.UTC
	_engine.SetMapper(core.SameMapper{})
	err = _engine.DropTables(new(IdGenertator), new(Host), new(Store), new(User), new(Theme), new(Template), new(Page), new(Image), new(Collection), new(Product), new(CustomField), new(Variation), new(collection_product), new(image_any))
	PanicIf(err)

	err = _engine.Sync(new(IdGenertator), new(Host), new(Store), new(User), new(Theme), new(Template), new(Page), new(Image), new(Collection), new(Product), new(CustomField), new(Variation), new(collection_product), new(image_any))
	PanicIf(err)

	createFakeData()
	getHostApp()
	getHostMappings()

	runtime.GOMAXPROCS(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
}

func main() {
	m := withoutLogging()
	martini.Env = martini.Dev

	//m.Use(render.Renderer())

	//public folder steup
	publicOption := martini.StaticOptions{SkipLogging: true}
	m.Use(martini.Static("public", publicOption))

	//find store app instance
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
