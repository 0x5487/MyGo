package main

import (
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
	"fmt"
	"strconv"
	//"time"
	"encoding/json"
)

type ApiOption struct {
	Store *Store
}

func (m *myClassic) UseApi(option ApiOption) error {

	m.Get("/api/v1/collections/:collectionId", func(r render.Render, params martini.Params) {
		collectionId, err := strconv.Atoi(params["collectionId"])

		if err != nil {
			r.JSON(500, "collectionId is invalid")
			return
		}

		collection, err := GetCollection(option.Store.Id, collectionId)

		if err != nil {
			r.JSON(500, "error")
			return
		}

		if collection == nil {
			r.JSON(404, "collection is not found")
			return
		}

		collection.ToJsonForm()
		r.JSON(200, collection)
	})

	m.Get("/api/v1/collections", func(r render.Render) {
		collections, err := GetCollections(option.Store.Id)

		if err != nil {
			r.JSON(500, "error")
			return
		}

		if collections == nil {
			r.JSON(200, "")
			return
		}

		for i := range collections {
			collection := &collections[i]
			collection.ToJsonForm()
		}

		r.JSON(200, collections)
	})

	m.Post("/api/v1/collections", binding.Json(Collection{}), binding.ErrorHandler, func(collection Collection, res http.ResponseWriter) string {
		collection.StoreId = option.Store.Id
		err := collection.create()

		if err != nil {
			if aE, ok := err.(*appError); ok {
				res.WriteHeader(500)
				return aE.Message
			}
		}

		location := fmt.Sprintf("/api/v1/collections/%d", collection.Id)
		res.Header().Add("location", location)
		res.WriteHeader(201)
		return ""
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

	m.Get("/api/v1/products/:productId", func(req *http.Request, r render.Render, params martini.Params) {
		var productId string = params["productId"]

		link := LinkModel{Url: "http://abc"}
		var product = Product{}
		product.All = link

		if productId == "1" {
			links := [2]LinkModel{{Url: "Jaosn"}, {Url: "Hello"}}
			product.All = links
		}

		r.JSON(200, product)
	})

	m.Post("/api/v1/products", binding.Json(Product{}), binding.ErrorHandler, func(product Product, res http.ResponseWriter, r render.Render) {

		println(product.Id)

		//product.All = LinkModel{Url: "abc"}

		/*_, ok := product.All.(LinkModel)
		if ok {
			r.JSON(200, "ok")
		} else {
			r.JSON(200, "failed")
		}*/

		switch v := product.All.(type) {
		case map[string]interface{}:
			println("single")
			j, err := json.Marshal(&product.All)
			if err != nil {
				fmt.Println(err)
			}

			var link LinkModel
			err = json.Unmarshal(j, &link)
			if err != nil {
				fmt.Println(err)
			}
			println(link.Url)
		case []interface{}:
			println("double")

			var links []LinkModel
			ok := MarshalToType(&product.All, &links)
			if ok {

			} else {
				println("Json failed")
			}
		default:
			fmt.Printf("unexpected type %T,", v)
		}
	})

	return nil
}
