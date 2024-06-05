package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/dev-soubhagya/employee_management/internal/db"
	"github.com/dev-soubhagya/employee_management/internal/handler"
	"github.com/dev-soubhagya/employee_management/internal/model"
	"github.com/gorilla/mux"
)

func setupTestRouter() *mux.Router {
	return handler.SetupRouter()
}

func TestCreateEmployeeHandler(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	router := setupTestRouter()
	employee := model.Employee{
		Name:     "John Doe",
		Position: "Software Engineer",
		Salary:   50000,
	}

	employeeJSON, _ := json.Marshal(employee)
	req, err := http.NewRequest("POST", "/employees", bytes.NewBuffer(employeeJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestGetEmployeeHandler(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	router := setupTestRouter()

	employee := model.Employee{
		Name:     "John Doe",
		Position: "Software Engineer",
		Salary:   50000,
	}

	err := db.CreateEmployee(employee)
	if err != nil {
		t.Fatalf("failed to create employee: %v", err)
	}

	var empID int
	err = db.Db.QueryRow("SELECT id FROM employees WHERE name = ?", employee.Name).Scan(&empID)
	if err != nil {
		t.Fatalf("failed to get employee ID: %v", err)
	}

	req, err := http.NewRequest("GET", "/employees/"+strconv.Itoa(empID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestUpdateEmployeeHandler(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	router := setupTestRouter()

	employee := model.Employee{
		Name:     "John Doe",
		Position: "Software Engineer",
		Salary:   50000,
	}

	err := db.CreateEmployee(employee)
	if err != nil {
		t.Fatalf("failed to create employee: %v", err)
	}

	var empID int
	err = db.Db.QueryRow("SELECT id FROM employees WHERE name = ?", employee.Name).Scan(&empID)
	if err != nil {
		t.Fatalf("failed to get employee ID: %v", err)
	}

	updatedEmployee := model.Employee{
		ID:       empID,
		Name:     "Jane Doe",
		Position: "Senior Software Engineer",
		Salary:   60000,
	}

	updatedEmployeeJSON, _ := json.Marshal(updatedEmployee)
	req, err := http.NewRequest("PUT", "/employees/"+strconv.Itoa(empID), bytes.NewBuffer(updatedEmployeeJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestListEmployeesHandler(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	router := setupTestRouter()

	req, err := http.NewRequest("GET", "/employees?page=1&per_page=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteEmployeeHandler(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	router := setupTestRouter()

	employee := model.Employee{
		Name:     "John Doe",
		Position: "Software Engineer",
		Salary:   50000,
	}

	err := db.CreateEmployee(employee)
	if err != nil {
		t.Fatalf("failed to create employee: %v", err)
	}

	var empID int
	err = db.Db.QueryRow("SELECT id FROM employees WHERE name = ?", employee.Name).Scan(&empID)
	if err != nil {
		t.Fatalf("failed to get employee ID: %v", err)
	}

	req, err := http.NewRequest("DELETE", "/employees/"+strconv.Itoa(empID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
