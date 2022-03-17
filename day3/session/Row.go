package sesssion

import (
	"database/sql"
	"goorm/day3/clause"
	"goorm/day3/dialect"
	"goorm/day3/log"
	"goorm/day3/schema"
	"strings"
)

// 对原始的SQL交互进行封装

type Session struct {
	db      *sql.DB         // 原生db
	sql     strings.Builder // 存储sql
	sqlVars []interface{}   // 存储传入的参数

	refTable *schema.Schema  // 表相关
	dialect  dialect.Dialect //session属于那个数据库类型

	clause clause.Clause //  sql构造的对象
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{db: db, dialect: dialect}
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

func (s *Session) Raw(sql string, vars ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, vars...)
	return s
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Infof(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Infof(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Session) QueryRaws() (results *sql.Rows, err error) {
	defer s.Clear()
	log.Infof(s.sql.String(), s.sqlVars)
	if results, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
