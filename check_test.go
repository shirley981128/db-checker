package main

import (
	"github.com/shirley981128/db-checker/conf"
	"github.com/shirley981128/db-checker/db"
	"github.com/shirley981128/db-checker/db/dal"
	"encoding/json"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	if conf.ReadConfig("./conf") != nil {
		panic("read config fail")
	}
	fmt.Println("success,config:", conf.Config)
	db.InitDB()
	for dbname, dbs := range db.Dbs {
		for _, db := range dbs {
			result, err := dal.GetTables(db, dbname)
			if err != nil {
				fmt.Println(err)
			}
			for _, v := range result {
				columns, err := dal.GetColumns(db, dbname, v.TableName)
				if err != nil {
					fmt.Println("err:", err)
				}
				cols, err := json.Marshal(columns)
				fmt.Println(string(cols))
			}
		}
	}
}
