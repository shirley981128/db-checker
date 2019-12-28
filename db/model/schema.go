package model

type Table struct {
	TableName string `gorm:"column:table_name" json:"table_name"`
}

type Column struct {
	ColumnName    string `gorm:"column:COLUMN_NAME" json:"column_name"`
	ColumnDefault string `gorm:"column:COLUMN_DEFAULT" json:"column_default"`
	IsNullable    string `gorm:"column:IS_NULLABLE" json:"is_nullable"`
	ColumnType    string `gorm:"column:COLUMN_TYPE" json:"column_type"`
	CharSetName   string `gorm:"column:CHARACTER_SET_NAME" json:"character_set_name"`
	CollationName string `gorm:"column:COLLATION_NAME" json:"collation_name"`
	ColumnKey     string `gorm:"column:COLUMN_KEY" json:"column_key"`
}
