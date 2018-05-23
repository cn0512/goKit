package yktimer

import (
	"fmt"
	"testing"
)

var x, y int

func Add(a, b int) {
	fmt.Println(a + b)
}

func TestSched(t *testing.T) {
	Sched(Add, "23:00:00", "3s", 2, 3)
}

func BenchmarkSched(b *testing.B) {
	x++
	Sched(Add, "23:16:50", "1ms", x, y)
}
