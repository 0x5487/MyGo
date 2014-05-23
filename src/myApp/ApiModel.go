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
	Id           int `xorm:"SERIAL index"`
	StoreId      int
	Url          string
	Position     int
	FileName     string
	Attachment   string
	CustomFields LinkModel
}

type CustomField struct {
	Id      int    `xorm:"SERIAL index"`
	StoreId int    `xorm:"INT not null unique(custom_field) index"`
	Key     string `xorm:"not null unique(custom_field)"`
	Value   string
}

type Collection struct {
	Id           int    `xorm:"SERIAL index"`
	StoreId      int    `xorm:"INT not null unique(collection) index" form:"-" json:"-"`
	ResourceId   string `xorm:"not null unique(collection) index"`
	DisplayName  string `xorm:"not null unique(collection) index"`
	IsVisible    bool
	Description  string
	Image        Image
	Tags         string
	CustomFields LinkModel
	CreatedAt    time.Time
	UpdatedAt    time.Time `xorm:"index"`
}

type Product struct {
	Id                        int `xorm:"SERIAL index"`
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
	ListPrice                 Money
	Price                     Money
	Description               string
	Vendor                    string
	InventoryQuantity         int
	ManageInventoryMethod     ManageInventoryMethod
	Weight                    int32
	Variations                LinkModel
	Images                    LinkModel
	CustomFields              LinkModel
	CreatedAt                 time.Time
	UpdatedAt                 time.Time `xorm:"index"`
}

type Variation struct {
	Id                        int `xorm:"SERIAL index"`
	StoreId                   int `xorm:"INT index"`
	Sku                       string
	DisplayName               string
	IsPurchasable             bool
	IsVisible                 bool
	IsBackOrderEnabled        bool
	IsPreOrderEnabled         bool
	IsShippingAddressRequired bool
	Tags                      string
	ListPrice                 Money
	Price                     Money
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
	Id           int `xorm:"SERIAL index"`
	ImageId      int
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
