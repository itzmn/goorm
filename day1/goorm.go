package goorm

import (
	"database/sql"
	"goorm/day1/log"
	sesssion "goorm/day1/session"
)

// 核心交互类
//1. 连接数据库且检验连通性
//2. 创建数据库连接

type Engine struct {
	db *sql.DB
}

func NewEngine(driver string, source string) (e *Engine, err error) {

	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	log.Infof("Connect %s, %s success", driver, source)
	e = &Engine{db: db}
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err  != nil{
		log.Error(err)
		return
	}
	log.Info("Close databases success")
}



func (e *Engine) NewSession() *sesssion.Session {
	return sesssion.New(e.db)
}

