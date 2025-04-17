package test

import (
	"gym-bro-backend/connection"
	"gym-bro-backend/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Initialize the database connection
	connection.CreateConnection()

	// Run the tests
	m.Run()
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	var router *gin.Engine = gin.Default()
	router.GET("/exercises", controllers.GetExercises)
	router.POST("/exercises", controllers.CreateExercise)
	return router
}

func TestGetExercises(test *testing.T) {
	gin.SetMode(gin.TestMode)
	var router *gin.Engine = gin.Default()
	router.GET("/exercises", controllers.GetExercises)

	req, err := http.NewRequest("GET", "/exercises", nil)

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "exercise")
}

func TestCreateExercise(test *testing.T) {
	var router *gin.Engine = setupRouter()

	var testCases []struct {
		name           string
		payload        string
		expectedStatus int
		expectedBody   string
	} = []struct {
		name           string
		payload        string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Exercise",
			payload:        `{"name": "Push-up", "description": "A basic upper body exercise."}`,
			expectedStatus: http.StatusOK,
			expectedBody:   "Exercise created successfully",
		},
		{
			name:           "Missing Name",
			payload:        `{"description": "A basic upper body exercise."}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request",
		},
		{
			name:           "Invalid JSON",
			payload:        `{"name": "Push-up", "description": "A basic upper body exercise"`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request",
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			req, err := http.NewRequest("POST", "/exercises", strings.NewReader(testCase.payload))
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
