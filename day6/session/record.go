package sesssion

import (
	"fmt"
	"goorm/day1/log"
	"goorm/day5/clause"
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
	//s.CallMethod(BeforeQuery, nil)
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
		s.CallMethod(AfterQuery, dest.Addr().Interface())
		destSlice.Set(reflect.Append(destSlice, dest))
	}

	return raws.Close()
}

func (s *Session) Update(kv ...interface{}) (int64, error) {
	// 传入参数支持两种
	// 1. map[string]interface{}
	// 2. kv list  name tom age 12
	m, ok := kv[0].(map[string]interface{})
	if !ok {
		m = make(map[string]interface{})
		for i := 0; i < len(kv); i += 2 {
			m[kv[i].(string)] = kv[i+1]
		}
	}

	s.clause.Set(clause.UPDATE, s.RefTable().Name, m)
	sql, args := s.clause.Build(clause.UPDATE, clause.WHERE)
	exec, err := s.Raw(sql, args...).Exec()
	if err != nil {
		log.Error("exec sql error:", exec)
		return 0, err
	}
	return exec.RowsAffected()
}

// Delete 根据传入的对象执行删除操作
func (s *Session) Delete() (int64, error) {

	s.clause.Set(clause.DELETE, s.RefTable().Name)
	sql, args := s.clause.Build(clause.DELETE, clause.WHERE)
	exec, err := s.Raw(sql, args...).Exec()
	if err != nil {
		log.Error("exec sql error:", sql)
		return 0, err
	}
	return exec.RowsAffected()
}

func (s *Session) Count() (int64, error) {
	s.clause.Set(clause.COUNT, s.RefTable().Name)
	sql, args := s.clause.Build(clause.COUNT, clause.WHERE)
	row := s.Raw(sql, args...).QueryRow()
	var tmp int64
	if err := row.Scan(&tmp); err != nil {
		return 0, err
	}
	return tmp, nil
}

//===========增加链式操作
// s.Where().Limit().Find()

func (s *Session) Where(desc string, values ...interface{}) *Session {
	var args []interface{}
	s.clause.Set(clause.WHERE, append(append(args, desc), values...)...)
	return s
}

func (s *Session) Limit(limit int) *Session {
	s.clause.Set(clause.LIMIT, limit)
	return s
}

func (s *Session) OrderBy(order string) *Session {
	s.clause.Set(clause.ORDERBY, order)
	return s
}

func (s *Session) First() *Session {
	s.clause.Set(clause.LIMIT, 1)
	return s
}
