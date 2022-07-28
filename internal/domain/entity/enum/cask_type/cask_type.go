package cask_type

type Ctype int64

const (
	FrenchOak Ctype = iota
	AmericanOak
	SpanishOak
	Acacia
	CherryTree
)

func (c Ctype) String() string {
	switch c {
	case FrenchOak:
		return "French Oak"
	case AmericanOak:
		return "American Oak"
	case SpanishOak:
		return "Spanish Oak"
	case Acacia:
		return "Acacia"
	case CherryTree:
		return "Cherry Tree"
	default:
		return "unknown"
	}
}
