package main

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
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
	Id           int `xorm:"PK SERIAL index"`
	StoreId      int `xorm:"INT not null index" form:"-" json:"-"`
	Url          string
	Position     int
	FileName     string
	Attachment   string
	CustomFields LinkModel `xorm:"-"`
}

type CustomField struct {
	Id        int    `xorm:"PK SERIAL index"`
	StoreId   int    `xorm:"INT not null unique(custom_field)" form:"-" json:"-"`
	TypeId    int    `xorm:"INT not null unique(custom_field)"`
	ParentId  int    `xorm:"INT not null unique(custom_field)"`
	Key       string `xorm:"not null unique(custom_field)"`
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time `xorm:"index"`
}

type Collection struct {
	Id           int    `xorm:"PK SERIAL index"`
	StoreId      int    `xorm:"INT not null unique(resourceId) unique(name)" form:"-" json:"-"`
	ResourceId   string `xorm:"not null unique(resourceId)"`
	DisplayName  string `xorm:"not null unique(name)"`
	IsVisible    bool
	Description  string
	Image        Image `xorm:"-"`
	Tags         string
	CustomFields LinkModel `xorm:"-"`
	CreatedAt    time.Time
	UpdatedAt    time.Time `xorm:"index"`
}

type Product struct {
	Id                        int `xorm:"PK SERIAL index"`
	StoreId                   int `xorm:"INT not null index" form:"-" json:"-"`
	Sku                       string
	ResourceId                string
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
	InventoryQuantity         int
	ManageInventoryMethod     ManageInventoryMethod
	Weight                    int32
	Variations                LinkModel `xorm:"-"`
	Images                    LinkModel `xorm:"-"`
	CustomFields              LinkModel `xorm:"-"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time `xorm:"index"`
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
	InventoryQuantity         int
	ManageInventoryMethod     ManageInventoryMethod
	Weight                    int32
	CreatedAt                 time.Time
	UpdatedAt                 time.Time `xorm:"index"`
}

//database bridge table

type collection_product struct {
	Id           int `xorm:"PK SERIAL index"`
	CollectionId int
	ProductId    int
	CreatedAt    time.Time
	UpdatedAt    time.Time `xorm:"index"`
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

func (source *Collection) create() error {
	source.CreatedAt = time.Now().UTC()
	source.UpdatedAt = time.Now().UTC()
	_, err := _engine.Insert(source)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "UNIQUE constraint failed:") {
			myErr := appError{Ex: err, Message: "The collection was already existing.", Code: 4001}
			return &myErr
		}
		return err
	} else {
		return nil
	}
}

func (source *CustomField) create() error {
	source.CreatedAt = time.Now().UTC()
	source.UpdatedAt = time.Now().UTC()
	_, err := _engine.Insert(source)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "UNIQUE constraint failed:") {
			myErr := appError{Ex: err, Message: "The custom Field was already existing.", Code: 4001}
			return &myErr
		}
		return err
	} else {
		return nil
	}
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
	} else {
		return nil
	}
}

func GetCollections(storeId int) []Collection {
	return []Collection{}
}

func (source *Collection) GetTags() []string {
	return []string{}
}

func (source *Collection) SetTags(tags string) error {
	return &appError{Message: "not implemented yet"}
}
