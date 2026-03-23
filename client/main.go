package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Go does not allow unused imports. These lines reference each imported package so
// the compiler accepts them before you use them in your implementation.
// Remove each entry once you actually call functions from that package.
var (
	_ = bytes.NewBuffer
	_ = json.Marshal
	_ = io.ReadAll
	_ = http.Get
)

// Student must match the server's data structure exactly.
type Student struct {
	Name  string `json:"name"`
	Major string `json:"major"`
	Year  int    `json:"year"`
}

// getStudent fetches a student by name from the server.
func getStudent(name string) (Student, error) {
	// TODO 1: Send a GET request to http://localhost:8080/student?name=<name>
	//         http.Get(url) returns (resp, err).

	// TODO 2: Check the error and handle it.

	// TODO 3: Close resp.Body at the end of the function using defer.
	//         defer resp.Body.Close()

	// TODO 4: Check the HTTP status code.
	//         If not http.StatusOK (200): return an error using fmt.Errorf.
	//         You can read the body with io.ReadAll(resp.Body) for the error message.

	// TODO 5: Decode the JSON response body into a Student struct and return it.
	//         var student Student
	//         json.NewDecoder(resp.Body).Decode(&student)

	return Student{}, fmt.Errorf("not implemented")
}

// addStudent sends a new student to the server.
func addStudent(s Student) error {
	// TODO 1: Convert the student struct to JSON bytes.
	//         json.Marshal(s) returns ([]byte, error).

	// TODO 2: Send a POST request to http://localhost:8080/student/add
	//         http.Post(url, "application/json", bytes.NewBuffer(data))

	// TODO 3: Check the error and handle it.

	// TODO 4: Close resp.Body using defer.

	// TODO 5: Check that the status code is http.StatusCreated (201).
	//         If not: return an error.

	return fmt.Errorf("not implemented")
}

func main() {
	// 1. Fetch an existing student
	student, err := getStudent("alice")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Fetched:     Name=%-10s  Major=%-20s  Year=%d\n", student.Name, student.Major, student.Year)

	// 2. Add a new student
	newStudent := Student{
		Name:  "charlie",
		Major: "Physics",
		Year:  1,
	}
	if err := addStudent(newStudent); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 3. Fetch the newly added student
	student, err = getStudent("charlie")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("New student: Name=%-10s  Major=%-20s  Year=%d\n", student.Name, student.Major, student.Year)
}
