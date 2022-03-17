package sesssion

import (
	"fmt"
	"goorm/day1/log"
	"goorm/day3/clause"
	"reflect"
)

// 对记录的增删改查操作

// Insert 实现记录的插入操作 多个对象插入
func (s *Session) Insert(values ...interface{}) (int64, error) {

	recordValues := make([]interface{}, 0)
	for _, value := range values {
		// 根据传入的对象 获取对象对应的表
		table := s.Model(value).RefTable()
		// 根据对应的表数据， 构建对应的insert 部分sql $table ($fields)
		tableName := table.Name
		fieldNames := table.FieldNames

		// 将传入的对象字段 调用方法，获取sql
		s.clause.Set(clause.INSERT, tableName, fieldNames)
		// 根据传入对象，获取到对应字段的参数
		recordValues = append(recordValues, table.RecordValues(value))
	}

	// 将参数 也绑定到sql
	s.clause.Set(clause.VALUES, recordValues...)

	// 执行构建sql, 获取到 sql 和参数
	sql, args := s.clause.Build(clause.INSERT, clause.VALUES)
	exec, err := s.Raw(sql, args...).Exec()
	if err != nil {
		log.Errorf("exec sql error:%s", sql)
		return 0, err
	}

	return exec.RowsAffected()
}

// Find 根据传入的对象，获取到值
func (s *Session) Find(values interface{}) error {
	// 获取指针指向的具体的值
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	// 得到值的类型
	destType := destSlice.Type().Elem()
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()

	// 构建查询的sql语句 select $field from $table
	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, args := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	raws, err := s.Raw(sql, args...).QueryRaws()
	if err != nil {
		log.Errorf("exec sql error:%s", sql)
		return err
	}
	// 将查询到的值 分解 写入对象
	for raws.Next() {
		fmt.Println("cnt")
		// 得到实例
		dest := reflect.New(destType).Elem()
		var vs []interface{}
		for _, name := range table.FieldNames {
			i := dest.FieldByName(name).Addr().Interface()
			vs = append(vs, i)
		}

		if err := raws.Scan(vs...); err != nil {
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return raws.Close()
}
