package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"io/ioutil"
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

type Page struct {
	Name        string
	Title       string
	Description string
	Content     string
	Template    string
}

type Store struct {
	Id           int
	Name         string
	DefaultTheme string
	StorageRoot  string
	DomainNames  []string
	App          *myClassic
}

type Theme struct {
	Id   int
	Name string
}

type Template struct {
	Id      int
	Name    string
	Content string
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
		renderOption := render.Options{Directory: templatesPath, Extensions: []string{".html"}, IndentJSON: true}
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
		publicOption := martini.StaticOptions{Prefix: "/public", SkipLogging: true}
		handler := martini.Static(publicPath, publicOption)
		_, err := c.Invoke(handler)
		if err != nil {
			panic(err)
		}
	})

	m.Get("/", func(r render.Render) {
		displayPage(r, &store, "home")
	})

	m.Get("/api/v1/themes/:themeName/", func(r render.Render, params martini.Params) {

	})

	m.Get("/api/v1/templates/", func(req *http.Request, r render.Render, params martini.Params) {
		theme := req.URL.Query().Get("theme")
		println("theme:" + theme)

		if len(theme) > 0 {
			templatesDir := filepath.Join(store.StorageRoot, "themes", theme, "templates")
			templates := []Template{}

			filepath.Walk(templatesDir, func(path string, fileInfo os.FileInfo, err error) error {

				ext := filepath.Ext(path)
				if ext == ".html" {
					buf, err := ioutil.ReadFile(path)
					if err != nil {
						panic(err)
					}

					content := string(buf[:])
					name := fileInfo.Name()

					template := Template{
						Name:    name,
						Content: content,
					}

					println(name)
					println(content)
					templates = append(templates, template)
				}

				return nil
			})

			r.JSON(200, templates)
		}
	})

	m.Get("/api/v1/pages/", func(req *http.Request, r render.Render, params martini.Params) {

		println("starting api/pages")

		pagesDir := filepath.Join(store.StorageRoot, "pages")
		pages := []Page{}

		filepath.Walk(pagesDir, func(path string, fileInfo os.FileInfo, err error) error {

			ext := filepath.Ext(path)
			if ext == ".json" {
				buf, err := ioutil.ReadFile(path)
				if err != nil {
					panic(err)
				}

				var page Page
				if err = json.Unmarshal(buf, &page); err != nil {
					println("json file parse error")
					return err
				}

				page.Name = fileInfo.Name()
				pages = append(pages, page)
			}

			return nil
		})

		r.JSON(200, pages)

		println("finished api/pages")
	})

	m.Get("/pages/:pageName", func(r render.Render, params martini.Params) {
		displayPage(r, &store, params["pageName"])
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
