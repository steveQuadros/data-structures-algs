package main

import (
	"testing"
	"github.com/stretchr/testify/require"
	"fmt"
)



func TestBisectLeft(t *testing.T) {
	tc := []struct{
		nums []int
		target int
		expected int
	} {
		{[]int{1,3,4,5,6}, 2, 1},
		{[]int{0,10,20,30,40}, 25, 3},
		{[]int{1, 3, 4, 4, 4, 6, 7}, 5, 5},
		{[]int{1, 3, 4, 4, 4, 6, 7}, 2, 1},
	}
	
	for _, tt := range tc {
		t.Run(fmt.Sprintf("%v, %v", tt.nums, tt.target), func(t *testing.T) {
			require.Equal(t, tt.expected, BisectLeft(tt.nums, tt.target), )
		})
	}
}
