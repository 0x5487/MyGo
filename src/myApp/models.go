package main

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"fmt"
	"github.com/JasonSoft/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	NoTrack = iota
	Track
	External
)

type ManageInventoryMethod int
type Money int64

type LinkModel struct {
	Url string
}

type Host struct {
	Id      int `xorm:"PK SERIAL index"`
	StoreId int `xorm:"INT index"`
	Name    string
}

type Image struct {
	Path       string
	Url        string
	Position   int
	FileName   string
	Attachment string
}

type CustomField struct {
	Name  string
	Value string
}

// FileInfo describes a file that has been uploaded.
type FileInfo struct {
	Key          string `json:"-"`
	Url          string `json:"url,omitempty"`
	ThumbnailUrl string `json:"thumbnail_url,omitempty"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Size         int64  `json:"size"`
	Error        string `json:"error,omitempty"`
	DeleteUrl    string `json:"delete_url,omitempty"`
	DeleteType   string `json:"delete_type,omitempty"`
}

type IdGenertator struct {
	Id   int `xorm:"PK SERIAL"`
	Name string
}

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

type Collection struct {
	Id             int    `xorm:"INT index"`
	StoreId        int    `xorm:"INT not null unique(resourceId) unique(name)" form:"-" json:"-"`
	ResourceId     string `xorm:"not null unique(resourceId)"`
	Path           string
	DisplayName    string `xorm:"not null unique(name)"`
	IsVisible      bool
	Content        string
	Image          *Image        `xorm:"-"`
	ImageDB        string        `json:"-"`
	Tags           []string      `xorm:"-"`
	TagsDB         string        `json:"-"`
	ProductIds     []int         `xorm:"-"`
	CustomFieldsDB string        `json:"-"`
	CustomFields   []CustomField `xorm:"-"`
	CreatedAt      time.Time     `xorm:"created" json:"-"`
	UpdatedAt      time.Time     `xorm:"updated index" json:"-"`
	DeletedAt      time.Time     `xorm:"unique(name) unique(resourceId)" json:"-"`
}

type Product struct {
	Id                        int    `xorm:"PK SERIAL index"`
	StoreId                   int    `xorm:"INT not null unique(resourceId) unique(name) unique(Sku)" form:"-" json:"-"`
	Sku                       string `xorm:"not null unique(Sku)"`
	SkuEx                     string `xorm:"not null unique(Sku)" form:"-" json:"-"`
	ResourceId                string `xorm:"not null unique(resourceId)"`
	Name                      string `xorm:"not null unique(name)"`
	IsPurchasable             bool
	IsVisible                 bool
	IsBackOrderEnabled        bool
	IsPreOrderEnabled         bool
	IsShippingAddressRequired bool
	Tags                      string `xorm:"not null"`
	ListPrice                 Money  `xorm:"INT index"`
	Price                     Money  `xorm:"INT index"`
	Content                   string `xorm:"not null"`
	Vendor                    string `xorm:"not null"`
	InventoryQuantity         int    `xorm:"INT"`
	Weight                    int
	ManageInventoryMethod     ManageInventoryMethod
	OptionSetId               int         `xorm:"INT"`
	PageTitle                 string      `xorm:"not null"`
	MetaDescription           string      `xorm:"not null"`
	Variations                interface{} `xorm:"-"`
	Images                    interface{} `xorm:"-"`
	CustomFields              interface{} `xorm:"-"`
	All                       interface{} `xorm:"-"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time `xorm:"index"`
	DeletedAt                 time.Time
}

type Variation struct {
	Id                        int `xorm:"PK SERIAL index"`
	StoreId                   int `xorm:"INT not null index" form:"-" json:"-"`
	Sku                       string
	DisplayName               string
	IsPurchasable             bool
	IsVisible                 bool
	IsBackOrderEnabled        bool
	IsPreOrderEnabled         bool
	IsShippingAddressRequired bool
	Tags                      string
	ListPrice                 Money `xorm:"INT index"`
	Price                     Money `xorm:"INT index"`
	Description               string
	Vendor                    string
	InventoryQuantity         int `xorm:"INT"`
	ManageInventoryMethod     ManageInventoryMethod
	Weight                    int
	CreatedAt                 time.Time
	UpdatedAt                 time.Time `xorm:"index"`
	DeletedAt                 time.Time
}

type OptionSet struct {
	Id        int
	Name      string
	Options   LinkModel `xorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time `xorm:"index"`
}

type OptionSetOption struct {
	Id           int
	OptionSetId  int
	OptionId     int
	Position     int
	IsRequired   bool
	Option       Option        `xorm:"-"`
	OptionValues []OptionValue `xorm:"-"`
	CreatedAt    time.Time
	UpdatedAt    time.Time `xorm:"index"`
}

type Option struct {
	Id          int
	Name        string
	DisplayName string
	Values      LinkModel `xorm:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time `xorm:"index"`
}

type OptionValue struct {
	Id        int
	OptionId  int
	Position  int
	Lable     string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time `xorm:"index"`
}

type collection_product struct {
	Id           int `xorm:"PK SERIAL index"`
	CollectionId int
	ProductId    int
	CreatedAt    time.Time
	UpdatedAt    time.Time `xorm:"index"`
	DeletedAt    time.Time
}

type image_any struct {
	Id           int `xorm:"PK SERIAL index"`
	CollectionId int
	ProductId    int
	CreatedAt    time.Time
	UpdatedAt    time.Time `xorm:"index"`
}

func getHostApp() map[string]*myClassic {
	if _hostApp == nil {
		updateHostApp()
	}
	return _hostApp
}

func updateHostApp() {
	hostMappings := make([]Host, 0)
	err := _engine.Find(&hostMappings)
	if err != nil {
		panic(err)
	}
	log.Printf("Host count: %d", len(hostMappings))

	var stores []Store
	err = _engine.Find(&stores)
	if err != nil {
		panic(err)
	}
	log.Printf("Store count: %d", len(stores))

	_hostApp = make(map[string]*myClassic)

	for _, hostMapping := range hostMappings {
		for _, store := range stores {

			if store.Id == hostMapping.StoreId {
				_hostApp[hostMapping.Name] = store.CreateApp()
			}
		}
	}
}

func (hostTable *Host) create() {
	//insert to database
	_, err := _engine.Insert(hostTable)
	if err != nil {
		panic(err)
	}
}

func getHostMappings() *[]Host {
	results := []Host{}
	err := _engine.Find(&results)
	if err != nil {
		panic(err)
	}
	return &results
}

func (source *Collection) toDatabaseForm() error {

	if source == nil {
		myErr := appError{Message: "entity can't be nil"}
		return &myErr
	}

	if len(source.CustomFields) == 0 {
		source.CustomFieldsDB = ""
	} else {
		ba, err := json.Marshal(&source.CustomFields)
		if err != nil {
			return err
		}
		source.CustomFieldsDB = string(ba[:])
	}

	if len(source.Tags) == 0 {
		source.TagsDB = ""
	} else {
		ba, err := json.Marshal(&source.Tags)
		if err != nil {
			return err
		}
		source.TagsDB = string(ba[:])
	}

	if source.Image == nil {
		source.ImageDB = ""
	} else {
		ba, err := json.Marshal(&source.Image)
		if err != nil {
			return err
		}
		source.ImageDB = string(ba[:])
	}

	return nil
}

func (source *Collection) toJsonForm() error {

	if source == nil {
		myErr := appError{Message: "entity can't be nil"}
		return &myErr
	}

	source.Path = "/collections/" + source.ResourceId

	if len(source.CustomFieldsDB) > 0 {
		byteArray := []byte(source.CustomFieldsDB)
		err := json.Unmarshal(byteArray, &source.CustomFields)
		if err != nil {
			return err
		}
	}

	if len(source.TagsDB) > 0 {
		byteArray := []byte(source.TagsDB)
		err := json.Unmarshal(byteArray, &source.Tags)
		if err != nil {
			return err
		}
	}

	if len(source.ImageDB) > 0 {
		ba := []byte(source.ImageDB)
		err := json.Unmarshal(ba, &source.Image)
		if err != nil {
			return err
		}
	} else {
		source.Image = nil
	}

	return nil
}

func (source *Collection) create() error {
	//download image
	var resp *http.Response
	var err error
	if source.Image != nil {
		if len(source.Image.Url) > 0 {
			resp, err = http.Get(source.Image.Url)
			if err != nil {
				logError(err.Error())
				return err
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				downloadErr := appError{Message: "image can't be downloaded"}
				return &downloadErr
			}
		}
		logInfo("got resp of downloaded image")
	}

	//create new transaction
	session := _engine.NewSession()
	defer session.Close()

	err = session.Begin()
	if err != nil {
		logError(err.Error())
		session.Rollback()
		return err
	}

	//get new id
	idGenerator := new(IdGenertator)
	_, err = session.Insert(idGenerator)
	if err != nil {
		logError(err.Error())
		session.Rollback()
		return err
	}
	source.Id = idGenerator.Id

	//save image
	if resp != nil {
		logDebug("saving image")
		fileName := uuid.New() + ".jpg"
		imagePath := filepath.Join(_appDir, fileName)
		out, err := os.Create(imagePath)
		defer out.Close()
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			logError(err.Error())
			session.Rollback()
			return err
		}
		source.Image.FileName = fileName
		source.Image.Url = ""
		source.Image.Path = fmt.Sprintf("/collections/images/%d/%s", idGenerator.Id, fileName)
	}

	//insert collection into database
	source.toDatabaseForm()
	_, err = session.Insert(source)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "UNIQUE constraint failed:") {
			myErr := appError{Ex: err, Message: "The collection was already existing.", Code: 4001}
			session.Rollback()
			return &myErr
		}
		session.Rollback()
		return err
	}

	//insert collection and product relationships
	_, err = session.Delete(&collection_product{CollectionId: source.Id})
	if err != nil {
		session.Rollback()
		panic(err)
	}

	if len(source.ProductIds) > 0 {
		productIds := make([]int, 0)

		for _, element := range source.ProductIds {
			AppendIfMissing(productIds, element)
		}

		col_prods := make([]collection_product, 0)

		for _, element := range productIds {
			col_prod := collection_product{CollectionId: source.Id, ProductId: element}
			col_prods = append(col_prods, col_prod)
		}
		_, err = session.Insert(&col_prods)
		if err != nil {
			session.Rollback()
			panic(err)
		}
	}

	err = session.Commit()
	if err != nil {
		panic(err)
	}

	return nil
}

func (source *Collection) update() error {
	source.UpdatedAt = time.Now().UTC()
	source.toDatabaseForm()

	_, err := _engine.Id(source.Id).Update(source)

	if err != nil {
		return err
	}
	return nil
}

func (source *Collection) delete() error {
	collection, err := GetCollection(source.StoreId, source.Id)
	if err != nil {
		return err
	}
	if collection == nil {
		return nil
	}
	collection.DeletedAt = time.Now().UTC()
	return collection.update()
}

func (source *Product) create() error {
	source.CreatedAt = time.Now().UTC()
	source.UpdatedAt = time.Now().UTC()
	_, err := _engine.Insert(source)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "UNIQUE constraint failed:") {
			myErr := appError{Ex: err, Message: "The product was already existing.", Code: 4001}
			return &myErr
		}
		return err
	}
	return nil
}

func GetCollection(storeId int, collectionId int) (*Collection, error) {
	collection := Collection{Id: collectionId, StoreId: storeId}
	has, err := _engine.Get(&collection)
	if err != nil {
		return nil, err
	}
	if has {
		return &collection, nil
	}
	return nil, nil
}

func GetCollections(storeId int) ([]Collection, error) {
	var collections []Collection
	err := _engine.Where("storeId == ?", storeId).And("DeletedAt < ?", _defaultDatabaseTime).OrderBy("Id").Find(&collections)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (source *Collection) GetTags() []string {
	return []string{}
}

func (source *Collection) SetTags(tags string) error {
	return &appError{Message: "not implemented yet"}
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

	//setup aip
	option := ApiOption{Store: store}
	m.UseApi(option)

	//setup upload server
	m.UseUploadServer()

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
