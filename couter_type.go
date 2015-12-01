package data

const (
	CounterTypeMin     CounterType = iota
	CounterTypeGAUGE               // keep the pass value
	CounterTypeCOUNTER             // this will calc the value as speed
	CounterTypeMax
)

type CounterType int

func (ct CounterType) String() string {
	switch ct {
	case CounterTypeGAUGE:
		return "GAUGE"
	case CounterTypeCOUNTER:
		return "COUNTER"
	default:
		return "GAUGE"
	}
}
