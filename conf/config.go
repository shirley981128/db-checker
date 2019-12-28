package conf

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	Config map[string]map[string]DBConfig
)

type DBConfig struct {
	DbName         string `yaml:"dbName"`
	WriteHost      string `yaml:"writeHost"`
	WritePort      string `yaml:"writePort"`
	WriteUser      string `yaml:"writeUser"`
	WritePassword  string `yaml:"writePassword"`
	WriteTimeout   string `yaml:"writeTimeout"`
	ReadHost       string `yaml:"readHost"`
	ReadPort       string `yaml:"readPort"`
	ReadUser       string `yaml:"readUser"`
	ReadPassword   string `yaml:"readPassword"`
	ReadTimeout    string `yaml:"readTimeout"`
	ConnectTimeout string `yaml:"connectTimeout"`
	LogMode        bool   `yaml:"logMode"`
}

func ReadConfig(path string) error {
	Config = make(map[string]map[string]DBConfig)
	fileName := path + "/" +"test.yml"
	if f, e := ioutil.ReadFile(fileName); e != nil {
		log.Fatalln("failed to read configuration file : %s", fileName)
		return errors.New("failed to read the configuration")
	} else {
		if e := yaml.Unmarshal(f, Config); e != nil {
			log.Fatalln("failed to unmarshal the configuration")
			return errors.New("failed to unmarshal the configuration")
		}
	}
	log.Fatalln("succ to read configuration file : %v", Config)
	return nil
}
