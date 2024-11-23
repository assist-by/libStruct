package libStruct

type TechnicalIndicators struct {
	EMA200       float64
	ParabolicSAR float64
	MACDLine     float64
	SignalLine   float64
}

func NewTechnicalIndicators(ema200, parabolicSar, macdLine, signalLine float64) (*TechnicalIndicators, error) {
	return &TechnicalIndicators{
		EMA200:       ema200,
		ParabolicSAR: parabolicSar,
		MACDLine:     macdLine,
		SignalLine:   signalLine,
	}, nil
}
