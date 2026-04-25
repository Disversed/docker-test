package networking

import (
	"net/http"

	"github.com/gorilla/mux"
)

type EmployeesServer struct {
	handlers *EmployeesHandlers
	server   *http.Server
}

func NewEmployeesServer(handlers *EmployeesHandlers) *EmployeesServer {
	return &EmployeesServer{handlers: handlers, server: &http.Server{}}
}

func (s *EmployeesServer) StartServer(addr string) error {
	router := mux.NewRouter()

	router.HandleFunc("/employees", s.handlers.GetEmployeesList).Methods("GET")
	router.HandleFunc("/employees", s.handlers.AddEmployee).Methods("POST")
	router.HandleFunc("/employees", s.handlers.DeleteEmployee).Methods("DELETE")

	s.server.Addr = addr
	s.server.Handler = router

	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
