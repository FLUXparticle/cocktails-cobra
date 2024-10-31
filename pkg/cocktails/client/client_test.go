package client

import (
	"testing"
)

func BenchmarkSequential(b *testing.B) {
	sumMilk := DoSequential()

	if sumMilk != 38 {
		b.Errorf("sumMilk: %d expected: 38", sumMilk)
	}
}

func BenchmarkParallel(b *testing.B) {
	sumMilk := DoParallel()

	if sumMilk != 38 {
		b.Errorf("sumMilk: %d expected: 38", sumMilk)
	}
}
