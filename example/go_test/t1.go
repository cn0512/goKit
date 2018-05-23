package main

import (
	"fmt"
	"os"
)

type sa struct {
	z string
	a int
	b string
}

func last() {
	fmt.Println("last")
}

func begin() int {
	defer last()
	fmt.Println("begin")
	return 0
}

func main() {
	wd, _ := os.Getwd()
	fmt.Println(wd)

	c := &sa{a: 1}
	d := sa{b: "helo", a: 1, z: "s"}
	fmt.Println(c, d)

	//fmt.Errorf("token.SignedString error %+v", err)

	if hostName := os.Getenv("HOSTNAME"); hostName != "" {
		fmt.Println(hostName)
	} else {
		fmt.Println(hostName)
	}

	e := fmt.Sprintf("%%%s%%", d.b)
	fmt.Print(e)

	fmt.Println(begin())
}
