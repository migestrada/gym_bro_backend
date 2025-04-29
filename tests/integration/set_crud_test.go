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

func TestGetSets(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.GET("/sets", controllers.GetSets)

	req, err := http.NewRequest("GET", "/sets", nil)

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Sets retrieved successfully")
}

func TestCreateTest(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.POST("/sets", controllers.CreateSet)
	router.POST("/exercises", controllers.CreateExercise)
	var exercise controllers.Exercise = createTestExercise()

	req, err := http.NewRequest("POST", "/sets", strings.NewReader(`{"reps": 1, "rest_time": 1, "weight": 1, "weight_unit": "kg", "exercise_id": `+fmt.Sprint(exercise.ID)+`}`))

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusCreated, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Set created successfully")
}

func TestGetSetByID(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.POST("/sets", controllers.CreateSet)
	router.GET("/sets/:id", controllers.GetSetByID)

	var set controllers.Set = createTestSet()

	req, err := http.NewRequest("GET", "/sets/"+fmt.Sprint(set.ID), nil)

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Set retrieved successfully")
}

func TestDeleteSet(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.POST("/sets", controllers.CreateSet)
	router.DELETE("/sets/:id", controllers.DeleteSet)

	var set controllers.Set = createTestSet()

	req, err := http.NewRequest("DELETE", "/sets/"+fmt.Sprint(set.ID), nil)

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Set deleted successfully")
}

func TestUpdateSet(test *testing.T) {
	var router *gin.Engine = setupRouter()
	router.POST("/sets", controllers.CreateSet)
	router.PUT("/sets/:id", controllers.UpdateSet)

	var set controllers.Set = createTestSet()

	req, err := http.NewRequest("PUT", "/sets/"+fmt.Sprint(set.ID), strings.NewReader(`{"reps": 2, "rest_time": 2, "weight": 2, "weight_unit": "kg"}`))

	if err != nil {
		test.Fatal(err)
	}

	var responseRecorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(test, http.StatusOK, responseRecorder.Code)
	assert.Contains(test, responseRecorder.Body.String(), "Set updated successfully")
}
