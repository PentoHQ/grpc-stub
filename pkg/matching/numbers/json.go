package numbers

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/pentohq/grpc-stub/pkg/jsonstrict"
)

var (
	ErrInvalidJSON = errors.New("cannot unmarshal json value into numbers.Matcher")
)

type anyMatcherJSON struct {
	Any struct{} `json:"$any"`
}

type ltMatcherJSON struct {
	Lt json.Number `json:"$lt"`
}

type lteMatcherJSON struct {
	Lte json.Number `json:"$lte"`
}

type gtMatcherJSON struct {
	Gt json.Number `json:"$gt"`
}

type gteMatcherJSON struct {
	Gte json.Number `json:"$gte"`
}

// TODO(rk): split to smaller functions
func (m *Matcher[T]) UnmarshalJSON(data []byte) error {
	var value json.Number
	err := jsonstrict.Unmarshal(data, &value)
	if err == nil {
		v, err := toNumber[T](json.Number(value))
		if err != nil {
			return fmt.Errorf("%w: %s", ErrInvalidJSON, err)
		}
		*m = MatchValue(v)
		return nil
	}

	var anyMatcher anyMatcherJSON
	if err := jsonstrict.Unmarshal(data, &anyMatcher); err == nil {
		*m = MatchAny[T]()
		return nil
	}

	var ltMatcher ltMatcherJSON
	if err := jsonstrict.Unmarshal(data, &ltMatcher); err == nil {
		v, err := toNumber[T](ltMatcher.Lt)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrInvalidJSON, err)
		}
		*m = MatchValueLt(v)
		return nil
	}

	var lteMatcher lteMatcherJSON
	if err := jsonstrict.Unmarshal(data, &lteMatcher); err == nil {
		v, err := toNumber[T](lteMatcher.Lte)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrInvalidJSON, err)
		}
		*m = MatchValueLte(v)
		return nil
	}

	var gtMatcher gtMatcherJSON
	if err := jsonstrict.Unmarshal(data, &gtMatcher); err == nil {
		v, err := toNumber[T](gtMatcher.Gt)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrInvalidJSON, err)
		}
		*m = MatchValueGt(v)
		return nil
	}

	var gteMatcher gteMatcherJSON
	if err := jsonstrict.Unmarshal(data, &gteMatcher); err == nil {
		v, err := toNumber[T](gteMatcher.Gte)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrInvalidJSON, err)
		}
		*m = MatchValueGte(v)
		return nil
	}

	return ErrInvalidJSON
}

func toNumber[T Number](n json.Number) (T, error) {
	var t T
	switch any(t).(type) {
	case float64:
		v, err := strconv.ParseFloat(n.String(), 64)
		return T(v), err
	case float32:
		v, err := strconv.ParseFloat(n.String(), 32)
		return T(float32(v)), err
	case int32:
		v, err := strconv.ParseInt(n.String(), 10, 32)
		return T(int32(v)), err
	case int64:
		v, err := strconv.ParseInt(n.String(), 10, 64)
		return T(v), err
	case uint32:
		v, err := strconv.ParseUint(n.String(), 10, 32)
		return T(uint32(v)), err
	case uint64:
		v, err := strconv.ParseUint(n.String(), 10, 64)
		return T(v), err
	default:
		return t, fmt.Errorf("cannot use %q as %s", n.String(), reflect.TypeOf(t))
	}
}
