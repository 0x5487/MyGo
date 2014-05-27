package main

import (
	//"encoding/json"
	"github.com/JasonSoft/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	//"github.com/martini-contrib/sessions"
	//"io/ioutil"
	//"log"
	"net/http"
	//"os"
	//"html/template"
	//"path/filepath"
	//"strconv"
	"fmt"
	//"time"
)

type ApiOption struct {
	Store *Store
}

func (m *myClassic) UseApi(option ApiOption) error {
	m.Get("/api/v1/collections", func(r render.Render) {
		themes := getThemes(option.Store.Id)
		r.JSON(200, themes)
	})

	m.Get("/api/v1/themes/:themeName", func(r render.Render) {
		themes := getThemes(option.Store.Id)
		r.JSON(200, themes)
	})

	m.Get("/api/v1/themes", func(r render.Render) {
		themes := getThemes(option.Store.Id)
		r.JSON(200, themes)
	})

	m.Post("/api/v1/themes", binding.Json(Theme{}), binding.ErrorHandler, func(theme Theme, res http.ResponseWriter) string {
		theme.StoreId = option.Store.Id
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
		pages := getPages(option.Store.Id)
		r.JSON(200, pages)
	})

	m.Post("/api/v1/pages", binding.Json(Page{}), binding.ErrorHandler, func(page Page, res http.ResponseWriter) string {
		page.StoreId = option.Store.Id
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
		targetTheme := option.Store.getTheme(themeName)

		if targetTheme == nil {
			r.JSON(404, "theme was not found")
			return
		}

		templates := []Template{}

		for _, template := range *option.Store.templates {
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
		targetTheme := option.Store.getTheme(themeName)

		if targetTheme == nil {
			r.JSON(404, "theme was not found")
			return
		}

		templates := []Template{}

		for _, template := range *option.Store.templates {
			if targetTheme.Id == template.ThemeId {
				templates = append(templates, template)
			}
		}

		r.JSON(200, templates)
	})

	m.Post("/api/v1/templates", binding.Json(Template{}), binding.ErrorHandler, func(template Template, res http.ResponseWriter) string {

		template.StoreId = option.Store.Id
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

	return nil
}
