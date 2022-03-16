package main

import (
	"fmt"
	"go/ast"
	"goorm/day2/dialect"
	_ "goorm/day2/dialect"
	"goorm/day2/schema"
	"reflect"
)

type Php struct {
	name string
	age int
}


type Java struct {
	Name string
	Php
	Age int
}

func (j *Java) Say() {
	fmt.Println("haha")
}

func testReflect() {
	j:= &Java{Name: "ja", Age: 12}

	of := reflect.ValueOf(j)
	t := reflect.Indirect(of).Type()

	fmt.Println(t.Field(1).Anonymous)
	fmt.Println(ast.IsExported(t.Field(1).Name))
	fmt.Println(t.Field(1).Type)
	fmt.Println(t.Field(1).Name)
	fmt.Println(t.Name())
	fmt.Println(t.NumField())
}

func main() {

	//testReflect()

	dialect, _ := dialect.GetDialect("sqlite3")

	schema := schema.Parse(&Java{}, dialect)
	fmt.Println(schema)

}
