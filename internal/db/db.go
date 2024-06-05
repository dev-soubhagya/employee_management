package db

import (
	"database/sql"
	"log"

	"github.com/dev-soubhagya/employee_management/internal/model"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS employees (
        id INT AUTO_INCREMENT,
        name VARCHAR(50),
        position VARCHAR(50),
        salary FLOAT,
        PRIMARY KEY (id)
    );`
	_, err = Db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	Db.Close()
}

func CreateEmployee(employee model.Employee) error {
	query := `INSERT INTO employees (name, position, salary) VALUES (?, ?, ?)`
	_, err := Db.Exec(query, employee.Name, employee.Position, employee.Salary)
	return err
}

func GetEmployeeByID(id int) (model.Employee, error) {
	var employee model.Employee
	query := `SELECT id, name, position, salary FROM employees WHERE id = ?`
	err := Db.QueryRow(query, id).Scan(&employee.ID, &employee.Name, &employee.Position, &employee.Salary)
	return employee, err
}

func UpdateEmployee(employee model.Employee) error {
	query := `UPDATE employees SET name = ?, position = ?, salary = ? WHERE id = ?`
	_, err := Db.Exec(query, employee.Name, employee.Position, employee.Salary, employee.ID)
	return err
}

func DeleteEmployee(id int) error {
	query := `DELETE FROM employees WHERE id = ?`
	_, err := Db.Exec(query, id)
	return err
}

func ListEmployees(page, perPage int) ([]model.Employee, error) {
	var employees []model.Employee
	query := `SELECT id, name, position, salary FROM employees LIMIT ? OFFSET ?`
	rows, err := Db.Query(query, perPage, (page-1)*perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee model.Employee
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Position, &employee.Salary); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}
