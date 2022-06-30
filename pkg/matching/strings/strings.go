package strings

import "strings"

type Operator string

const (
	Equals   Operator = "$equals"
	Any      Operator = "$any"
	Contains Operator = "$contains"
)

type Matcher struct {
	Operator Operator
	Value    string
}

func (m Matcher) Matches(value string) bool {
	switch m.Operator {
	case Equals:
		return m.Value == value
	case Any:
		return true
	case Contains:
		return strings.Contains(value, m.Value)
	default:
		return false
	}
}

func MatchAny() Matcher {
	return Matcher{
		Operator: Any,
	}
}

func MatchValue(value string) Matcher {
	return Matcher{
		Operator: Equals,
		Value:    value,
	}
}

func MatchValueContaining(value string) Matcher {
	return Matcher{
		Operator: Contains,
		Value:    value,
	}
}
