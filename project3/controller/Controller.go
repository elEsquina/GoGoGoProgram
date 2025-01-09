package controller

import (
	"encoding/json"
	"net/http"
	"project3/repository"
	"project3/model"
	_ "fmt"
	"strconv"
)

var studentRepository = repository.NewStudentRepository() 

func StudentsRouter(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		GetAllStudents(w, r)
	} else if r.Method == http.MethodPost {
		CreateStudent(w, r)
	} else {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
	}
}


func StudentsPathParamRouter(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		GetStudentById(w, r)
	} else if r.Method == http.MethodPut {
		UpdateStudentById(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteStudentById(w, r)
	} else {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
	}
}


func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(studentRepository.GetStudents()); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		return
	}
}


func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student model.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	
	if err := studentRepository.AddStudent(student); err != nil {
		http.Error(w, "Failed to add student "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(student); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}


func GetStudentById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	ret, err := studentRepository.GetStudentById(id)

	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(ret); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		return
	}
}


func UpdateStudentById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	var student model.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := studentRepository.UpdateStudent(student, id); err != nil {
		http.Error(w, "Failed to update student: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(student); err != nil {
		http.Error(w, "Failed to encode student response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}


func DeleteStudentById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	if err := studentRepository.DeleteStudent(id); err != nil {
		http.Error(w, "Failed to delete student", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

