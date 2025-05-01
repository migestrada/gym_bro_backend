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

func TestGetWorkoutExerciseByID(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.GET("/workout_exercises/:id", controllers.GetWorkoutExerciseByID)

	var workoutExercise controllers.WorkoutExercise = createTestWorkoutExercise()
	req, err := http.NewRequest("GET", "/workout_exercises/"+fmt.Sprint(workoutExercise.ID), nil)
	if err != nil {
		test.Fatal(err)
	}
	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Workout exercise retrieved successfully")
}

func TestCreateWorkoutExercise(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.POST("/workout_exercises", controllers.CreateWorkoutExercise)

	var workout controllers.Workout = createTestWorkout()
	var exercise controllers.Exercise = createTestExercise()
	req, err := http.NewRequest("POST", "/workout_exercises", strings.NewReader(`{"workout_id":`+fmt.Sprint(workout.ID)+`,"exercise_id":`+fmt.Sprint(exercise.ID)+`}`))
	if err != nil {
		test.Fatal(err)
	}
	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusCreated, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Workout exercise created successfully")
}

func TestUpdateWorkoutExercise(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.PUT("/workout_exercises/:id", controllers.UpdateWorkoutExercise)

	var workoutExercise controllers.WorkoutExercise = createTestWorkoutExercise()
	var workout controllers.Workout = createTestWorkout()
	var exercise controllers.Exercise = createTestExercise()
	req, err := http.NewRequest("PUT", "/workout_exercises/"+fmt.Sprint(workoutExercise.ID), strings.NewReader(`{"workout_id":`+fmt.Sprint(workout.ID)+`,"exercise_id":`+fmt.Sprint(exercise.ID)+`}`))
	if err != nil {
		test.Fatal(err)
	}
	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Workout exercise updated successfully")
}
