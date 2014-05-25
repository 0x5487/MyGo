package main

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type HostMapping struct {
	Id      int `xorm:"SERIAL index"`
	StoreId int `xorm:"INT index"`
	Host    string
}

func getHostApp() map[string]*myClassic {
	if _hostApp == nil {
		updateHostApp()
	}
	return _hostApp
}

func updateHostApp() {
	hostMappings := make([]HostMapping, 0)
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
				_hostApp[hostMapping.Host] = store.CreateApp()
			}
		}
	}
}

func (hostTable *HostMapping) create() {
	//insert to database
	_, err := _engine.Insert(hostTable)
	if err != nil {
		panic(err)
	}
}

func getHostMappings() *[]HostMapping {
	results := []HostMapping{}
	err := _engine.Find(&results)
	if err != nil {
		panic(err)
	}
	return &results
}
