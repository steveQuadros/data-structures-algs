package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQuickSelect(t *testing.T) {
	type in struct {
		nums []int
		k    int
	}
	tc := []struct {
		name     string
		in       in
		expected int
		f        func([]int, int) int
	}{
		{"k largest 1", in{[]int{4, 2, 1, 3}, 2}, 3, QuickSelectKthLargest},
		{"k largest 2", in{[]int{4, 2, 1, 3}, 1}, 4, QuickSelectKthLargest},
		{"k smallest 1", in{[]int{4, 2, 1, 3}, 1}, 1, QuickSelectKthSmallest},
		{"k smallest 2", in{[]int{4, 2, 1, 3}, 4}, 4, QuickSelectKthSmallest},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.f(tt.in.nums, tt.in.k))
		})
	}
}
