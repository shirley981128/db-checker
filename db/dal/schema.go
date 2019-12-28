package dal

import (
	"github.com/jinzhu/gorm"
	"github.com/shirley981128/db-checker/db/model"
	"log"
)

func GetTables(db *gorm.DB, dbname string) ([]model.Table, error) {
	var result []model.Table
	if err := db.Table("tables").Select("table_name").Where("table_schema=?", dbname).Scan(&result).Error; err != nil {
		log.Fatalln("db error:%+v", err)
		return nil, err
	}
	return result, nil
}

func GetColumns(db *gorm.DB, dbname string, tableName string) ([]model.Column, error) {
	var result []model.Column
	if err := db.Table("columns").Where("table_schema=? and table_name=?", dbname, tableName).Scan(&result).Error; err != nil {
		log.Fatalln("db error:%+v", err)
		return nil, err
	}
	return result, nil
}
