package test

import (
	"fmt"
	"gym-bro-backend/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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

func TestCreateWorkoutSession(testing *testing.T) {
	var router *gin.Engine = setupRouter()
	router.POST("/workout_sessions", controllers.CreateWorkoutSession)

	var set controllers.Set = createTestSet()
	var workoutExercise controllers.WorkoutExercise = createTestWorkoutExercise()

	req, err := http.NewRequest("POST", "/workout_sessions", strings.NewReader(`{
		"date": "2023-10-01T00:00:00Z",
		"workout_exercise_id":`+fmt.Sprint(workoutExercise.ID)+`,
		"set_id":`+fmt.Sprint(set.ID)+`
	}`))

	if err != nil {
		testing.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(testing, http.StatusCreated, responseRecorder.Code)
	assert.Contains(testing, responseRecorder.Body.String(), "Workout session created successfully")
}

func TestUpdateWorkoutSession(testing *testing.T) {
	var router *gin.Engine = setupRouter()
	router.PUT("/workout_sessions/:id", controllers.UpdateWorkoutSession)

	var workoutSession controllers.WorkoutSession = createTestWorkoutSession()
	var workoutExercise controllers.WorkoutExercise = createTestWorkoutExercise()
	var set controllers.Set = createTestSet()

	req, err := http.NewRequest("PUT", "/workout_sessions/"+fmt.Sprint(workoutSession.ID), strings.NewReader(`{
		"date": "`+fmt.Sprint(time.Now().Format(time.RFC3339))+`",
		"workout_exercise_id":`+fmt.Sprint(workoutExercise.ID)+`,
		"set_id":`+fmt.Sprint(set.ID)+`
	}`))

	if err != nil {
		testing.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(testing, http.StatusOK, responseRecorder.Code)
	assert.Contains(testing, responseRecorder.Body.String(), "Workout session updated successfully")
}

func TestDeleteWorkoutSession(testing *testing.T) {
	var router *gin.Engine = setupRouter()
	router.DELETE("/workout_sessions/:id", controllers.DeleteWorkoutSession)

	var workoutSession controllers.WorkoutSession = createTestWorkoutSession()

	req, err := http.NewRequest("DELETE", "/workout_sessions/"+fmt.Sprint(workoutSession.ID), nil)
	if err != nil {
		testing.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(testing, http.StatusOK, responseRecorder.Code)
	assert.Contains(testing, responseRecorder.Body.String(), "Workout session deleted successfully")
}
