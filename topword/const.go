package topword

import "strings"

const (
	N = 10
)

var (
	filter = strings.NewReplacer(",", "", ".", "")
)
