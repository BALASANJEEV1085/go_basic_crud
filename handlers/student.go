package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"crud-api/database"
	"crud-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CREATE
func CreateStudent(w http.ResponseWriter, r *http.Request) {

	var student models.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	result, err := database.StudentCollection.InsertOne(
		context.Background(),
		student,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// READ ALL
func GetStudents(w http.ResponseWriter, r *http.Request) {

	cursor, err := database.StudentCollection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var students []models.Student

	err = cursor.All(context.Background(), &students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// READ ONE
func GetStudent(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var student models.Student

	err = database.StudentCollection.
		FindOne(
			context.Background(),
			bson.M{"_id": objID},
		).
		Decode(&student)

	if err != nil {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// UPDATE
func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var student models.Student

	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	update := bson.M{
		"$set": bson.M{
			"name":   student.Name,
			"branch": student.Branch,
		},
	}

	result, err := database.StudentCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		update,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// DELETE
func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result, err := database.StudentCollection.DeleteOne(
		context.Background(),
		bson.M{"_id": objID},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}