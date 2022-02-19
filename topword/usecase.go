package topword

import (
	"context"
	"strings"

	"github.com/hysem/top-word-service/shared/heap"
)

// Usecase interface
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

	// Step 2: create a heap with capacity N (min heap because of the Less method defined for WordInfo)
	h := heap.New[*WordInfo](N)

	for word, count := range wordCount {
		w := &WordInfo{
			Word:  word,
			Count: count,
		}

		// Step 3: Try to push to the heap. If the heap is full Push will return false
		if !h.Push(w) {

			// Step 4: If the heap is full try comparing the heap top (word with min count) and the current word
			if minWord, _ := h.Peek(); w.IsLess(minWord) {
				// Step 5: If true then replace stack top with the current word
				h.Pop()
				h.Push(w)
			}
		}
	}

	topWords := make([]*WordInfo, 0, N)
	// Step 6: Get the top words from the heap in order; top word will be at the end of the list
	for i := 0; i < N; i++ {
		if w, ok := h.Pop(); ok {
			topWords = append(topWords, w)
		} else {
			break
		}
	}

	// Step 7: return the result
	return topWords
}
