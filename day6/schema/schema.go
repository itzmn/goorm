package schema

import (
	"go/ast"
	"goorm/day3/dialect"
	"reflect"
)

// ORM框架 将表和对象的转换关系表述

type Field struct {
	Name string
	Type string
	Tag  string
}

// Schema 表和对象的映射关系
type Schema struct {
	Model      interface{}       // 对应的对象
	Name       string            // 对象名称，后续的表名
	Fields     []*Field          // 字段列表
	FieldNames []string          // 字段名称列表
	FieldMap   map[string]*Field // 字段名称和字段的映射关系 减少遍历
}

func (s *Schema) GetField(name string) *Field {
	return s.FieldMap[name]
}

// Parse 将对象 根据不同的数据库类型 转换成schema
func Parse(dest interface{}, dialect dialect.Dialect) *Schema {
	of := reflect.ValueOf(dest)
	model := reflect.Indirect(of).Type()
	schema := &Schema{Model: dest, Name: model.Name(), FieldMap: make(map[string]*Field)}
	for i := 0; i < model.NumField(); i++ {
		f := model.Field(i)
		// 如果字段不是嵌入式字段 且 是可以导出的
		if !f.Anonymous && ast.IsExported(f.Name) {
			field := &Field{Name: f.Name, Type: dialect.DataTypeOf(reflect.Indirect(reflect.New(f.Type)))}
			// 找到字段tag
			if lookup, ok := f.Tag.Lookup("goorm:"); ok {
				field.Tag = lookup
			}
			// 将字段 保存入schema
			schema.FieldNames = append(schema.FieldNames, f.Name)
			schema.Fields = append(schema.Fields, field)
			schema.FieldMap[f.Name] = field
		}
	}
	return schema
}

// RecordValues 将传入的对象， 得到所有字段的值的 切片
func (s *Schema) RecordValues(desc interface{}) []interface{} {
	value := reflect.Indirect(reflect.ValueOf(desc))
	var fieldValues []interface{}

	for _, name := range s.FieldNames {
		byName := value.FieldByName(name).Interface()
		fieldValues = append(fieldValues, byName)
	}
	return fieldValues
}
