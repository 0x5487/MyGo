package main

import (
	//"encoding/json"
	"github.com/JasonSoft/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessions"
	//"io/ioutil"
	"log"
	"net/http"
	//"os"
	"html/template"
	"path/filepath"
	"strconv"
	"time"
)

type Store struct {
	Id               int64
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
	log.Println("create app entity")

	store.storageRoot = filepath.Join(_appDir, "storage", store.Name)
	store.themes = getThemes(store.Id)
	store.templates = getTemplates(store.Id)
	store.pages = getPages(store.Id)
	store.templatesService = map[string]*template.Template{}

	//compile templates
	for _, tmpl := range *store.templates {
		themeKey := "__" + strconv.FormatInt(tmpl.ThemeId, 10)
		t := store.templatesService[themeKey]
		if t == nil {
			t = template.New(tmpl.Name)
		}
		template.Must(t.New(tmpl.Name).Parse(tmpl.Content))
		store.templatesService[themeKey] = t
	}

	m := withoutLogging()

	//session setup
	session_store := sessions.NewCookieStore([]byte("xyz123"))
	m.Use(sessions.Sessions("sid", session_store))

	//setup theme
	m.Use(func(res http.ResponseWriter, req *http.Request, c martini.Context, sess sessions.Session) {
		themeName := req.URL.Query().Get("theme")

		if len(themeName) > 0 {
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

		renderOption := render.Options{Template: store.templatesService["__1"], IndentJSON: true}
		handler := render.Renderer(renderOption)
		c.Invoke(handler)
		c.Next()
	})

	//files folder setup
	filesPath := filepath.Join(store.storageRoot, "files")
	filesOption := martini.StaticOptions{Prefix: "/files/"}
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
		return getPage("main")
	})

	m.Get("/admin", func() string {
		return getPage("index")
	})

	m.Get("/pages/:pageName", func(r render.Render, params martini.Params) {
		displayPage(r, store, params["pageName"])
	})

	m.Get("/api/v1/themes", func(r render.Render) {
		themes := getThemes(store.Id)
		r.JSON(200, themes)
	})

	m.Post("/api/v1/themes", binding.Json(Theme{}), binding.ErrorHandler, func(theme Theme) string {
		log.Println("starting api/v1/themes ")
		theme.StoreId = 5
		theme.Create()

		return theme.Name
	})

	m.Get("/api/v1/pages/", func(req *http.Request, r render.Render, params martini.Params) {

		log.Println("starting api/pages")

		pages := getPages(store.Id)
		r.JSON(200, pages)

		log.Println("finished api/pages")
	})

	m.Get("/api/v1/templates/", func(req *http.Request, r render.Render, params martini.Params) {
		println("calling /api/v1/templates/")

		themeName := req.URL.Query().Get("theme")

		if len(themeName) <= 0 {
			r.JSON(404, "theme parameter is missing")
		}

		//ensure theme is valid
		targetTheme := store.getTheme(themeName)

		if targetTheme == nil {
			r.JSON(404, "theme is not found")
			return
		}

		templates := []Template{}

		for _, template := range *store.templates {
			if targetTheme.Id == template.ThemeId {
				templates = append(templates, template)
			}
		}

		r.JSON(200, templates)

	})

	return m
}

func (store *Store) Create() {
	//insert to database
	_, err := _engine.Insert(store)
	if err != nil {
		panic(err)
	}
}

func (hostTable *HostTable) Create() {
	//insert to database
	_, err := _engine.Insert(hostTable)
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
