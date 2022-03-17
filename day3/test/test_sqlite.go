package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// 测试Sqlite3

func main() {

	db, _ := sql.Open("sqlite3", "/Users/zhangmengnan/env/sqlite-tools-osx-x86-3380100/file/gee.db")
	defer db.Close()
	//exec, err := db.Exec("INSERT into User (`name`) values (?), (?)", "zhangsan", "lisi")
	//if err == nil {
	//	fmt.Println(exec.RowsAffected())
	//}

	//row := db.QueryRow("select name from User")
	//var name string

	type User struct {
		Name string
		Age  int
	}
	rows, _ := db.Query("select Name,Age from User")

	var (
		Name string
		Age int
	)
	for rows.Next() {

		if err := rows.Scan(&Name, &Age); err == nil {
			fmt.Println(Name, Age)
		}else {
			fmt.Println(err)
		}
	}

}
