package sesssion

import (
	"goorm/day6/log"
)

// 事务操作

func (s *Session) Begin() (err error) {

	log.Info("transaction begin")
	if s.tx, err = s.db.Begin(); err != nil {
		log.Error("begin transaction error")
		return
	}
	return
}

func (s *Session) RollBack() (err error) {
	log.Info("transaction rollback")
	if err = s.tx.Rollback(); err != nil {
		return
	}
	return
}

func (s *Session) Commit() (err error) {
	log.Info("transaction commit")
	if err = s.tx.Commit(); err != nil {
		log.Error("transaction commit error")
		return
	}
	return
}

type TransactionFunc func(session *Session) (interface{}, error)

func (s *Session) Transaction(f TransactionFunc) (result interface{}, err error) {
	err = s.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			s.RollBack()
			panic(p)
		} else if err != nil {
			s.RollBack()
		} else {
			s.Commit()
		}

	}()
	return f(s)
}
