package main

import (
	"encoding/json"
	. "fmt"
	"net/http"
)

type Student struct {
	Name  string `json:"nama"`
	Id    int    `json:"nis"`
	Grade int    `json:"nilai"`
}

var students = []Student{
	{"Ali", 201, 67},
	{"Loren", 202, 67},
	{"Budhi", 203, 67},
	{"Jessar", 204, 67},
	{"Krisna", 205, 67},
	{"Aziz", 206, 67},
}

func studentsData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var jsonData, err = json.Marshal(students)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
		return
	}

	http.Error(w, "Gagal mengambil data siswa", http.StatusBadRequest)
}

func student(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	if req.Method == "POST" {
		var studentId = req.FormValue("id")
		var result []byte
		var err error

		for _, student := range students {
			if Sprint(student.Id) == studentId {
				result, err = json.Marshal(student)
				if err != nil {
					http.Error(res, err.Error(), http.StatusInternalServerError)
					return
				}

				res.Write(result)
				return
			}
		}
	}

	http.Error(res, "Bad Request :(", http.StatusBadRequest)
}
func main() {
	http.HandleFunc("/students", studentsData)
	http.HandleFunc("/student", student)

	Println("Server running at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
