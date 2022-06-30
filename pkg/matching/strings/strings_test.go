package strings_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/pentohq/grpc-stub/pkg/matching/strings"
)

type testCases []struct {
	value    string
	expected bool
}

func (tcs testCases) Run(t *testing.T, m strings.Matcher) {
	for _, tc := range tcs {
		t.Run(tc.value, func(t *testing.T) {
			actual := m.Matches(tc.value)

			if actual != tc.expected {
				t.Errorf("expected result to be %v got %v", tc.expected, actual)
			}
		})
	}
}

func TestMatchAny(t *testing.T) {
	m := strings.MatchAny()
	v := strconv.Itoa(rand.Int())

	result := m.Matches(v)
	if !result {
		t.Errorf("expected to match value %q", v)
	}
}

func TestMatchValue(t *testing.T) {
	m := strings.MatchValue("Damiano")
	tcs := testCases{
		{
			value:    "Damiano",
			expected: true,
		},
		{
			value:    "Dam",
			expected: false,
		},
	}

	tcs.Run(t, m)
}

func TestMatchValueContaining(t *testing.T) {
	m := strings.MatchValueContaining("Dam")
	tcs := testCases{
		{
			value:    "Damiano",
			expected: true,
		},
		{
			value:    "dam",
			expected: false,
		},
		{
			value:    "Rafaello",
			expected: false,
		},
	}

	tcs.Run(t, m)
}
