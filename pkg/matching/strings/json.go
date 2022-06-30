package strings

import (
	"errors"

	"github.com/pentohq/grpc-stub/pkg/jsonstrict"
)

var ErrInvalidJSON = errors.New("cannot unmarshal json value into strings.Matcher")

type anyMatcherJSON struct {
	Any struct{} `json:"$any"`
}

type containsMatcherJSON struct {
	Contains string `json:"$contains"`
}

func (m *Matcher) UnmarshalJSON(data []byte) error {
	var value string
	if err := jsonstrict.Unmarshal(data, &value); err == nil {
		*m = MatchValue(value)
		return nil
	}

	var anyMatcher anyMatcherJSON
	if err := jsonstrict.Unmarshal(data, &anyMatcher); err == nil {
		*m = MatchAny()
		return nil
	}

	var containsMatcher containsMatcherJSON
	if err := jsonstrict.Unmarshal(data, &containsMatcher); err == nil {
		*m = MatchValueContaining(containsMatcher.Contains)
		return nil
	}

	return ErrInvalidJSON
}
