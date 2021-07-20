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

	ida, _ := mid.GetNext("aaa")
	fmt.Println("ida =", ida)
	// rs, _ := mid.ResetID("aaa", 0)
	// fmt.Println("rs =", rs)

	idb, _ := mid.GetNext("bbb")
	fmt.Println("idb =", idb)

	// Hang thread Main.
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C) SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c
	log.Println("################# End Main #################")
}
