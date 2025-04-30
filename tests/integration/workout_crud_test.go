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

func TestDeleteWorkout(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.DELETE("/workouts/:id", controllers.DeleteWorkout)

	var workout controllers.Workout = createTestWorkout()
	req, err := http.NewRequest("DELETE", "/workouts/"+fmt.Sprint(workout.ID), nil)

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Workout deleted successfully")
}

func TestCreateWorkout(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.POST("/workouts", controllers.CreateWorkout)

	var trainingPlan controllers.TrainingPlan = createTestTrainingPlan()
	req, err := http.NewRequest("POST", "/workouts", strings.NewReader(`{"name": "Test Workout", "description": "Test Workout Description", "order": 1, "training_plan_id": `+fmt.Sprint(trainingPlan.ID)+`}`))

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusCreated, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Workout created successfully")
}
