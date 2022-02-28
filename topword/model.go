package topword

// WordInfo holds top word detail
type WordInfo struct {
	Word  string `json:"word"`
	Count uint64 `json:"count"`
}

// IsLess implements the _comparable interface in heap package
func (tw *WordInfo) IsLess(tw1 *WordInfo) bool {
	if tw.Count == tw1.Count {
		return tw.Word > tw1.Word
	}
	return tw.Count > tw1.Count
}
