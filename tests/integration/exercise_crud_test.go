package test

import (
	"fmt"
	"gym-bro-backend/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetExercises(test *testing.T) {
	var router *gin.Engine = setupRouter()
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
	router.POST("/exercises", controllers.CreateExercise)

	var testCases []TestCase = []TestCase{
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

	RunTests(test, router, "POST", "/exercises", testCases)
}

func TestDeleteExercise(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.POST("/exercises", controllers.CreateExercise)
	router.DELETE("/exercises/:id", controllers.DeleteExercise)

	var exercise controllers.Exercise = createTestExercise()
	deleteRequest, err := http.NewRequest("DELETE", fmt.Sprintf("%s%d", "/exercises/", exercise.ID), nil)
	if err != nil {
		test.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, deleteRequest)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Exercise deleted successfully")
}

func TestUpdateExercise(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.POST("/exercises", controllers.CreateExercise)
	router.PUT("/exercises/:id", controllers.UpdateExercise)

	var payload = `{"name":"Updated exercise", "description": "Updated exercise description"}`
	var exercise controllers.Exercise = createTestExercise()
	updateRequest, err := http.NewRequest("PUT", fmt.Sprintf("%s%d", "/exercises/", exercise.ID), strings.NewReader(payload))

	if err != nil {
		test.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, updateRequest)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Exercise updated successfully")
}
