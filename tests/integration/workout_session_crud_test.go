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

func TestGetWorkoutSessions(testing *testing.T) {
	var router *gin.Engine = setupRouter()
	router.GET("/workout_sessions", controllers.GetWorkoutSessions)

	req, err := http.NewRequest("GET", "/workout_sessions", nil)
	if err != nil {
		testing.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(testing, http.StatusOK, responseRecorder.Code)
	assert.Contains(testing, responseRecorder.Body.String(), "Workout sessions retrieved successfully")
}

func TestGetWorkoutSessionByID(testing *testing.T) {
	var router *gin.Engine = setupRouter()
	router.GET("/workout_sessions/:id", controllers.GetWorkoutSessionByID)

	var workoutSession controllers.WorkoutSession = createTestWorkoutSession()
	req, err := http.NewRequest("GET", "/workout_sessions/"+fmt.Sprint(workoutSession.ID), nil)
	if err != nil {
		testing.Fatal(err)
	}
	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(testing, http.StatusOK, responseRecorder.Code)
	assert.Contains(testing, responseRecorder.Body.String(), "Workout session retrieved successfully")
}
