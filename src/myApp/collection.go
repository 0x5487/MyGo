package main

type Collection struct {
	Id          int
	Name        string
	Description string
}

func GetCollections(storeId int64) []Collection {
	return []Collection{}
}
