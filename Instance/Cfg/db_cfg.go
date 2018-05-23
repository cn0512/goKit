package Cfg

import (
	cons "../../ykconstant"
)

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

type DBs struct {
	Db []DB `xml:"db"` //big Capital
}

type DB struct {
	Sq    Sqlite `xml:"sqlite"`
	Redis Redis  `xml:"redis"`
	Mysql Mysql  `xml:"mysql"`
}

func (yk *DB) GetSqliteFile() string {
	return yk.Sq.File
}

type Sqlite struct {
	File       string `xml:"file"`
	Db         string `xml:"db"`
	Table_user string `xml:"table_user"`
}

type Redis struct {
	Uri string `xml:"uri"`
}

type Mysql struct {
	Uri string `xml:"uri"`
}

var so_db sync.Once
var cfg_db *DB

func GetDBCfg() *DB {

	so_db.Do(func() {
		cfg_db = &DB{}
		if true != cfg_db.load() {
			fmt.Println("db.cfg load err")
		}
	})

	return cfg_db
}

func (yk *DB) load() bool {

	content, err := ioutil.ReadFile(cons.Cfg_db_file)
	//fmt.Println(content)
	if err != nil {
		log.Fatal(err)
		return false
	}

	err = xml.Unmarshal(content, yk)
	if err != nil {
		log.Fatal(err)
		return false
	}

	//fmt.Println(cfg)
	//fmt.Println(cfg.Sq)
	return true
}
