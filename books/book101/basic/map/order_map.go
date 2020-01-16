package _map

import (
	"reflect"
	"sort"
)

type OrderedMap struct {
	keys []interface{}
	m map[interface{}]interface{}
}

func (omap *OrderedMap) Len() int {
	return len(omap.keys)
}

func (omap *OrderedMap) Less(i, j int) bool {
	return false
}

func (omap *OrderedMap) Swap(i, j int) {
	omap.keys[i], omap.keys[j] = omap.keys[j], omap.keys[i]
}

type Keys interface {
	// 在Keys接口类型中嵌入sort.Interface接口类型，就意味着Keys类型的值一定是可排序的。
	sort.Interface
	Add(k interface{}) bool
	Remove(k interface{}) bool
	Clear()
	Get(index int) interface{}
	GetAll() []interface{}
	Search(k interface{}) (index int, contains bool)
	ElemType() reflect.Type
	// 通过把比较两个元素值大小的问题抛给使用者，
	// 我们既解决了需要动态确定元素类型的问题，
	// 又明确了比较两个元素值大小的解决方式。
	CompareFunc() func(interface{}, interface{}) int8
}

type myKeys struct {
	container []interface{}
	compareFunc func(interface{}, interface{}) int8
	elemType reflect.Type
}




