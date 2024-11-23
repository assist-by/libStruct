package libStruct

type SignalResult struct {
	Symbol     string
	Signal     string
	Timestamp  int64
	Price      float64
	StopLoss   float64
	TakeProfit float64
	Conditions SignalConditions
}
