package ch03

import "sort"

type Sequence struct {
	len int
	cap int
	Sortable
	sortableArray sort.Interface
}



