package test

import (
	"gym-bro-backend/connection"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	name           string
	payload        string
	expectedStatus int
	expectedBody   string
}

func TestMain(m *testing.M) {
	connection.CreateConnection()
	m.Run()
}

func RunTests(test *testing.T, router *gin.Engine, method string, path string, testCases []TestCase) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			req, err := http.NewRequest(method, path, strings.NewReader(testCase.payload))
			if err != nil {
				test.Fatal(err)
			}

			var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
			router.ServeHTTP(responseRecorder, req)
			assert.Equal(test, testCase.expectedStatus, responseRecorder.Code)
			assert.Contains(test, responseRecorder.Body.String(), testCase.expectedBody)
		})
	}
}
