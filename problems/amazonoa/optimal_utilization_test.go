package amazonoa

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOU(t *testing.T) {
	type testInput struct {
		a, b   [][]int
		target int
	}
	type testCase struct {
		in       testInput
		expected [][]int
	}
	tc := map[string]testCase{
		"ex 1": {
			in: testInput{
				a:      [][]int{{1, 2}, {2, 4}, {3, 6}},
				b:      [][]int{{1, 2}},
				target: 7,
			},
			expected: [][]int{{2, 1}},
		},
		"ex 2": {
			in: testInput{
				a:      [][]int{{1, 3}, {2, 5}, {3, 7}, {4, 10}},
				b:      [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
				target: 10,
			},
			expected: [][]int{{2, 4}, {3, 2}},
		},
		"ex 3": {
			in: testInput{
				a:      [][]int{{1, 8}, {2, 7}, {3, 14}},
				b:      [][]int{{1, 5}, {2, 10}, {3, 14}},
				target: 20,
			},
			expected: [][]int{{3, 1}},
		},
		"ex 4": {
			in: testInput{
				a:      [][]int{{1, 8}, {2, 15}, {3, 9}},
				b:      [][]int{{1, 8}, {2, 11}, {3, 12}},
				target: 20,
			},
			expected: [][]int{{1, 3}, {3, 2}},
		},
	}

	for name, tt := range tc {
		t.Run("BF: "+name, func(t *testing.T) {
			actual := OptimalUtilization(tt.in.a, tt.in.b, tt.in.target)
			require.ElementsMatch(t, actual, tt.expected)
		})
	}

	for name, tt := range tc {
		t.Run("2P: "+name, func(t *testing.T) {
			actual := OptimalUtilization2P(tt.in.a, tt.in.b, tt.in.target)
			require.ElementsMatch(t, actual, tt.expected)
		})
	}

}
