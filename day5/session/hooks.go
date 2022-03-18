package sesssion

import (
	"goorm/day2/log"
	"reflect"
)

// 钩子函数

const (
	BeforeQuery = "BeforeQuery"
	AfterQuery  = "AfterQuery"
)

// CallMethod 定义执行钩子函数的 函数
func (s *Session) CallMethod(method string, values interface{}) {

	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method)
	if values != nil {
		fm = reflect.ValueOf(values).MethodByName(method)
	}
	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		if call := fm.Call(param); len(call) > 0 {
			if err, ok := call[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
	return
}
