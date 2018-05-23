package main

import (
	"fmt"
)

type RedisDB struct {
	IDB
}

func (yk *RedisDB) About() {
	fmt.Println("will start Redis DB")

}

func (yk *RedisDB) Run() {

}
