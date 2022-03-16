package sesssion

import (
	"database/sql"
	"goorm/day1/log"
	"strings"
)

// 对原始的SQL交互进行封装

type Session struct {
	db      *sql.DB         // 原生db
	sql     strings.Builder // 存储sql
	sqlVars []interface{}   // 存储传入的参数
}

func New(db *sql.DB) *Session {
	return &Session{db: db}
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
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
