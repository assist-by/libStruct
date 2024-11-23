package libStruct

type SignalType int

const (
	SIGNAL_LONG SignalType = iota + 1
	SIGNAL_SHORT
	SIGNAL_MAYBE_LONG
	SIGNAL_MAYBE_SHORT
	SIGNAL_NO_SIGANL
)

var signalTypeNames = [...]string{
	"LONG", "SHORT", "MAYBE LONG", "MAYBE SHORT", "NO SIGNAL",
}

func (t SignalType) String() string {
	idx := t - 1
	if idx < SIGNAL_LONG || idx > SIGNAL_NO_SIGANL {
		return "Unknown"
	}
	return signalTypeNames[idx]
}

type SignalResult struct {
	Symbol     string
	Signal     SignalType
	Timestamp  int64
	Price      float64
	StopLoss   float64
	TakeProfit float64
	Conditions SignalConditions
}

func NewSignalResult(symbol string, signal SignalType, timestamp int64, price, stoploss, takeProfit float64, conditions SignalConditions) (*SignalResult, error) {
	return &SignalResult{
		Symbol:     symbol,
		Signal:     signal,
		Timestamp:  timestamp,
		Price:      price,
		StopLoss:   stoploss,
		TakeProfit: takeProfit,
		Conditions: conditions,
	}, nil
}
