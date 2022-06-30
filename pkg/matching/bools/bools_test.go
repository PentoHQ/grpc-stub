package bools_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/pentohq/grpc-stub/pkg/matching/bools"
)

func TestMatchAny(t *testing.T) {
	m := bools.MatchAny()
	v := rand.Int()%2 == 0

	result := m.Matches(v)
	if !result {
		t.Errorf("expected to match value %v", v)
	}
}

func TestMatchValue(t *testing.T) {
	m := bools.MatchValue(false)
	tcs := []struct {
		value    bool
		expected bool
	}{
		{
			value:    true,
			expected: false,
		},
		{
			value:    false,
			expected: true,
		},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v", tc.value), func(t *testing.T) {
			actual := m.Matches(tc.value)

			if actual != tc.expected {
				t.Errorf("expected result to be %v got %v", tc.expected, actual)
			}
		})
	}
}
