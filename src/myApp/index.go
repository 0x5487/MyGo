package main

import (
	"github.com/go-martini/martini"
	//"github.com/martini-contrib/render"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	//"strings"
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

type Store struct {
	Id           int
	Name         string
	DefaultTheme string
	StorageRoot  string
	DomainNames  []string
	App          *myClassic
}

func NewStore(name string) *Store {
	domainName := name + ".mystore.com"
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	m := withoutLogging()
	m.Get("/", func() {
		println("hello" + name)
	})

	store := Store{}
	store.Name = name
	store.DefaultTheme = "simple"
	store.DomainNames = []string{domainName}
	store.App = m
	store.StorageRoot = filepath.Join(dir, name)

	return &store
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

	m.Get("/", func() {
		println("hello app")
	})

	m.Run()
}
