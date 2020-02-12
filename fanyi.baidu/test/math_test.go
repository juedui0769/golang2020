package test

import (
	"math/rand"
	"testing"
)

func TestMathRand(t *testing.T) {
	t.Log(rand.Intn(100))
	t.Log(rand.Intn(10000))

	// 每次执行都是一样的
	t.Log(rand.Int())

	// 每次执行都是一样的
	t.Log(rand.Float32())
}

