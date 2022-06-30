package bools_test

import (
	"encoding/json"
	"testing"

	"github.com/pentohq/grpc-stub/pkg/matching/bools"
)

func TestMatcher_Unmarshal(t *testing.T) {
	tcs := []struct {
		json    string
		matcher bools.Matcher
		error   error
	}{
		{
			json:    `true`,
			matcher: bools.MatchValue(true),
		},
		{
			json:    `{"$any":{}}`,
			matcher: bools.MatchAny(),
		},
		{
			json:  `"foo"`,
			error: bools.ErrInvalidJSON,
		},
		{
			json:  `{"foo": "bar"}`,
			error: bools.ErrInvalidJSON,
		},
		{
			json:  `{"$any": false}`,
			error: bools.ErrInvalidJSON,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.json, func(t *testing.T) {
			actual := &bools.Matcher{}

			err := json.Unmarshal([]byte(tc.json), actual)
			if err != tc.error {
				t.Fatalf("expected error %v got %v", tc.error, err)
			}

			if actual.Operator != tc.matcher.Operator {
				t.Errorf("expected operator to be %s got %s", tc.matcher.Operator, actual.Operator)
			}

			if actual.Value != tc.matcher.Value {
				t.Errorf("expected value to be %v got %v", tc.matcher.Value, actual.Value)
			}
		})
	}
}
