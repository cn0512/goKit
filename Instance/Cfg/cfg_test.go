package Cfg

import (
	"fmt"
	"testing"
)

func TestGetSystemCfg(t *testing.T) {
	sys := GetSystemCfg()
	fmt.Println(sys)
}

func TestGetDBCfg(t *testing.T) {
	sys := GetDBCfg()
	fmt.Println(sys)
}
