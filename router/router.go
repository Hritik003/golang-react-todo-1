package router

import (
	"golang-react-todo-1/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	//mux is used to match the incoming requests to their respective handler functions.

	router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTasks", middleware.DeleteAllTasks).Methods("DELETE", "OPTIONS")
	return router
}
