package teachers

import (
	"net/http"
)

func router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/teachers/", TeacherHandler)
	return mux
}
