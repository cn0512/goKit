package main

import (
	_ "fmt"
	"os"
	"os/signal"

	cons "../../ykconstant"
	"../Cfg"
)

func main() {
	var idb IDB
	sysCfg := Cfg.GetSystemCfg()
	switch sysCfg.Db.Use {
	case cons.DB_Redis:
		idb = &RedisDB{}
	case cons.DB_Sqlite:
		idb = &SqliteDB{}
	default:
		return
	}

	idb.About()
	idb.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
