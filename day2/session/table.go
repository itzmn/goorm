package sesssion

import (
	"fmt"
	"goorm/day2/log"
	"goorm/day2/schema"
	"reflect"
	"strings"
)

// 和表相关的操作封装

// Model 初始化session中的表
func (s *Session) Model(data interface{}) *Session {

	if s.refTable == nil || reflect.TypeOf(data) != reflect.TypeOf(s.refTable) {
		schema := schema.Parse(data, s.dialect)
		s.refTable = schema
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set, refTable is nil")
	}
	return s.refTable
}

// =========表级别的增删改查=========

//CreateTable 创建表
func (s *Session) CreateTable() error {
	refTable := s.RefTable()
	var columns []string
	for _, field := range refTable.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	create := fmt.Sprintf("create table %s (%s);", refTable.Name, desc)
	log.Info(create)
	_, err := s.Raw(create).Exec()
	return err
}

// DropTable 创建表
func (s *Session) DropTable() error {
	refTable := s.RefTable()
	drop := fmt.Sprintf("drop table if exists %s;", refTable.Name)
	log.Info(drop)
	_, err := s.Raw(drop).Exec()
	return err
}

// HasTable 创建表
func (s *Session) HasTable() bool {
	refTable := s.RefTable()
	sql, args := s.dialect.TableExistSql(refTable.Name)
	has := fmt.Sprintf("%s, args:%v", sql, args)
	log.Info(has)
	row := s.Raw(sql, args...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == refTable.Name
}
