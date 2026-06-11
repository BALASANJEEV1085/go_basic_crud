package routes

import (
	"net/http"

	"crud-api/handlers"
)

func SetupRoutes() {

	http.HandleFunc("/students", handlers.GetStudents)

	http.HandleFunc("/student", handlers.GetStudent)

	http.HandleFunc("/student/create", handlers.CreateStudent)

	http.HandleFunc("/student/update", handlers.UpdateStudent)

	http.HandleFunc("/student/delete", handlers.DeleteStudent)
}