package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Db Db
}

type Db struct {
	Addr    string
	Port    int
	User    string
	Pass    string
	LibName string
}

var (
	cfgPath string
	gConfig Config
)

func init() {
	flag.Parse()
	flag.StringVar(&cfgPath, "config", ".", "config path")

	buff, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(buff, &gConfig)
	if err != nil {
		panic(err)
	}
}
