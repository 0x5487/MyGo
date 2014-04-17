package main

import (
	"github.com/go-martini/martini"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"path/filepath"
	//"runtime"
)

type myClassic struct {
	*martini.Martini
	martini.Router
}

func withoutLogging() *myClassic {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Recovery())
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return &myClassic{m, r}
}

var _appDir string
var _engine *xorm.Engine

func init() {
	//runtime.GOMAXPROCS(runtime.NumCPU())

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

	_engine, err = xorm.NewEngine("sqlite3", "./database/test.db")
	//defer _engine.Close()
	_engine.SetMapper(xorm.SameMapper{})
	_engine.Sync(new(HostTable), new(Store), new(User), new(Theme), new(Template), new(Page))
	createFakeData()
}

func main() {
	m := withoutLogging()
	martini.Env = martini.Dev

	//m.Use(render.Renderer())

	//public folder steup
	publicOption := martini.StaticOptions{SkipLogging: true}
	m.Use(martini.Static("public", publicOption))

	m.Use(func(res http.ResponseWriter, req *http.Request, c martini.Context) {

		var isHostMatch = false

		for key, value := range getHostApp() {
			if req.Host == key {
				log.Println("Matched: " + req.Host)
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
