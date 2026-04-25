package networking

import (
	"context"
	"docker-study/employees"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type EmployeesHandlers struct {
	//employees map[int]employees.EmployeeDto
	ctx    context.Context
	dbConn *pgx.Conn
}

func NewEmployeesHandlers(ctx context.Context, dbConn *pgx.Conn) *EmployeesHandlers {
	return &EmployeesHandlers{
		ctx:    ctx,
		dbConn: dbConn,
	}
}

func (h *EmployeesHandlers) GetEmployeesList(w http.ResponseWriter, r *http.Request) {
	sql := `
		SELECT * FROM employees;
	`
	rows, err := h.dbConn.Query(h.ctx, sql)
	if err != nil {
		panic(err)
	}

	var (
		emp       employees.Employee
		employees []employees.Employee
	)
	for rows.Next() {
		rows.Scan(
			&emp.ID,
			&emp.FullName,
			&emp.Position,
		)

		employees = append(employees, emp)
	}

	b, err := json.MarshalIndent(employees, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		panic(err)
	}
}

func (h *EmployeesHandlers) AddEmployee(w http.ResponseWriter, r *http.Request) {
	var emp EmployeeDto
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
		return
	}

	//h.employees[rand.Int()] = emp
	sql := `
		INSERT INTO employees (full_name, position)
		VALUES ($1, $2);
	`
	if _, err := h.dbConn.Exec(h.ctx, sql, emp.FullName, emp.Position); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(emp); err != nil {
		panic(err)
	}
}

func (h *EmployeesHandlers) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
		return
	}

	id, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
		return
	}

	//delete(h.employees, int(id))
	sql := `
		DELETE FROM employees
		WHERE id=$1;
	`
	if _, err := h.dbConn.Exec(h.ctx, sql, id); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}
