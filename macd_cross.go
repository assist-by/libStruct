package libStruct

// MACD 크로스 체크를 위한 구조체
type MACDCross struct {
	CurrentMACDLine   float64
	CurrentSignalLine float64
	PrevMACDLine      float64
	PrevSignalLine    float64
}
