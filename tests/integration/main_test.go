package test

import (
	"fmt"
	"gym-bro-backend/connection"
	"gym-bro-backend/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	var router *gin.Engine = gin.Default()
	return router
}

func createTestExercise() controllers.Exercise {
	exercise := controllers.Exercise{
		Name:        "Push-up",
		Description: "A basic upper body exercise.",
	}

	if err := connection.DB.Create(&exercise).Error; err != nil {
		panic("Failed to create test exercise: " + err.Error())
	}

	return exercise
}

func createTestSet() controllers.Set {
	var exercise controllers.Exercise = createTestExercise()
	set := controllers.Set{
		Reps:       10,
		RestTime:   60,
		Weight:     70.0,
		WeightUnit: "kg",
		ExerciseID: exercise.ID,
	}

	if err := connection.DB.Create(&set).Error; err != nil {
		panic("Failed to create test set: " + err.Error())
	}

	return set
}

func createTestTrainingPlan() controllers.TrainingPlan {
	var trainingPlan controllers.TrainingPlan = controllers.TrainingPlan{
		Name:        "Full Body Workout" + fmt.Sprintf("%d", time.Now().UnixMilli()),
		Description: "A comprehensive workout plan.",
	}

	if err := connection.DB.Create(&trainingPlan).Error; err != nil {
		panic("Failed to create test training plan: " + err.Error())
	}

	return trainingPlan
}

func createTestWorkout() controllers.Workout {
	var trainingPlan controllers.TrainingPlan = createTestTrainingPlan()

	var workout controllers.Workout = controllers.Workout{
		Name:           "Leg Day " + fmt.Sprintf("%d", time.Now().UnixMilli()),
		Order:          1,
		TrainingPlanID: trainingPlan.ID,
	}

	if err := connection.DB.Create(&workout).Error; err != nil {
		panic("Failed to create test workout: " + err.Error())
	}

	return workout
}
