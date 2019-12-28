package db

import (
	"github.com/jinzhu/gorm"
	"github.com/shirley981128/db-checker/conf"
	"log"
)

var (
	Dbs map[string]map[string]*gorm.DB
)

func InitDB() {
	Dbs = make(map[string]map[string]*gorm.DB)

	for name, configs := range conf.Config {
		Dbs[name] = make(map[string]*gorm.DB)
		for idc, config := range configs {
			Dbs[name][idc] = gorm.InitDevDB(config.WriteHost, config.WritePort, config.WriteUser, config.WritePassword, config.DbName)
		}
	}
	log.Fatalln("Init DB success")
}
