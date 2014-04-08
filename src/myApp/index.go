package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
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
	store := Store{}
	store.Name = name
	store.DefaultTheme = "simple"
	store.DomainNames = []string{domainName}
	store.App = m
	store.StorageRoot = filepath.Join(dir, "storage", name)

	//session setup
	session_store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("my_session", session_store))

	//setup theme
	m.Use(func(res http.ResponseWriter, req *http.Request, c martini.Context, sess sessions.Session) {
		theme := req.URL.Query().Get("theme")

		if len(theme) > 0 {
			sess.Set("theme", theme)
		} else {
			v := sess.Get("theme")
			if v == nil {
				sess.Set("theme", store.DefaultTheme)
			}
		}

		//templates folder setup
		templatesPath := filepath.Join(store.StorageRoot, "themes", sess.Get("theme").(string), "templates")
		renderOption := render.Options{Directory: templatesPath, Extensions: []string{".html"}}
		handler := render.Renderer(renderOption)
		c.Invoke(handler)
		c.Next()
	})

	//files folder setup
	filesPath := filepath.Join(store.StorageRoot, "files")
	filesOption := martini.StaticOptions{Prefix: "/files/"}
	m.Use(martini.Static(filesPath, filesOption))

	//public folder steup
	m.Get("/public/.*", func(res http.ResponseWriter, req *http.Request, c martini.Context, sess sessions.Session) {
		v := sess.Get("theme")
		publicPath := filepath.Join(store.StorageRoot, "themes", v.(string), "public")
		publicOption := martini.StaticOptions{Prefix: "/public"}
		handler := martini.Static(publicPath, publicOption)
		_, err := c.Invoke(handler)
		if err != nil {
			panic(err)
		}
	})

	m.Get("/", func(r render.Render) {
		r.HTML(200, "home", "abc")
	})

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

	m.Get("/", func() string {
		return "hello app"
	})

	m.Run()
}
