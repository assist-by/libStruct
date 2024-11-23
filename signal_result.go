package libStruct

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
