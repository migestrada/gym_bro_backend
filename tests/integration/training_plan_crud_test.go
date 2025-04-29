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

func TestGetTrainingPlans(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.GET("/training-plans", controllers.GetTrainingPlans)
	req, err := http.NewRequest("GET", "/training-plans", nil)

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Training plans retrieved successfully")
}

func TestGetTrainingPlanByID(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.GET("/training-plans/:id", controllers.GetTrainingPlanByID)

	var trainingPlan controllers.TrainingPlan = createTestTrainingPlan()
	req, err := http.NewRequest("GET", "/training-plans/"+fmt.Sprint(trainingPlan.ID), nil)

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Training plan retrieved successfully")
}

func TestCreateTrainingPlan(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.POST("/training-plans", controllers.CreateTrainingPlan)

	req, err := http.NewRequest("POST", "/training-plans", strings.NewReader(`{"name": "Test Training Plan`+fmt.Sprintf("%d", time.Now().Unix())+`", "description": "This is a test training plan"}`))

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusCreated, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Training plan created successfully")
}

func TestDeleteTrainingPlan(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.DELETE("/training-plans/:id", controllers.DeleteTrainingPlan)

	var trainingPlan controllers.TrainingPlan = createTestTrainingPlan()
	req, err := http.NewRequest("DELETE", "/training-plans/"+fmt.Sprint(trainingPlan.ID), nil)

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Training plan deleted successfully")
}

func TestUpdateTrainingPlan(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.PUT("/training-plans/:id", controllers.UpdateTrainingPlan)

	var trainingPlan controllers.TrainingPlan = createTestTrainingPlan()
	req, err := http.NewRequest("PUT", "/training-plans/"+fmt.Sprint(trainingPlan.ID), strings.NewReader(`{"name": "Updated Training Plan`+fmt.Sprintf("%d", time.Now().UnixMilli())+`", "description": "This is an updated test training plan"}`))

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Training plan updated successfully")
}
