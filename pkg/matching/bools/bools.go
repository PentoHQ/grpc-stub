package bools

type Operator string

const (
	Equals Operator = "$equals"
	Any    Operator = "$any"
)

type Matcher struct {
	Operator Operator
	Value    bool
}

func (m Matcher) Matches(value bool) bool {
	switch m.Operator {
	case Equals:
		return m.Value == value
	case Any:
		return true
	default:
		return false
	}
}

func MatchAny() Matcher {
	return Matcher{
		Operator: Any,
	}
}

func MatchValue(value bool) Matcher {
	return Matcher{
		Operator: Equals,
		Value:    value,
	}
}
