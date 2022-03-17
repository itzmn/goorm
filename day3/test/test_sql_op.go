package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	goorm "goorm/day3"
	"goorm/day3/log"
)

func main() {

	type User struct {
		Name string
		Age int
	}
	engine, err := goorm.NewEngine("sqlite3", "/Users/zhangmengnan/env/sqlite-tools-osx-x86-3380100/file/gee.db")
	if err != nil {
		log.Error(err)
	}

	session := engine.NewSession()

	var u []User
	err = session.Find(&u)
	fmt.Println("eer:", err)
	//session.Insert(&User{Name: "test", Age: 12})
	//fmt.Sprintf(u.Name)

	fmt.Sprintf("fff")

	for _, user := range u {
		fmt.Println(user.Name)
	}




}
