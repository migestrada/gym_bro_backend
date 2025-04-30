package test

import (
	"gym-bro-backend/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetWorkoutExercises(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.GET("/workout_exercises", controllers.GetWorkoutExercises)
	req, err := http.NewRequest("GET", "/workout_exercises", nil)
	if err != nil {
		test.Fatal(err)
	}
	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Workout exercises retrieved successfully")
}
