package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
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

	//ida, _ := mid.GetNext("aaa")
	//fmt.Println("ida =", ida)
	//// rs, _ := mid.ResetID("aaa", 0)
	//// fmt.Println("rs =", rs)
	//gida, _ := mid.GetID("aaa")
	//fmt.Println("gida =", gida)
	//
	//idb, _ := mid.GetNext("bbb")
	//fmt.Println("idb =", idb)
	//gidb, _ := mid.GetID("bbb")
	//fmt.Println("gidb =", gidb)

	// get id not exits
	gidNotExits, err := mid.GetID("gidNotExits")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("gidNotExits =", gidNotExits)

	//// Test counter increment
	//c1, _ := mid.IncrementCounter("counter1", 5)
	//fmt.Println("c1 =", c1)
	//gc1, _ := mid.GetID("counter1")
	//fmt.Println("gc1 =", gc1)
	//// Test counter decrement
	//c2, _ := mid.IncrementCounter("counter1", -8)
	//fmt.Println("c2 =", c2)
	//gc2, _ := mid.GetID("counter1")
	//fmt.Println("gc2 =", gc2)
	//
	//c3, _ := mid.IncrementCounter("counter3", 5)
	//fmt.Println("c3 =", c3)

	// Hang thread Main.
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C) SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c
	log.Println("################# End Main #################")
}
