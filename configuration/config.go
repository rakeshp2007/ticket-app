package configuration

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Configuration struct {
	MongoDB struct {
		Host     string
		Username string
		Password string
		Database string
		Port     string
	}
	JWT struct {
		JwtKey    string
		TtlMinute int
	}
}

var Config = Configuration{}

func SetConfigParams() {
	_, b, _, _ := runtime.Caller(0)
	currentPath := filepath.Dir(b)
	c := flag.String("c", currentPath+"/../config.json", "Specify configuration file.")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		log.Fatal("Can't open config file", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("Can't decode config file", err)
	}

}
