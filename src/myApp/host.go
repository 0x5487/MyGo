package main

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type HostTable struct {
	Id      int64
	Host    string
	StoreId int64 `xorm:"index"`
}

var _hostApp map[string]*myClassic

func GetHostApp() map[string]*myClassic {
	if _hostApp == nil {
		log.Println("fill hostApp")
		UpdateHostApp()
	}
	return _hostApp
}

func UpdateHostApp() {
	hostTables := make([]HostTable, 0)
	_engine.Find(&hostTables)
	log.Printf("Host count: %d", len(hostTables))

	stores := make(map[int64]*Store)
	_engine.Find(&stores)
	log.Printf("Store count: %d", len(stores))

	_hostApp = make(map[string]*myClassic)

	for _, value := range hostTables {
		_hostApp[value.Host] = stores[value.StoreId].CreateApp()
	}
}
