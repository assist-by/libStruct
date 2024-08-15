package autrolibrary

type SignalType int

const (
	Long SignalType = iota
	Short
	Maybe_Long
	Maybe_Short
	No_Signal
)

var signalTypeNames = [...]string{
	"LONG", "SHORT", "MAYBE LONG", "MAYBE SHORT", "NO SIGNAL",
}

func (t SignalType) String() string {
	if t < Long || t > No_Signal {
		return "Unknown"
	}
	return signalTypeNames[t]
}
