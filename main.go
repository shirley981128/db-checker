package main

import (
	"log"
	"github.com/shirley981128/db-checker/conf"
	"github.com/shirley981128/db-checker/db"
	"github.com/shirley981128/db-checker/db/dal"
	"github.com/shirley981128/db-checker/db/model"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	if conf.ReadConfig("./conf") != nil {
		panic("read config fail")
	}
	db.InitDB()

	for dbname, dbs := range db.Dbs {
		wg.Add(1)
		go compare(dbs, dbname)
	}
	wg.Wait()
	time.Sleep(time.Second * 3)
}

func compare(dbs map[string]*gorm.DB, dbname string) {
	defer wg.Done()
	log.Println("start compare dbs:%s", dbname)
	if len(dbs) != 2 {
		log.Fatalln("compare fail,instance number of %s error", dbname)
		return
	}
	tables := make(map[string]map[string]model.Table)
	var idcs []string
	for idc, DB := range dbs {
		tables[idc] = make(map[string]model.Table)
		idcs = append(idcs, idc)
		tbs, err := dal.GetTables(DB, dbname)
		for _, tb := range tbs {
			tables[idc][tb.TableName] = tb
		}
		if err != nil {
			log.Fatalln("get %s tables err:%+v", dbname+"."+idc, err)
			return
		}
	}
	for tbname, _ := range tables[idcs[0]] {
		if _, exist := tables[idcs[1]][tbname]; exist {
			compareTable(dbs, dbname, tbname)
			delete(tables[idcs[1]], tbname)
		} else {
			log.Fatalln("%s miss table:%s", dbname+"."+idcs[1], tbname)
		}
	}
	if len(tables[idcs[1]]) > 0 {
		for tbname, _ := range tables[idcs[1]] {
			log.Fatalln("%s miss table:%s", dbname+"."+idcs[0], tbname)
		}
	}

}

func compareTable(dbs map[string]*gorm.DB, dbname string, tablename string) {
	log.Println("start compare tables:%s", tablename)
	columns := make(map[string]map[string]model.Column)
	var idcs []string
	for idc, DB := range dbs {
		columns[idc] = make(map[string]model.Column)
		idcs = append(idcs, idc)
		cols, err := dal.GetColumns(DB, dbname, tablename)
		if err != nil {
			log.Fatalln("get %s columns err:%+v", tablename+"."+idc, err)
			return
		}
		for _, col := range cols {
			columns[idc][col.ColumnName] = col
		}
	}
	for colname, col := range columns[idcs[0]] {
		if col2, exist := columns[idcs[1]][colname]; exist {
			if col != col2 {
				log.Fatalln("there are some differences in %s,col in %s:%+v,col in %s:%+v", tablename+"."+colname, idcs[0], col, idcs[1], col2)
			}
			delete(columns[idcs[1]], colname)
		} else {
			log.Fatalln("%s miss column:%+v", tablename+"."+idcs[1], col)
		}
	}
	if len(columns[idcs[1]]) > 0 {
		for _, col := range columns[idcs[1]] {
			log.Fatalln("%s miss column:%+v", tablename+"."+idcs[0], col)
		}
	}
}
