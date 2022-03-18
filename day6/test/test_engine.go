package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"goorm/day1"
	"goorm/day1/log"
)

func main() {

	engine, err := goorm.NewEngine("sqlite3", "/Users/zhangmengnan/env/sqlite-tools-osx-x86-3380100/file/gee.db")
	if err != nil {
		log.Error(err)
	}

	session := engine.NewSession()

	//result, _ := session.Raw("insert into User(`name`) values (?), (?)", "aa", "bb").Exec()
	raws, err := session.Raw("select name, age from User").QueryRaws()
	for raws.Next() {
		var (
			name string
			age  int
		)
		if err := raws.Scan(&name, &age); err != nil {
			log.Error(err)
		}
		fmt.Printf("scan data name:%s, age:%v\n", name, age)
	}
	//count, _ := result.RowsAffected()
	//fmt.Println(count)
}
