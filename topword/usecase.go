package topword

import (
	"context"
	"strings"

	"github.com/hysem/top-word-service/shared/heap"
)

type Usecase interface {
	FindTopWords(ctx context.Context, request *FindTopWordsRequest) []*WordInfo
}

type usecase struct {
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (u *usecase) FindTopWords(ctx context.Context, request *FindTopWordsRequest) []*WordInfo {
	wordCount := map[string]uint64{}
	for _, word := range strings.Fields(request.Text) {
		wordCount[word]++
	}

	h := heap.New[*WordInfo](N)

	for word, count := range wordCount {
		w := &WordInfo{
			Word:  word,
			Count: count,
		}

		if !h.Push(w) {
			if minWord, _ := h.Peek(); minWord.IsLess(minWord) {
				h.Pop()
				h.Push(w)
			}
		}
	}

	topWords := make([]*WordInfo, 0, N)
	for i := 0; i < N; i++ {
		if w, ok := h.Pop(); ok {
			topWords = append(topWords, w)
		} else {
			break
		}

	}

	return topWords
}
