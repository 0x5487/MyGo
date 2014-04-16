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

func getHostApp() map[string]*myClassic {
	if _hostApp == nil {
		log.Println("filling hostApp")
		updateHostApp()
	}
	return _hostApp
}

func updateHostApp() {
	hostTables := make([]HostTable, 0)
	err := _engine.Find(&hostTables)
	log.Printf("Host count: %d", len(hostTables))
	if err != nil {
		panic(err)
	}

	stores := make(map[int64]*Store)
	err = _engine.Find(&stores)
	log.Printf("Store count: %d", len(stores))
	if err != nil {
		panic(err)
	}

	_hostApp = make(map[string]*myClassic)

	for _, value := range hostTables {
		_hostApp[value.Host] = stores[value.StoreId].CreateApp()
	}
}
