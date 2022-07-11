package entity

type Ptype int64

const (
	TypeReception Ptype = iota
	TypeMaceration
	TypeFermentation
	TypeAgeing
	TypeBottling
)

func (p Ptype) String() string {
	switch p {
	case TypeReception:
		return "reception"
	case TypeMaceration:
		return "maceration"
	case TypeFermentation:
		return "fermentation"
	case TypeAgeing:
		return "ageing"
	case TypeBottling:
		return "bottling"
	default:
		return "unknown"
	}
}
