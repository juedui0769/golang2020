package ch03

import (
	"sort"
	"testing"
)

type SortableStrings [3]string

func (self SortableStrings) Len() int {
	return len(self)
}

func (self SortableStrings) Less(i, j int) bool {
	return self[i] < self[j]
}

func (self SortableStrings) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func Test101(t *testing.T) {
	_, ok := interface{}(SortableStrings{}).(sort.Interface)
	if ok {
		t.Log("True")
	} else {
		t.Log("False")
	}
}
