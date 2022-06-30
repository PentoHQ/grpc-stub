package numbers_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/pentohq/grpc-stub/pkg/matching/numbers"
)

type testCases[T numbers.Number] []struct {
	value    T
	expected bool
}

func (tcs testCases[T]) Run(t *testing.T, m numbers.Matcher[T]) {
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v", tc.value), func(t *testing.T) {
			actual := m.Matches(tc.value)

			if actual != tc.expected {
				t.Errorf("expected result to be %v got %v", tc.expected, actual)
			}
		})
	}
}

func TestMatchAny(t *testing.T) {
	m := numbers.MatchAny[int32]()
	v := rand.Int31()

	result := m.Matches(v)
	if !result {
		t.Errorf("expected to match value %q", v)
	}
}

func TestMatchValue(t *testing.T) {
	m := numbers.MatchValue[float64](0.5)
	tcs := testCases[float64]{
		{
			value:    0.5,
			expected: true,
		},
		{
			value:    0.2,
			expected: false,
		},
	}

	tcs.Run(t, m)
}

func TestMatchValueLt(t *testing.T) {
	m := numbers.MatchValueLt[uint32](10)
	tcs := testCases[uint32]{
		{
			value:    5,
			expected: true,
		},
		{
			value:    10,
			expected: false,
		},
		{
			value:    15,
			expected: false,
		},
	}

	tcs.Run(t, m)
}

func TestMatchValueLte(t *testing.T) {
	m := numbers.MatchValueLte[uint32](10)
	tcs := testCases[uint32]{
		{
			value:    5,
			expected: true,
		},
		{
			value:    10,
			expected: true,
		},
		{
			value:    15,
			expected: false,
		},
	}

	tcs.Run(t, m)
}

func TestMatchValueGt(t *testing.T) {
	m := numbers.MatchValueGt[uint32](10)
	tcs := testCases[uint32]{
		{
			value:    5,
			expected: false,
		},
		{
			value:    10,
			expected: false,
		},
		{
			value:    15,
			expected: true,
		},
	}

	tcs.Run(t, m)
}

func TestMatchValueGte(t *testing.T) {
	m := numbers.MatchValueGte[int64](10)
	tcs := testCases[int64]{
		{
			value:    5,
			expected: false,
		},
		{
			value:    10,
			expected: true,
		},
		{
			value:    15,
			expected: true,
		},
	}

	tcs.Run(t, m)
}
