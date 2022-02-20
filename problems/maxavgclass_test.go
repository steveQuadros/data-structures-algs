package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxAvgClass(t *testing.T) {
	classes := [][]int{{1, 2}, {3, 5}, {2, 2}}
	extraStudents := 2
	require.Equal(t, 0.78333, maxAverageRatio(classes, extraStudents))
}
