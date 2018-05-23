package Cfg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	_ "log"
	"sync"

	cons "../../ykconstant"
	yklog "../../yklog"
)

type Syss struct {
	SysNode []Sys `xml:"system"` //big Capital
}

type Sys struct {
	Protocol Sel `xml:"protocol"`
	Db       Sel `xml:"db"`
	Mq       Sel `xml:"mq"`
}

func (yk *Sys) GetProtocolType() string {
	return yk.Protocol.Use
}
func (yk *Sys) GetDBType() string {
	return yk.Db.Use
}

func (yk *Sys) GetMqType() string {
	return yk.Mq.Use
}

type Sel struct {
	Use string `xml:"use"`
}

var so_sys sync.Once
var cfg_sys *Sys

func GetSystemCfg() *Sys {

	so_sys.Do(func() {
		cfg_sys = &Sys{}
		if true != cfg_sys.load() {
			fmt.Println("system.cfg load err")
		}
	})

	return cfg_sys
}
func (yk *Sys) load() bool {

	content, err := ioutil.ReadFile(cons.Cfg_system_file)
	//fmt.Println(content)
	if err != nil {
		yklog.Logout("errors=", err)
		return false
	}

	err = xml.Unmarshal(content, yk)
	if err != nil {
		yklog.Logout("errors=", err)
		return false
	}

	//fmt.Println(yk)
	return true
}
