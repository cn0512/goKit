package yksqlite

/*
	基于go-sqlite3项目的API封装
*/

import (
	"database/sql"
	"fmt"
	_ "os"

	"../Instance/Cfg"
	yklog "../yklog"
	_ "github.com/mattn/go-sqlite3"
)

type YKSqliteDB struct {
	Db *sql.DB
}

func (yk *YKSqliteDB) Open() error {
	dbcfg := Cfg.GetDBCfg()
	dbfile := dbcfg.GetSqliteFile()
	var err error
	yk.Db, err = sql.Open("sqlite3", dbfile)
	if err != nil {
		yklog.Logout("error=", err)
		fmt.Println(err)
		return err
	}
	//defer yk.Db.Close()
	return nil
}

func (yk *YKSqliteDB) Update(sql string) error {
	_, err := yk.Db.Exec(sql)
	if err != nil {
		yklog.Logout("%q: %s\n", err, sql)
		return err
	}
	return nil
}
func (yk *YKSqliteDB) Query(sql string, dest ...interface{}) error {
	rows, err := yk.Db.Query(sql)
	if err != nil {
		yklog.Logout("error-1-=", err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			yklog.Logout("error-2=", err)
			return err
		}
	}
	err = rows.Err()
	if err != nil {
		yklog.Logout("error-3-=", err)
		return err
	}
	return nil
}
