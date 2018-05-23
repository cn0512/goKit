package main

import (
	"fmt"

	"../../yksqlite"
)

type SqliteDB struct {
}

func (yk *SqliteDB) About() {
	fmt.Println("will start Sqlite DB")

}

func (yk *SqliteDB) Run() {
	/*sqlite*/
	sqlitedb := yksqlite.YKSqliteDB{nil}
	err := sqlitedb.Open()
	if err != nil {
		fmt.Println("-1-")
		return
	}
	sql := "select gameid,name from game_cfg"
	var gameid int
	var name string
	err = sqlitedb.Query(sql, &gameid, &name)
	if err != nil {
		fmt.Println("-2-")
		return
	}
	fmt.Println(gameid, name)
}
