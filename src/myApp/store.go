package main

import (
	//"encoding/json"
	"github.com/JasonSoft/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessions"
	//"io/ioutil"
	//"log"
	"net/http"
	//"os"
	"html/template"
	"path/filepath"
	//"strconv"
	"fmt"
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
		return displayShare("main.html")
	})

	m.Get("/admin", func() string {
		return displayShare("index.html")
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

	m.Get("/api/v1/collections", func(r render.Render) {
		themes := getThemes(store.Id)
		r.JSON(200, themes)
	})

	m.Get("/api/v1/themes/:themeName", func(r render.Render) {
		themes := getThemes(store.Id)
		r.JSON(200, themes)
	})

	m.Get("/api/v1/themes", func(r render.Render) {
		themes := getThemes(store.Id)
		r.JSON(200, themes)
	})

	m.Post("/api/v1/themes", binding.Json(Theme{}), binding.ErrorHandler, func(theme Theme, res http.ResponseWriter) string {
		theme.StoreId = store.Id
		err := theme.create()

		if err != nil {
			if aE, ok := err.(*appError); ok {
				res.WriteHeader(500)
				return aE.Message
			}
		}

		location := fmt.Sprintf("/api/v1/themes/%d", theme.Id)
		res.Header().Add("location", location)
		res.WriteHeader(201)
		return ""
	})

	m.Get("/api/v1/pages", func(req *http.Request, r render.Render, params martini.Params) {
		pages := getPages(store.Id)
		r.JSON(200, pages)
	})

	m.Post("/api/v1/pages", binding.Json(Page{}), binding.ErrorHandler, func(page Page, res http.ResponseWriter) string {
		page.StoreId = store.Id
		err := page.create()

		if err != nil {
			if aE, ok := err.(*appError); ok {
				res.WriteHeader(500)
				return aE.Message
			}
		}

		location := fmt.Sprintf("/api/v1/pages/%d", page.Id)
		res.Header().Add("location", location)
		res.WriteHeader(201)
		return ""
	})

	m.Get("/api/v1/templates/:templateName", func(req *http.Request, r render.Render, params martini.Params) {
		themeName := req.URL.Query().Get("theme")

		if len(themeName) <= 0 {
			r.JSON(404, "theme parameter was missing")
			return
		}

		//ensure theme is valid
		targetTheme := store.getTheme(themeName)

		if targetTheme == nil {
			r.JSON(404, "theme was not found")
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

	m.Get("/api/v1/templates", func(req *http.Request, r render.Render, params martini.Params) {
		themeName := req.URL.Query().Get("theme")

		if len(themeName) <= 0 {
			r.JSON(404, "theme parameter was missing")
			return
		}

		//ensure theme is valid
		targetTheme := store.getTheme(themeName)

		if targetTheme == nil {
			r.JSON(404, "theme was not found")
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

	m.Post("/api/v1/templates", binding.Json(Template{}), binding.ErrorHandler, func(template Template, res http.ResponseWriter) string {

		template.StoreId = store.Id
		err := template.create()

		if err != nil {
			if aE, ok := err.(*appError); ok {
				res.WriteHeader(500)
				return aE.Message
			}
		}

		location := fmt.Sprintf("/api/v1/templates/%d?theme", template.Id)
		res.Header().Add("location", location)
		res.WriteHeader(201)
		return ""
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
