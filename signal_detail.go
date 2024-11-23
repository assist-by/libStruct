package libStruct

type SignalDetail struct {
	EMA200Condition       bool
	ParabolicSARCondition bool
	MACDCondition         bool
	EMA200Value           float64
	EMA200Diff            float64
	ParabolicSARValue     float64
	ParabolicSARDiff      float64
	MACDMACDLine          float64
	MACDSignalLine        float64
	MACDHistogram         float64
}

// func New(ema200Condition, parabolicSarCondition, macdCondition bool, ema200Value, ema200Diff, parabolicSarValue, parabolicSarDiff, macdLine, signalLine, macdHistogram float64) (*SignalDetail, error) {
// 	return &SignalDetail{}, nil
// }
