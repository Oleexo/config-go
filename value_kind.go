package config

type ValueKind int

const (
	KindString ValueKind = iota
	KindInt
	KindFloat
	KindBool
)

func (k ValueKind) String() string {
	switch k {
	case KindString:
		return "string"
	case KindInt:
		return "int"
	case KindFloat:
		return "float"
	case KindBool:
		return "bool"
	default:
		return "unknown"
	}
}
