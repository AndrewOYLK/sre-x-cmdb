package config

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

var Config *AppConfig

func init() {
	if Config != nil {
		return
	}

	var once sync.Once
	once.Do(func() {
		Config = &AppConfig{}

		bytes, err := ioutil.ReadFile("project.yaml")
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(bytes, Config)
		if err != nil {
			panic(err)
		}
	})
}

type AppConfig struct {
	Server  Server  `yaml:"server"`
	Metrics Metrics `yaml:"metrics"`
	Mysql   Mysql   `yaml:"mysql"`
	Redis   Redis   `yaml:"redis"`
	MongoDB MongoDB `yaml:"mongodb"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Metrics struct {
	Path string `yaml:"path"`
	Port string `yaml:"port"`
}

type Mysql struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	DBName      string `yaml:"dbname"`
	MaxConn     int    `yaml:"maxconn"`
	MaxIdleConn int    `yaml:"maxidleconn"`
	Charset     string `yaml:"charset"`
	ShowSQL     bool   `yaml:"showsql"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type MongoDB struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
