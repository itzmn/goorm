package main

import (
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	goorm "goorm/day6"
	"goorm/day6/log"
	sesssion "goorm/day6/session"
)

func test_sql_op() {
	type User struct {
		Name string
		Age  int
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

func test_update_delete_and_chain() {
	type User struct {
		Name string
		Age  int
	}
	engine, err := goorm.NewEngine("sqlite3", "/Users/zhangmengnan/env/sqlite-tools-osx-x86-3380100/file/gee.db")
	if err != nil {
		log.Error(err)
	}

	session := engine.NewSession()

	var u []User
	err = session.Where("Name=?", "Tom").OrderBy("Age desc").Limit(1).Find(&u)
	fmt.Println("eer:", err)
	fmt.Println(u)

	//i, err := session.Model(&User{}).Where("Name=? or Name=?", "aa", "bb").Delete()
	//fmt.Println(i)
	//fmt.Println(err)

	//i, err := session.Model(&User{}).Where("Name=?", "Tom").Update("Age", 30)
	//fmt.Println(i)
	//fmt.Println(err)

}

type User struct {
	Name string
	Age  int
}

func (us *User) AfterQuery(s *sesssion.Session) {
	log.Info("after query")
	us.Age = 5
}
func test_hook() {

	engine, err := goorm.NewEngine("sqlite3", "/Users/zhangmengnan/env/sqlite-tools-osx-x86-3380100/file/gee.db")
	if err != nil {
		log.Error(err)
	}
	session := engine.NewSession()
	var u []User
	err = session.Where("Name=?", "Tom").OrderBy("Age desc").Limit(1).Find(&u)
	fmt.Println("eer:", err)
	fmt.Println(u)

}

func test_transaction() {

	engine, err := goorm.NewEngine("sqlite3", "/Users/zhangmengnan/env/sqlite-tools-osx-x86-3380100/file/gee.db")
	if err != nil {
		log.Error(err)
	}

	session := engine.NewSession()

	//var u []User
	session.Transaction(func(session *sesssion.Session) (interface{}, error) {
		insert, err := session.Insert(&User{Name: "test", Age: 25})
		//i, err := session.Where("Name=?", "lisi").Delete()
		//err = session.Where("Name=?", "Tom").OrderBy("Age desc").Limit(1).Find(&u)
		fmt.Println("eer:", err)
		//fmt.Println(i)
		fmt.Println(insert)
		return nil, errors.New("tt")
	})

}

func main() {

	test_transaction()

}
