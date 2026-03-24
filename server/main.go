package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Student represents a student in the directory.
// Struct tags control the JSON field names.
type Student struct {
	Name  string `json:"name"`
	Major string `json:"major"`
	Year  int    `json:"year"`
}

// Go does not allow unused imports. This line references the json package so the
// compiler accepts the import before you use it in your implementation.
// Remove it once you call json functions in your handler.
var (
	_ = json.Marshal
	_ = strings.TrimPrefix
)

// In-memory store: name -> Student
var students = map[string]Student{
	"alice": {Name: "alice", Major: "Computer Science", Year: 3},
	"bob":   {Name: "bob", Major: "Mathematics", Year: 2},
}

// handleGetStudent handles GET /students/{name}
func handleGetStudent(w http.ResponseWriter, r *http.Request) {
	// TODO 1: Check that the HTTP method is GET.
	//         If not: http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) and return.

	// TODO 2: Extract the student name from the URL path.
	//         strings.TrimPrefix(r.URL.Path, "/students/") returns the name.
	//         If empty: http.Error(w, "...", http.StatusBadRequest) and return.

	// TODO 3: Look up the student in the "students" map.
	//         Use the comma-ok idiom: student, ok := students[name]
	//         If not found: http.Error(w, "...", http.StatusNotFound) and return.

	// TODO 4: Set the Content-Type header to "application/json".
	//         w.Header().Set("Content-Type", "application/json")

	// TODO 5: Encode the student as JSON and write it into the response.
	//         json.NewEncoder(w).Encode(student)

	http.Error(w, "not implemented", http.StatusNotImplemented)
}

// handleAddStudent handles POST /students
func handleAddStudent(w http.ResponseWriter, r *http.Request) {
	// TODO 1: Check that the HTTP method is POST.

	// TODO 2: Decode the JSON body into a Student struct.
	//         var student Student
	//         err := json.NewDecoder(r.Body).Decode(&student)
	//         Don't forget to check the error!

	// TODO 3: Save the student in the map: students[student.Name] = student

	// TODO 4: Set the status code to 201 Created: w.WriteHeader(http.StatusCreated)
	//         Return the new student as JSON.

	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func main() {
	http.HandleFunc("/students/", handleGetStudent)
	http.HandleFunc("/students", handleAddStudent)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
