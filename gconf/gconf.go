//Author: nghiatc
//Since: Jul 15, 2021

package gconf

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper
var env string

func Init(wdir string) {
	//// init Configuration
	environment := flag.String("e", "development", "run project with mode [-e development | test | production]")
	flag.Usage = func() {
		fmt.Println("Usage: ./[appname] -e development | test | production")
		os.Exit(1)
	}
	flag.Parse()
	log.Printf("============== GConf environment: %s", *environment)
	InitEnv(wdir, *environment)
}

// Init is an exported method that takes the environment starts the viper (external lib) and
// returns the configuration struct.
func InitEnv(wdir string, environment string) {
	log.Printf("============== GConfig Init Environment: %s ==============", environment)
	var err error
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(environment)
	v.AddConfigPath(filepath.Join(wdir, "conf"))
	err = v.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
	config = v
	env = environment
}

func JoinPath(basedir string, path string) string {
	return filepath.Join(basedir, path)
}

func GetConfig() *viper.Viper {
	return config
}

func GetEnv() string {
	return env
}
