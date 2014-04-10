package main

import (
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
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

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	_appDir = dir
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	m := withoutLogging()
	martini.Env = martini.Dev

	//m.Use(render.Renderer())

	//public folder steup
	publicOption := martini.StaticOptions{SkipLogging: true}
	m.Use(martini.Static("public", publicOption))

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
