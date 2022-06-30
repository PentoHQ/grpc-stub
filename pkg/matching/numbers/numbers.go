package numbers

type Number interface {
	~float64 | ~float32 |
		~int32 | ~int64 | ~uint32 | ~uint64
}

type Operator string

const (
	Equals Operator = "$equals"
	Any    Operator = "$any"
	Lt     Operator = "$lt"
	Lte    Operator = "$lte"
	Gt     Operator = "$gt"
	Gte    Operator = "$gte"
)

type Matcher[T Number] struct {
	Operator Operator
	Value    T
}

func (m Matcher[T]) Matches(value T) bool {
	switch m.Operator {
	case Equals:
		return value == m.Value
	case Any:
		return true
	case Lt:
		return value < m.Value
	case Lte:
		return value <= m.Value
	case Gt:
		return value > m.Value
	case Gte:
		return value >= m.Value
	default:
		return false
	}
}

func MatchAny[T Number]() Matcher[T] {
	return Matcher[T]{
		Operator: Any,
	}
}

func MatchValue[T Number](value T) Matcher[T] {
	return Matcher[T]{
		Operator: Equals,
		Value:    value,
	}
}

func MatchValueLt[T Number](value T) Matcher[T] {
	return Matcher[T]{
		Operator: Lt,
		Value:    value,
	}
}

func MatchValueLte[T Number](value T) Matcher[T] {
	return Matcher[T]{
		Operator: Lte,
		Value:    value,
	}
}

func MatchValueGt[T Number](value T) Matcher[T] {
	return Matcher[T]{
		Operator: Gt,
		Value:    value,
	}
}

func MatchValueGte[T Number](value T) Matcher[T] {
	return Matcher[T]{
		Operator: Gte,
		Value:    value,
	}
}
