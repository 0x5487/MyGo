package main

import (
	"encoding/json"
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
	Id         int `xorm:"PK SERIAL index"`
	StoreId    int `xorm:"INT not null index" form:"-" json:"-"`
	Url        string
	Position   int
	FileName   string
	Attachment string
}

type CustomField struct {
	Name  string
	Value string
}

type Collection struct {
	Id             int    `xorm:"PK SERIAL index"`
	StoreId        int    `xorm:"INT not null unique(resourceId) unique(name)" form:"-" json:"-"`
	ResourceId     string `xorm:"not null unique(resourceId)"`
	DisplayName    string `xorm:"not null unique(name)"`
	IsVisible      bool
	Content        string
	Image          Image `xorm:"-"`
	Tags           string
	ProductIds     []int         `xorm:"-"`
	CustomFieldsDB string        `json:"-"`
	CustomFields   []CustomField `xorm:"-"`
	CreatedAt      time.Time     `json:"-"`
	UpdatedAt      time.Time     `xorm:"index" json:"-"`
	DeletedAt      time.Time     `json:"-"`
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

func (source *Collection) ToDatabaseForm() error {

	if source == nil {
		myErr := appError{Message: "entity can't be nil"}
		return &myErr
	}

	if source.CustomFields == nil || len(source.CustomFields) == 0 {
		source.CustomFieldsDB = ""
	} else {
		ba, err := json.Marshal(source.CustomFields)
		if err != nil {
			return err
		}
		source.CustomFieldsDB = string(ba[:])
	}

	return nil
}

func (source *Collection) ToJsonForm() error {

	if source == nil {
		myErr := appError{Message: "entity can't be nil"}
		return &myErr
	}

	if len(source.CustomFieldsDB) > 0 {
		ba := []byte(source.CustomFieldsDB)
		var customfields []CustomField
		err := json.Unmarshal(ba, &customfields)
		if err != nil {
			return err
		}
		source.CustomFields = customfields
	}

	return nil
}

func (source *Collection) create() error {
	source.CreatedAt = time.Now().UTC()
	source.UpdatedAt = time.Now().UTC()
	source.ToDatabaseForm()

	_, err := _engine.Insert(source)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "UNIQUE constraint failed:") {
			myErr := appError{Ex: err, Message: "The collection was already existing.", Code: 4001}
			return &myErr
		}
		return err
	}

	return nil
}

func (source *Collection) update() error {
	source.UpdatedAt = time.Now().UTC()
	source.ToDatabaseForm()

	_, err := _engine.Id(source.Id).Update(source)

	if err != nil {
		return err
	}

	return nil
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
	err := _engine.Where("storeId == ?", storeId).Find(&collections)

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
