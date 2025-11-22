package router

import (
	teachers "apiproject02/internal/modules/teachers"
	"net/http"
)

func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/teachers/", teachers.TeacherHandler)
	return mux
}
