package main

import (
	"github.com/go-martini/martini"
	"github.com/lunny/xorm"
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

	_engine.Sync(new(HostTable), new(Store), new(User), new(Theme), new(Template), new(Page))

	store := Store{Name: "jason", DefaultTheme: "simple"}
	store.Create()

	hostTable := HostTable{Host: "jason.mystore.com:3000", StoreId: store.Id, StoreName: "jason"}
	hostTable.Create()

}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	m := withoutLogging()
	martini.Env = martini.Dev

	//m.Use(render.Renderer())

	//public folder steup
	publicOption := martini.StaticOptions{SkipLogging: true}
	m.Use(martini.Static("public", publicOption))

	hostTables := make([]HostTable, 0)
	err := engine.Find(&hostTables)
	log.Printf("Host count: %d", len(hostTables))

	stores := make(map[int64]Store)
	_engine.Find(&stores)
	log.Printf("Store count: %d", len(stores))

	jasonStore := NewStore("jason")

	m.Use(func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		if req.Host == "jason.mystore.com:3000" {
			jasonStore.App.ServeHTTP(res, req)
		} else {
			c.Next()
		}
	})

	m.Get("/", func() string {
		return "hello world"
	})

	m.Run()
}
