package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/congnghia0609/ntc-gmongo/gconf"
	"github.com/congnghia0609/ntc-gmongo/gmongo"
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

}
