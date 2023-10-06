package utils

import (
	"fmt"
	"testing"
)

var examples = []struct {
	num      int64
	expected int64
}{
	{0, 0},
	{1, 1},
	{25, 325},
	{100, 5050},
	{1000, 500500},
	{4998, 12492501},
	{9999, 49995000},
	{123333222, 7605541886117253},
}

func TestConsecutiveSumConcurrent(t *testing.T) {
	fmt.Println("\nTesting ConsecutiveSumAsync.")

	for i, value := range examples {
		sum := ConsecutiveSumConcurrent(value.num)

		if sum != value.expected {
			t.Errorf("Incorrect sum, it has %d and was expected %d", sum, value.expected)
		} else {
			fmt.Printf("Test #%d success\n", i+1)
		}
	}
}

func TestConsecutiveSum(t *testing.T) {
	fmt.Println("\nTesting ConsecutiveSum (sync).")

	for i, value := range examples {
		sum := ConsecutiveSum(value.num)

		if sum != value.expected {
			t.Errorf("Incorrect sum, it has %d and was expected %d", sum, value.expected)
		} else {
			fmt.Printf("Test #%d success\n", i+1)
		}
	}
}
