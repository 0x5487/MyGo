package main

import (
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

type Image struct {
	Id           int `xorm:"PK SERIAL index"`
	StoreId      int `xorm:"INT index"`
	Url          string
	Position     int
	FileName     string
	Attachment   string
	CustomFields LinkModel `xorm:"-"`
}

type CustomField struct {
	Id      int    `xorm:"PK SERIAL index"`
	StoreId int    `xorm:"INT not null unique(custom_field) index"`
	Key     string `xorm:"not null unique(custom_field)"`
	Value   string
}

type Collection struct {
	Id           int    `xorm:"PK SERIAL index"`
	StoreId      int    `xorm:"INT not null unique(resourceId) unique(name) index" form:"-" json:"-"`
	ResourceId   string `xorm:"not null unique(resourceId) index"`
	DisplayName  string `xorm:"not null unique(name) index"`
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
	StoreId                   int `xorm:"INT index"`
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
	StoreId                   int `xorm:"INT index"`
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
	Id           int `xorm:"SERIAL index"`
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

func (source *Collection) create() error {
	source.CreatedAt = time.Now().UTC()
	source.UpdatedAt = time.Now().UTC()
	_, err := _engine.Insert(source)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "UNIQUE constraint failed:") {
			myErr := appError{Ex: err, Message: "collection name was already existing.", Code: 4001}
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
