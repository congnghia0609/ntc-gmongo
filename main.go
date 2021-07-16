package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/congnghia0609/ntc-gmongo/gconf"
	"github.com/congnghia0609/ntc-gmongo/gmongo"
	"github.com/congnghia0609/ntc-gmongo/mid"
)

func InitGConf() {
	_, b, _, _ := runtime.Caller(0)
	wdir := filepath.Dir(b)
	fmt.Println("wdir:", wdir)
	gconf.Init(wdir)
}

func main() {
	InitGConf()

	fmt.Println("Hello world...")

	gmongo.InitMongo()
	defer gmongo.MClose()

	// id, _ := mid.GetNext("bbb")
	// fmt.Println("id =", id)
	// rs, _ := mid.ResetID("bbb", 0)
	// fmt.Println("rs =", rs)

	id, _ := mid.GetNext("bbb")
	fmt.Println("id =", id)
}
