package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	goorm "goorm/day2"
	_ "goorm/day2/dialect"
	"goorm/day2/log"
)


func main() {
	type Java struct {
		Name string
		Age int
	}
	engine, err := goorm.NewEngine("sqlite3", "/Users/zhangmengnan/env/sqlite-tools-osx-x86-3380100/file/gee.db")
	if err != nil {
		log.Error(err)
	}
	session := engine.NewSession()

	session = session.Model(&Java{})
	err = session.CreateTable()
	if err != nil{
		fmt.Println("create table error")
	}

}

