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

func getHostApp() map[string]*myClassic {
	if _hostApp == nil {
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

	stores := map[int64]*Store{}
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

func (hostTable *HostTable) Create() {
	//insert to database
	_, err := _engine.Insert(hostTable)
	if err != nil {
		panic(err)
	}
}

func getHostTables() *[]HostTable {
	results := []HostTable{}
	err := _engine.Find(&results)
	if err != nil {
		panic(err)
	}
	return &results
}
