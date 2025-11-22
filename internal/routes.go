package router

import (
	teachers "apiproject02/internal/modules/teachers"
	"net/http"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *http.ServeMux {
	handlers := &teachers.TeacherHandlers{DB: db}

	mux := http.NewServeMux()
	mux.HandleFunc("/teachers/", handlers.GetTeachers)
	return mux
}
