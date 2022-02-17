package topword

// WordInfo holds top word detail
type WordInfo struct {
	Word  string
	Count uint64
}

func (tw *WordInfo) IsLess(tw1 *WordInfo) bool {
	if tw.Count == tw1.Count {
		return tw.Word < tw1.Word
	}
	return tw.Count < tw1.Count
}
