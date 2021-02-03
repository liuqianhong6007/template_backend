package config

import (
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Db struct {
	Driver  string `yaml:"driver"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	User    string `yaml:"user"`
	Pass    string `yaml:"pass"`
	LibName string `yaml:"lib_name"`
}

var (
	gConfig Config
)

func init() {
	var cfgPath string
	flag.StringVar(&cfgPath, "config", "config.yaml", "config path")
	flag.Parse()

	buff, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(buff, &gConfig)
	if err != nil {
		panic(err)
	}
}

func ServerCfg() Server {
	return gConfig.Server
}

func DbCfg() Db {
	return gConfig.Db
}
