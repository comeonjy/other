package configure

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func init() {
	LoadConfigure()
}

type Config struct {
	Database database `yaml:"database"`
}

type database struct {
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
}

func Get() *Config {
	return cfg
}

var cfg *Config

func LoadConfigure() {
	cfg = &Config{}
	bytes, e := ioutil.ReadFile("D:/workspace/go/other/src/config_test.yaml")
	if e != nil {
		log.Fatal(e)
	}

	if e := yaml.Unmarshal(bytes, cfg); e != nil {
		log.Fatal(e)
	}

	log.Println(*cfg)
}
