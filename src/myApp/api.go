package main

import (
	"github.com/JasonSoft/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	//"github.com/martini-contrib/sessions"
	//"io/ioutil"
	"net/http"
	//"os"
	//"html/template"
	//"path/filepath"
	"fmt"
	"strconv"
	//"time"
	"encoding/json"
)

var aipOption ApiOption

type ApiOption struct {
	Store *Store
}

func (m *myClassic) UseApi(option ApiOption) error {

	aipOption = option

	m.Get("/api/v1/collections/:collectionId", getCollectionHandler)
	m.Put("/api/v1/collections/:collectionId", updateCollectionHandler)
	m.Delete("/api/v1/collections/:collectionId", deleteCollectionHandler)
	m.Get("/api/v1/collections", getCollectionsHandler)
	m.Post("/api/v1/collections", binding.Json(Collection{}), binding.ErrorHandler, createCollectionHandler)

	m.Get("/api/v1/products/:productId", getProductHandler)
	m.Post("/api/v1/products", binding.Json(Product{}), binding.ErrorHandler, createProductHandler)

	m.Get("/api/v1/themes/:themeName", getThemeHandler)
	m.Get("/api/v1/themes", getThemesHandler)
	m.Post("/api/v1/themes", binding.Json(Theme{}), binding.ErrorHandler, createThemeHandler)

	m.Get("/api/v1/pages", getPageHandler)
	m.Post("/api/v1/pages", binding.Json(Page{}), binding.ErrorHandler, createPageHandler)

	m.Get("/api/v1/templates/:templateName", getTemplateHandler)
	m.Get("/api/v1/templates", getTemplatesHandler)
	m.Post("/api/v1/templates", binding.Json(Template{}), binding.ErrorHandler, createTemplateHandler)

	return nil
}

func getCollectionsHandler(r render.Render) {
	collections, err := GetCollections(aipOption.Store.Id)

	if err != nil {
		if appErr, ok := err.(*appError); ok {
			r.JSON(500, appErr.Error())
		} else {
			r.JSON(500, "error")
		}
	}

	if collections == nil {
		r.JSON(200, "")
		return
	}

	for i := range collections {
		collection := &collections[i]
		collection.toJsonForm()
	}

	r.JSON(200, collections)
}

func getCollectionHandler(r render.Render, params martini.Params) {
	collectionId, err := strconv.Atoi(params["collectionId"])

	if err != nil {
		r.JSON(500, "bad request")
		return
	}

	collection, err := GetCollection(aipOption.Store.Id, collectionId)

	if err != nil {
		r.JSON(500, "error")
		return
	}

	if collection == nil {
		r.JSON(404, "collection is not found")
		return
	}

	collection.toJsonForm()
	r.JSON(200, collection)
}

func createCollectionHandler(r render.Render, collection Collection) {
	collection.StoreId = aipOption.Store.Id
	err := collection.create()

	if err != nil {
		if appErr, ok := err.(*appError); ok {
			r.JSON(500, appErr.Error())
			return
		} else {
			logError(err.Error())
			r.JSON(500, "error")
			return
		}
	}

	location := fmt.Sprintf("/api/v1/collections/%d", collection.Id)
	r.Header().Add("location", location)
	r.Status(201)
}

func updateCollectionHandler(r render.Render, collection Collection, params martini.Params) {
	collection.StoreId = aipOption.Store.Id
	err := collection.update()

	if err != nil {
		if appErr, ok := err.(*appError); ok {
			r.JSON(500, appErr.Error())
			return
		} else {
			logError(err.Error())
			r.JSON(500, "error")
			return
		}
	}
	r.Status(200)
}

func deleteCollectionHandler(r render.Render, params martini.Params) {
	collectionId, _ := strconv.Atoi(params["collectionId"])
	collection := Collection{Id: collectionId, StoreId: aipOption.Store.Id}
	err := collection.delete()
	if err != nil {
		logError(err.Error())
		r.JSON(500, "error")
		return
	}
	r.Status(200)
}

func getThemeHandler(r render.Render) {
	themes := getThemes(aipOption.Store.Id)
	r.JSON(200, themes)
}

func getThemesHandler(r render.Render) {
	themes := getThemes(aipOption.Store.Id)
	r.JSON(200, themes)
}

func createThemeHandler(theme Theme, res http.ResponseWriter) string {
	theme.StoreId = aipOption.Store.Id
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
}

func getPageHandler(req *http.Request, r render.Render, params martini.Params) {
	pages := getPages(aipOption.Store.Id)
	r.JSON(200, pages)
}

func createPageHandler(page Page, res http.ResponseWriter) string {
	page.StoreId = aipOption.Store.Id
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
}

func getTemplateHandler(req *http.Request, r render.Render, params martini.Params) {
	themeName := req.URL.Query().Get("theme")

	if len(themeName) <= 0 {
		r.JSON(404, "theme parameter was missing")
		return
	}

	//ensure theme is valid
	targetTheme := aipOption.Store.getTheme(themeName)

	if targetTheme == nil {
		r.JSON(404, "theme was not found")
		return
	}

	templates := []Template{}

	for _, template := range *aipOption.Store.templates {
		if targetTheme.Id == template.ThemeId {
			templates = append(templates, template)
		}
	}

	r.JSON(200, templates)
}

func getTemplatesHandler(req *http.Request, r render.Render, params martini.Params) {
	themeName := req.URL.Query().Get("theme")

	if len(themeName) <= 0 {
		r.JSON(404, "theme parameter was missing")
		return
	}

	//ensure theme is valid
	targetTheme := aipOption.Store.getTheme(themeName)

	if targetTheme == nil {
		r.JSON(404, "theme was not found")
		return
	}

	templates := []Template{}

	for _, template := range *aipOption.Store.templates {
		if targetTheme.Id == template.ThemeId {
			templates = append(templates, template)
		}
	}

	r.JSON(200, templates)
}

func createTemplateHandler(template Template, res http.ResponseWriter) string {

	template.StoreId = aipOption.Store.Id
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
}

func getProductHandler(req *http.Request, r render.Render, params martini.Params) {
	var productId string = params["productId"]

	link := LinkModel{Url: "http://abc"}
	var product = Product{}
	product.All = link

	if productId == "1" {
		links := [2]LinkModel{{Url: "Jaosn"}, {Url: "Hello"}}
		product.All = links
	}

	r.JSON(200, product)
}

func createProductHandler(product Product, res http.ResponseWriter, r render.Render) {

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
}
