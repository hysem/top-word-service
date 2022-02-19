package client_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hysem/top-word-service/client"
	"github.com/hysem/top-word-service/topword"
	"github.com/stretchr/testify/require"
)

func TestClient_FindTopWords(t *testing.T) {
	testCases := map[string]struct {
		responseStatusCode int
		expectedErr        string
		expectedResponse   []*topword.WordInfo
	}{
		`error case: invalid response`: {
			responseStatusCode: http.StatusInternalServerError,
			expectedErr:        `invalid response: 500`,
		},
		`error case: failed to parse response body`: {
			responseStatusCode: http.StatusOK,
			expectedErr:        `failed to decode response: EOF`,
		},
		`success case: got response`: {
			responseStatusCode: http.StatusOK,
			expectedResponse: []*topword.WordInfo{
				{Word: "test", Count: 1},
			},
		},
	}
	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				res.WriteHeader(tc.responseStatusCode)
				if tc.expectedResponse != nil {
					require.NoError(t, json.NewEncoder(res).Encode(tc.expectedResponse))
				}
			}))

			c := client.NewClient(s.URL)

			actualResponse, actualErr := c.FindTopWords(context.Background(), &topword.FindTopWordsRequest{
				Text: "name",
			})
			if tc.expectedErr == "" {
				require.NoError(t, actualErr)
			} else {
				require.Contains(t, actualErr.Error(), tc.expectedErr)
			}
			require.Equal(t, tc.expectedResponse, actualResponse)

		})
	}
}
