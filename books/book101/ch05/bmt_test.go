package ch05

import (
	"testing"
	"time"
)

func Benchmark(b *testing.B) {
	customTimerTag := false
	if customTimerTag {
		b.StopTimer()
	}
	time.Sleep(time.Second)
	if customTimerTag {
		b.StartTimer()
	}
}


