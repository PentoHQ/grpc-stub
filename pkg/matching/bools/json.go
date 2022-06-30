package bools

import (
	"errors"

	"github.com/pentohq/grpc-stub/pkg/jsonstrict"
)

var ErrInvalidJSON = errors.New("cannot unmarshal json value into bools.Matcher")

type anyMatcherJSON struct {
	Any struct{} `json:"$any"`
}

func (m *Matcher) UnmarshalJSON(data []byte) error {
	var value bool
	if err := jsonstrict.Unmarshal(data, &value); err == nil {
		*m = MatchValue(value)
		return nil
	}

	var anyMatcher anyMatcherJSON
	if err := jsonstrict.Unmarshal(data, &anyMatcher); err == nil {
		*m = MatchAny()
		return nil
	}

	return ErrInvalidJSON
}
