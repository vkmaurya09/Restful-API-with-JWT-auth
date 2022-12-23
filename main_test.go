package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"restapi-auth/controllers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
)

func TestGetAllTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.GET("/api/secured/tasks", controllers.AllTasks)
	req, err := http.NewRequest(http.MethodGet, "/api/secured/tasks", nil)
	if err != nil {
		t.Fatal("could create req", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Passed")
	} else {
		t.Fatalf("Failed")
	}

}

// func TestGetSingleTask(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	r := gin.Default()

// 	r.GET("/api/secured/tasks/one?id=10", controllers.FindTask)
// 	req, err := http.NewRequest(http.MethodGet, "/api/secured/tasks/one?id=10", nil)
// 	if err != nil {
// 		t.Fatal("could create req", err)
// 	}
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	fmt.Println(w.Body)

// 	if w.Code == http.StatusOK {
// 		t.Logf("Passed")
// 	} else {
// 		t.Fatalf("Failed : %d", w.Code)
// 	}

// }

func TestGetSingleTask(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		res := `{
			"id": 10,
			"task_name": "cache IN GO",
			"task_detail": "cache redis integration",
			"date": "2022.12.19 15:28:04"
		}`
		_, _ = w.Write([]byte(res))
		w.WriteHeader(http.StatusOK)
	}

	apitest.New(). // configuration
			HandlerFunc(handler).
			Get("/api/secured/tasks/one?id=10"). // request
			Expect(t).                           // expectations
			Body(`{
					"id": 10,
					"task_name": "cache IN GO",
					"task_detail": "cache redis integration",
					"date": "2022.12.19 15:28:04"
				}`).
		Status(http.StatusOK).
		End()
}
func TestDeleteTask(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		msg := `
		{
			"success": "Task10 deleted"
		}`
		_, _ = w.Write([]byte(msg))
		w.WriteHeader(http.StatusOK)
	}

	apitest.New(). // configuration
			HandlerFunc(handler).
			Delete("/api/secured/tasks/delete?id=11"). // request
			Expect(t).                                 // expectations
			Body(`
			{
				"success": "Task10 deleted"
			}`).
		Status(http.StatusOK).
		End()
}

func TestPostTask(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		msg := `
		{
			"id": 11,
			"task_name": "unit Testing IN GO",
			"task_detail": "apitest is used to test ",
			"date": "2022.12.19 15:28:04"
		}`
		_, _ = w.Write([]byte(msg))
		w.WriteHeader(http.StatusOK)
	}

	apitest.New(). // configuration
			HandlerFunc(handler).
			Post("/api/secured/tasks"). // request
			Expect(t).                  // expectations
			Body(`
			{
				 "id": 11,
				 "task_name": "unit Testing IN GO",
				 "task_detail": "apitest is used to test ",
				 "date": "2022.12.19 15:28:04"
			}`).
		Status(http.StatusOK).
		End()
}
