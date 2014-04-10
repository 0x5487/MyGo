package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type Api struct {
	Version int
}

type myClassic struct {
	*martini.Martini
	martini.Router
}

func displayPage(r render.Render, myStore *Store, pageName string) {
	pagePath := filepath.Join(myStore.StorageRoot, "pages", pageName+".json")
	pageFile, err := os.Open(pagePath)
	if err != nil {
		println("file not found")
		return
	}

	var pageJSON Page
	jsonParser := json.NewDecoder(pageFile)
	if err = jsonParser.Decode(&pageJSON); err != nil {
		println("json file parse error")
		return
	}

	r.HTML(200, pageJSON.Template, pageJSON)
}

func (a Api) Collections() []Collection {
	result := []Collection{{1, "name1", "desc1"}}
	return result
}

func withoutLogging() *myClassic {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Recovery())
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return &myClassic{m, r}
}

func data() string {
	//api := Api{}
	//api.Collection = []collection{{1, "name1", "desc1"}}
	return "abc"
}

func getAllStores() *[]Store {
	stores := []Store{{Id: 1, Name: "Jason"}}
	return &stores
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	m := withoutLogging()
	martini.Env = martini.Dev
	//m.Use(render.Renderer())

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
