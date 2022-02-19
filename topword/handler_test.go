package topword_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/hysem/top-word-service/topword"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type handlerMocks struct {
	topWordUsecase topword.UsecaseMock
}

func (m *handlerMocks) assertExpectations(t *testing.T) {
	m.topWordUsecase.AssertExpectations(t)
}
func newHandler(t *testing.T) (*topword.Handler, *handlerMocks) {
	m := handlerMocks{}
	h := topword.NewHandler(&m.topWordUsecase)
	return h, &m
}

func TestHandler_FindTopWords(t *testing.T) {
	const testParagraph = "paragraph to test"
	var testResponse = []*topword.WordInfo{{
		Word:  "paragraph",
		Count: 1,
	}, {
		Word:  "to",
		Count: 1,
	}, {
		Word:  "test",
		Count: 1,
	}}
	testCases := map[string]struct {
		requestMethod    string
		requestText      string
		expectedStatus   int
		expectedResponse []*topword.WordInfo
		setMocks         func(m *handlerMocks)
	}{
		`error case: invalid request method`: {
			requestMethod:  http.MethodGet,
			expectedStatus: http.StatusMethodNotAllowed,
		},
		`error case: nil response`: {
			requestMethod: http.MethodPost,
			requestText:   testParagraph,
			setMocks: func(m *handlerMocks) {
				m.topWordUsecase.On("FindTopWords", mock.Anything, &topword.FindTopWordsRequest{
					Text: testParagraph,
				}).Return([]*topword.WordInfo{})
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: []*topword.WordInfo{},
		},
		`success case: found top words`: {
			requestMethod: http.MethodPost,
			requestText:   testParagraph,
			setMocks: func(m *handlerMocks) {
				m.topWordUsecase.On("FindTopWords", mock.Anything, &topword.FindTopWordsRequest{
					Text: testParagraph,
				}).Return(testResponse, nil)
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: testResponse,
		},
	}
	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			h, m := newHandler(t)
			defer m.assertExpectations(t)
			if tc.setMocks != nil {
				tc.setMocks(m)
			}

			form := url.Values{}
			form.Add("text", tc.requestText)

			req, err := http.NewRequest(tc.requestMethod, "/endpoint", nil)
			require.NoError(t, err)
			req.PostForm = form
			rec := httptest.NewRecorder()
			h.FindTopWords(rec, req)

			require.Equal(t, tc.expectedStatus, rec.Result().StatusCode)
			if tc.expectedResponse == nil {
				require.Zero(t, rec.Body.String())
			} else {
				var actualResponse []*topword.WordInfo
				err := json.NewDecoder(rec.Body).Decode(&actualResponse)
				require.NoError(t, err)
				require.Equal(t, actualResponse, tc.expectedResponse)
			}
		})
	}
}
