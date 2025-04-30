package test

import (
	"fmt"
	"gym-bro-backend/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetWorkouts(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.GET("/workouts", controllers.GetWorkouts)
	req, err := http.NewRequest("GET", "/workouts", nil)

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Workouts retrieved successfully")
}

func TestGetWorkoutByID(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.GET("/workouts/:id", controllers.GetWorkoutByID)

	var workout controllers.Workout = createTestWorkout()
	req, err := http.NewRequest("GET", "/workouts/"+fmt.Sprint(workout.ID), nil)

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Workout retrieved successfully")
}
