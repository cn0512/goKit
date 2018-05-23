package main

import (
	"flag"
	"fmt"
	"log"
	_ "os"
	"time"

	c "./ykconstant"
	//net "./yknet"
	"github.com/golang/protobuf/proto"

	ykmsg "./ykmsg"
	"./ykprotoco"

	"./Instance/Cfg"
	consistent "./ykconsistent"
)

var (
	version = flag.String("v", "0.0.1", "version = ?")
	tcp_ip  = flag.String("ip", "127.0.0.1:6018", "port = ?")
)

type gobSt struct {
	Age  int
	Name string
}

func main() {
	flag.Parse()
	fmt.Println(c.Version)
	fmt.Println(c.TCP_IP)
	fmt.Println(*version)
	fmt.Println(*tcp_ip)

	header := ykmsg.MsgHeader{1029, 20, 1}
	headerbuf := header.Get()
	fmt.Println(headerbuf)

	var login YKGameMsg.LoginMsgReq
	login.Uid = 1000
	login.Pwd = "passord"
	login.CheckCode = "code"
	buf := login.String()
	fmt.Println("login1=", login, buf)
	data, _ := proto.Marshal(&login)
	fmt.Println("login1,buf=", data)
	var login2 YKGameMsg.LoginMsgReq
	proto.Unmarshal(data, &login2)
	fmt.Println("login2=", login2)

	var begin YKGameMsg.GameBegin
	begin.Seat = 1
	begin.Players = make([]*YKGameMsg.StPlayer, 2)
	cards0 := []int32{1, 2, 3, 4, 5}
	begin.Players[0] = &YKGameMsg.StPlayer{1, cards0}

	fmt.Println(begin)

	/*
		Gob
	*/
	//encode
	st := gobSt{12, "123"}
	var ykgob ykmsg.GobMsg
	bufMsg, err := ykgob.Encode(&st, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("encode=", bufMsg)
	//decode
	var st2 gobSt
	err = ykgob.Decode(bufMsg, &st2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("decode=", st2)

	/*
		consistent 分库分表测试;为了更少的数据迁移；
	*/

	c := consistent.New()
	c.Add("table1")
	c.Add("table2")
	//c.Add("table3")
	users := []string{"user_1", "user_2", "user_3", "user_4", "user_5", "user_5", "user_6", "user_7", "user_8", "user_9", "user_10"}
	fmt.Println("----init table----")
	for _, u := range users {
		server, err := c.Get(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %s\n", u, server)
	}

	c.Add("table3")
	//c.Add("table4")
	//c.Add("table5")
	//c.Add("table6")

	fmt.Println("\n----add some new table----")
	for _, u := range users {
		server, err := c.Get(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s => %s\n", u, server)
	}

	/*load xml*/
	db := Cfg.GetDBCfg()
	fmt.Println(db)
	sys := Cfg.GetSystemCfg()
	fmt.Println(sys)

	/*timer*/
	now := time.Now()
	fmt.Println(now)
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	zone, _ := now.UTC().Zone()
	fmt.Printf("UTC time is %d-%d-%d %02d:%02d:%02d %s\n",
		year, mon, day, hour, min, sec, zone)
	year, mon, day = now.Date()
	hour, min, sec = now.Clock()
	zone, _ = now.Zone()
	fmt.Printf("local time is %d-%d-%d %02d:%02d:%02d %s\n",
		year, mon, day, hour, min, sec, zone)

	time.Sleep(3 * time.Second)
	end_time := time.Now()
	var dur_time time.Duration = end_time.Sub(now)
	fmt.Println(dur_time)
}
