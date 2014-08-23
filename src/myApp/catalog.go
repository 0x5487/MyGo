package main

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
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

type IdGenertator struct {
	Id   int `xorm:"PK SERIAL"`
	Name string
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
