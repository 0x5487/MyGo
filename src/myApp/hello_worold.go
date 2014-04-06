package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"runtime"
)

type myClassic struct {
	*martini.Martini
	martini.Router
}

type Collection struct {
	Id          int
	Name        string
	Description string
}

type Api struct {
	Version int
}

func (a Api) Collections() []Collection {
	result := []Collection{{1, "name1", "desc1"}}
	return result
}

func withoutLogging() *myClassic {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Recovery())
	m.Use(martini.Static("public"))
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return &myClassic{m, r}
}

func data() string {
	//api := Api{}
	//api.Collection = []collection{{1, "name1", "desc1"}}
	return "abc"
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	m := withoutLogging()
	martini.Env = martini.Dev
	m.Use(render.Renderer())

	api := Api{Version: 1}
	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello", api)
	})
	m.Run()
}
