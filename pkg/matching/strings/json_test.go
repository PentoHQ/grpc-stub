package strings_test

import (
	"encoding/json"
	"testing"

	"github.com/pentohq/grpc-stub/pkg/matching/strings"
)

func TestMatcher_Unmarshal(t *testing.T) {
	tcs := []struct {
		json    string
		matcher strings.Matcher
		error   error
	}{
		{
			json:    `"foo"`,
			matcher: strings.MatchValue("foo"),
		},
		{
			json:  `1`,
			error: strings.ErrInvalidJSON,
		},
		{
			json:    `{"$any":{}}`,
			matcher: strings.MatchAny(),
		},
		{
			json:  `{"$any": false}`,
			error: strings.ErrInvalidJSON,
		},
		{
			json:    `{"$contains":"bar"}`,
			matcher: strings.MatchValueContaining("bar"),
		},
		{
			json:  `{"$contains":{}}`,
			error: strings.ErrInvalidJSON,
		},
		{
			json:  `{"foo": "bar"}`,
			error: strings.ErrInvalidJSON,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.json, func(t *testing.T) {
			actual := &strings.Matcher{}

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
