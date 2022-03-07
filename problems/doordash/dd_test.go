package doordash

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeliveryOrder(t *testing.T) {
	tc := []struct {
		in  []string
		out bool
	}{
		{[]string{"P1", "P2", "D1", "D2"}, true},
		{[]string{"P1", "D1", "P2", "D2"}, true},
		{[]string{"P1", "D2", "D1", "P2"}, false},
		{[]string{"P1", "D2"}, false},
		{[]string{"P1", "P2"}, false},
		{[]string{"P1", "D1", "D1"}, false},
		{[]string{}, true},
		{[]string{"P1", "P1", "D1"}, false},
		{[]string{"P1", "P1", "D1", "D1"}, false},
		{[]string{"P1", "D1", "P1"}, false},
		{[]string{"P1", "D1", "P1", "D1"}, false},
	}

	for _, tt := range tc {
		t.Run(fmt.Sprintf("%q -> %t", tt.in, tt.out), func(t *testing.T) {
			require.Equal(t, tt.out, DeliveryOrder(tt.in))
		})
	}
}
