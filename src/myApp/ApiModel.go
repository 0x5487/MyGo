package main

type ManageInventoryMethod int
type Money int64

const (
	NoTrack = iota
	Track
	External
)

type LinkModel struct {
	Url string
}

type Image struct {
	Id           int32
	StoreId      int
	Url          string
	Position     int
	FileName     string
	Attachment   string
	CustomFields LinkModel
}

type CustomField struct {
	Id      int32
	StoreId int
	Key     string
	Value   string
}

type Collection struct {
	Id           int32
	StoreId      int
	ResourceId   string
	DisplayName  string
	IsVisible    bool
	Description  string
	Image        Image
	Tags         string
	CustomFields LinkModel
}

type Product struct {
	Id                        int32
	StoreId                   int
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
}

type Variation struct {
	Id                        int32
	StoreId                   int
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
}

//database bridge table

type collection_product struct {
	Id           int32
	CollectionId int32
	ProductId    int32
}

type image_any struct {
	Id           int32
	ImageId      int32
	CollectionId int32
	ProductId    int32
}
