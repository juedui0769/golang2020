package ch04

import (
	"reflect"
	"testing"
)

type myKeys struct {
	container []interface{}
	compareFunc func(interface{}, interface{}) int8
	elemType reflect.Type
}

func TestMyKeys01(t *testing.T) {
	int64Keys := &myKeys{
		container:   make([]interface{}, 0),
		compareFunc: func(e1 interface{}, e2 interface{}) int8 {
			k1 := e1.(int64)
			k2 := e2.(int64)
			if k1 < k2 {
				return -1
			} else if k1 > k2 {
				return 1
			} else {
				return 0
			}
		},
		elemType:    reflect.TypeOf(int64(1)),
	}
	_ = int64Keys
}



