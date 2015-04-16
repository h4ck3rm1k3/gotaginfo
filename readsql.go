package main

import _ "github.com/mattn/go-sqlite3"
//import "database/sql"
import "fmt"
import "github.com/h4ck3rm1k3/gorm"


type SqliteMaster struct {
	Type string
	Name string
	TblName string `gorm:"column:tbl_name"`
	RootPage int
	Sql string
};

func (c SqliteMaster) TableName() string {
	return "sqlite_master"
}

func DescribeDb(dbfile string) error{
	db, err := gorm.Open("sqlite3", dbfile)
	if err != nil {
		fmt.Printf("ERROR:%v\n", err)
		return err
	}

	db.DB()

	//field name
	var rows []SqliteMaster
	db.Find(&rows)
	if err != nil {
		fmt.Printf("ERROR:%v\n", err)
		return err
	}
	for _, row :=range rows{
		fmt.Printf("%v\n", row)
	}
	return nil
}

