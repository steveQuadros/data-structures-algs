package amazonoa

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumsToTarget(t *testing.T) {
	expected := [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 2},
		{1, 1, 3},
		{1, 4},
		{1, 2, 2},
		{2, 3},
	}
	res := SumToTarget(5)
	require.ElementsMatch(t, expected, res)
}
