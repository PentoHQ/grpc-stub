package numbers_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/pentohq/grpc-stub/pkg/matching/numbers"
)

func TestMatcher_Unmarshal(t *testing.T) {
	t.Run("generic", func(t *testing.T) {
		tcs := []struct {
			json    string
			matcher numbers.Matcher[int32]
			error   error
		}{
			{
				json:    `10`,
				matcher: numbers.MatchValue[int32](10),
			},
			{
				json:  `10.15`,
				error: numbers.ErrInvalidJSON,
			},
			{
				json:    `{"$any": {}}`,
				matcher: numbers.MatchAny[int32](),
			},
			{
				json:  `{"$any": 10}`,
				error: numbers.ErrInvalidJSON,
			},
			{
				json:    `{"$lt": 10}`,
				matcher: numbers.MatchValueLt[int32](10),
			},
			{
				json:  `{"$lt": {}}`,
				error: numbers.ErrInvalidJSON,
			},
			{
				json:    `{"$lte": 10}`,
				matcher: numbers.MatchValueLte[int32](10),
			},
			{
				json:    `{"$gt": 10}`,
				matcher: numbers.MatchValueGt[int32](10),
			},
			{
				json:    `{"$gte": 10}`,
				matcher: numbers.MatchValueGte[int32](10),
			},
			{
				json:  `{"$foo": {}}`,
				error: numbers.ErrInvalidJSON,
			},
		}

		for _, tc := range tcs {
			t.Run(tc.json, func(t *testing.T) {
				actual := &numbers.Matcher[int32]{}

				err := json.Unmarshal([]byte(tc.json), &actual)
				if tc.error == nil && err != nil {
					t.Fatalf("expected no error got %v", err)
				} else if !errors.Is(err, tc.error) {
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
	})

	t.Run("types", func(t *testing.T) {
		t.Run("float64", func(t *testing.T) {
			j := `0.12345678912121212`
			var expected float64 = 0.12345678912121212
			m := numbers.Matcher[float64]{}
			err := json.Unmarshal([]byte(j), &m)
			if err != nil {
				t.Fatal(err)
			}

			if m.Value != expected {
				t.Errorf("expected value to be %v got %v", expected, m.Value)
			}
		})

		t.Run("float32", func(t *testing.T) {
			j := `0.12345679`
			var expected float32 = 0.12345679
			m := numbers.Matcher[float32]{}
			err := json.Unmarshal([]byte(j), &m)
			if err != nil {
				t.Fatal(err)
			}

			if m.Value != expected {
				t.Errorf("expected value to be %v got %v", expected, m.Value)
			}
		})

		t.Run("int32", func(t *testing.T) {
			j := `-2147483648`
			var expected int32 = -2147483648
			m := numbers.Matcher[int32]{}
			err := json.Unmarshal([]byte(j), &m)
			if err != nil {
				t.Fatal(err)
			}

			if m.Value != expected {
				t.Errorf("expected value to be %v got %v", expected, m.Value)
			}
		})

		t.Run("int64", func(t *testing.T) {
			j := `-9223372036854775808`
			var expected int64 = -9223372036854775808
			m := numbers.Matcher[int64]{}
			err := json.Unmarshal([]byte(j), &m)
			if err != nil {
				t.Fatal(err)
			}

			if m.Value != expected {
				t.Errorf("expected value to be %v got %v", expected, m.Value)
			}
		})

		t.Run("uint32", func(t *testing.T) {
			j := `4294967295`
			var expected uint32 = 4294967295
			m := numbers.Matcher[uint32]{}
			err := json.Unmarshal([]byte(j), &m)
			if err != nil {
				t.Fatal(err)
			}

			if m.Value != expected {
				t.Errorf("expected value to be %v got %v", expected, m.Value)
			}
		})

		t.Run("uint64", func(t *testing.T) {
			j := `18446744073709551615`
			var expected uint64 = 18446744073709551615
			m := numbers.Matcher[uint64]{}
			err := json.Unmarshal([]byte(j), &m)
			if err != nil {
				t.Fatal(err)
			}

			if m.Value != expected {
				t.Errorf("expected value to be %v got %v", expected, m.Value)
			}
		})
	})
}
