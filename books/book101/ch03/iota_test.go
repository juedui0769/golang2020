package ch03

import "testing"

const x = iota
const y = iota

const (
	a = iota
	b
	c
)

const (
	u = 1 << iota
	v
	w
)

const (
	e, f = iota, 1 << iota
	_, _
	g, h
	i, j
)

func TestIota101(t *testing.T){
	t.Log("x: ", x)
	t.Log("y: ", x)
	t.Logf("a = %d, b = %d, c = %d", a, b, c)
	t.Logf("u = %d, v = %d, w = %d", u, v, w)
	t.Logf("e: %d, g: %d, i: %d", e, g, i)
	t.Logf("f: %d, h: %d, j: %d", f, h, j)
}
