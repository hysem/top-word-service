package topword_test

import (
	"context"
	"testing"

	"github.com/hysem/top-word-service/topword"
	"github.com/stretchr/testify/assert"
)

type usecaseMock struct {
}

func (m *usecaseMock) assertExpectations(t *testing.T) {
}
func newUsecase(t *testing.T) (topword.Usecase, *usecaseMock) {
	m := usecaseMock{}
	u := topword.NewUsecase()
	return u, &m
}

func TestUsecase_FindTopWords(t *testing.T) {
	const testParagraph = "paragraph to test test"
	var testResponse = []*topword.WordInfo{{
		Count: 2,
		Word:  "test",
	}, {
		Word:  "to",
		Count: 1,
	}, {
		Word:  "paragraph",
		Count: 1,
	}}
	testCases := map[string]struct {
		expectedResponse []*topword.WordInfo
	}{
		`success case: found top words`: {
			expectedResponse: testResponse,
		},
	}
	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			u, m := newUsecase(t)
			defer m.assertExpectations(t)

			actualResponse := u.FindTopWords(context.Background(), &topword.FindTopWordsRequest{
				Text: testParagraph,
			})

			assert.Equal(t, tc.expectedResponse, actualResponse)
		})
	}
}
