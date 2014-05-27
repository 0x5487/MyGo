package main

import (
	//"encoding/json"
	"github.com/JasonSoft/render"
	"github.com/go-martini/martini"
	//"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessions"
	//"io/ioutil"
	//"log"
	"net/http"
	//"os"
	"html/template"
	"path/filepath"
	//"strconv"
	//"fmt"
	"time"
)

type Store struct {
	Id               int    `xorm:"SERIAL index"`
	Name             string `xorm:"not null unique"`
	DefaultTheme     string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	storageRoot      string                        `xorm:"-"`
	themes           *[]Theme                      `xorm:"-"`
	templates        *[]Template                   `xorm:"-"`
	pages            *[]Page                       `xorm:"-"`
	templatesService map[string]*template.Template `xorm:"-"`
}

func (store *Store) CreateApp() *myClassic {

	store.storageRoot = filepath.Join(_appDir, "storage", store.Name)
	store.themes = getThemes(store.Id)
	store.templates = getTemplates(store.Id)
	store.pages = getPages(store.Id)
	store.templatesService = map[string]*template.Template{}

	//compile templates
	for _, tmpl := range *store.templates {
		theme := store.getThemeById(tmpl.ThemeId)
		t := store.templatesService[theme.Name]
		if t == nil {
			t = template.New(tmpl.Name)
		}
		template.Must(t.New(tmpl.Name).Parse(tmpl.Content))
		store.templatesService[theme.Name] = t

	}

	m := withoutLogging()

	//session setup
	session_store := sessions.NewCookieStore([]byte("xyz123"))
	m.Use(sessions.Sessions("sid", session_store))

	//setup theme
	m.Use(func(res http.ResponseWriter, req *http.Request, c martini.Context, sess sessions.Session) {
		themeName := req.URL.Query().Get("theme")

		if len(themeName) > 0 {
			//ensure the themeName is valid
			targetTheme := store.getTheme(themeName)
			if targetTheme != nil {
				sess.Set("theme", themeName)
			}
		} else {
			v := sess.Get("theme")
			if v == nil {
				sess.Set("theme", store.DefaultTheme)
			} else {
				themeName = sess.Get("theme").(string)
			}

		}

		/*
			templatesPath := filepath.Join(store.storageRoot, "themes", themeName, "templates")
			renderOption := render.Options{Directory: templatesPath, Extensions: []string{".html"}, IndentJSON: true}
			handler := render.Renderer(renderOption)
			c.Invoke(handler)
			c.Next()
		*/

		renderOption := render.Options{Template: store.templatesService[themeName], IndentJSON: true}
		handler := render.Renderer(renderOption)
		c.Invoke(handler)
		c.Next()
	})

	//files folder setup
	filesPath := filepath.Join(store.storageRoot, "files")
	filesOption := martini.StaticOptions{Prefix: "/files/", SkipLogging: true}
	m.Use(martini.Static(filesPath, filesOption))

	//public folder steup
	m.Get("/public/.*", func(res http.ResponseWriter, req *http.Request, c martini.Context, sess sessions.Session) {
		v := sess.Get("theme")
		publicPath := filepath.Join(store.storageRoot, "themes", v.(string), "public")
		publicOption := martini.StaticOptions{Prefix: "/public", SkipLogging: true}
		handler := martini.Static(publicPath, publicOption)
		_, err := c.Invoke(handler)
		if err != nil {
			panic(err)
		}
	})

	m.Get("/", func(r render.Render) {
		displayPage(r, store, "home")
	})

	m.Get("/admin/main", func() string {
		return displayPrivate("main.html")
	})

	m.Get("/admin", func() string {
		return displayPrivate("index.html")
	})

	m.Get("/pages/:pageName", func(r render.Render, params martini.Params) {
		displayPage(r, store, params["pageName"])
	})

	m.Get("/products/:productName", func(r render.Render, params martini.Params) {
		displayPage(r, store, "product_detail")
	})

	m.Get("/products", func(r render.Render, params martini.Params) {
		displayPage(r, store, "product_list")
	})

	m.Get("/collections/:collectionName", func(r render.Render, params martini.Params) {
		displayPage(r, store, "collection_detail")
	})

	m.Get("/collections", func(r render.Render, params martini.Params) {
		displayPage(r, store, "collection_list")
	})

	m.Get("/cart", func(r render.Render, params martini.Params) {
		displayPage(r, store, "cart")
	})

	m.Get("/session", func(r render.Render, sess sessions.Session) string {
		return sess.Get("theme").(string)
	})

	option := ApiOption{Store: store}
	m.UseApi(option)

	return m
}

func (store *Store) Create() {
	//insert to database
	_, err := _engine.Insert(store)
	if err != nil {
		panic(err)
	}
}

func (store *Store) getTheme(themeName string) *Theme {
	var targetTheme *Theme

	for _, theme := range *store.themes {
		if theme.Name == themeName {
			targetTheme = &theme
			break
		}
	}
	return targetTheme
}

func (store *Store) getThemeById(themeId int) *Theme {
	var targetTheme *Theme

	for _, theme := range *store.themes {
		if theme.Id == themeId {
			targetTheme = &theme
			break
		}
	}
	return targetTheme
}
