package topword

import (
	"context"
	"strings"

	"github.com/hysem/top-word-service/shared/heap"
)

// Usecase interface
//go:generate mockery --name=Usecase --structname=UsecaseMock --filename=usecase_mock.go --inpackage
type Usecase interface {

	// FindTopWords finds the top words in the given text using a min heap
	// If two words have the same count then the word will be selected based on the alphabetic order
	FindTopWords(ctx context.Context, request *FindTopWordsRequest) []*WordInfo
}

// usecase implements Usecase interface
type usecase struct {
}

// NewUsecase returns an instance of usecase implementation
func NewUsecase() *usecase {
	return &usecase{}
}

// FindTopWords finds the top words in the given text using a min heap
// If two words have the same count then the word will be selected based on the alphabetic order
func (u *usecase) FindTopWords(ctx context.Context, request *FindTopWordsRequest) []*WordInfo {
	wordCount := map[string]uint64{}

	// Step 1: count the words
	for _, word := range strings.Fields(request.Text) {
		wordCount[word]++
	}

	// Step 2: create a heap with enough capacity (max heap because of the Less method defined for WordInfo)
	h := heap.New[*WordInfo](len(wordCount))

	for word, count := range wordCount {
		w := &WordInfo{
			Word:  word,
			Count: count,
		}

		h.Push(w)
	}

	topWords := make([]*WordInfo, 0, N)
	// Step 3: Get the top words from the heap in order; top word will be at the start of the list
	for i := 0; i < N; i++ {
		if w, ok := h.Pop(); ok {
			topWords = append(topWords, w)
		} else {
			break
		}
	}

	// Step 4: return the result
	return topWords
}
