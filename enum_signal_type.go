package autrolibrary

type SignalType int

const (
	LONG SignalType = iota
	SHORT
	MAYBE_LONG
	MAYBE_SHORT
	NO_SIGNAL
)

func (t SignalType) String() string {
	return [...]string{"LONG", "SHORT", "MAYBE LONG", "MAYBE SHORT", "NO SIGNAL"}[t]
}
