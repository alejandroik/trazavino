package process

type Ptype int64

const (
	Reception Ptype = iota
	Maceration
	Fermentation
	Ageing
	Bottling
)

func (p Ptype) String() string {
	switch p {
	case Reception:
		return "reception"
	case Maceration:
		return "maceration"
	case Fermentation:
		return "fermentation"
	case Ageing:
		return "ageing"
	case Bottling:
		return "bottling"
	default:
		return "unknown"
	}
}
