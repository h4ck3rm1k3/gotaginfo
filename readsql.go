package main

import _ "github.com/mattn/go-sqlite3"
//import "database/sql"
import "fmt"
import "github.com/h4ck3rm1k3/gorm"
import "database/sql"

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

	db2, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		fmt.Printf("Failed to open database:", err)
	}

	//field name
	var rows []SqliteMaster
	db.Find(&rows)
	if err != nil {
		fmt.Printf("ERROR:%v\n", err)
		return err
	}
	for _, row :=range rows{
		fmt.Printf("NAME:%s\n", row.Name)
		fmt.Printf("%v\n", row)

		sql := "SELECT * FROM " + row.Name + " limit 1"
		fmt.Printf("SQL:%s\n", sql)
		rows2, err2 := db2.Query(sql)
		fmt.Printf("ROWS2:%v\n", rows2)
		//fmt.Printf("ROWS2[0]:%v\n", rows2[0])
		fmt.Printf("err2:%v\n", err2)

		if (rows2 != nil) {
			cols,err3 := rows2.Columns()
			fmt.Printf("COLS:%v\n", cols)
			for i, name := range cols {
				fmt.Printf("COL:%n %v = %v\n", i,name)
			}
			fmt.Printf("err3:%v\n", err3)
		}
		

	}
	return nil
}

