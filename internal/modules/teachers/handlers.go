package teachers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	modules "apiproject02/internal/modules"
	"gorm.io/gorm"
	"errors"
)

func (h *TeacherHandlers) TeachersHandler(w http.ResponseWriter, r *http.Request) {

    switch r.Method {
    case http.MethodGet:
        h.GetTeachers(w, r)
    case http.MethodPost:
        h.AddTeacherHandler(w, r)
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

func (h *TeacherHandlers) GetTeachers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusForbidden)
		return
	}
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")

	if idStr == "" {
		crud := modules.CrudGeneric[Teacher]{DB: h.DB}
		responseModel, err := crud.ReadAll()

		if err != nil {
			fmt.Println("Parsing model error:", err)
			http.Error(w, "", http.StatusInternalServerError)
   			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseModel)

	} else {
		idInt, err := strconv.Atoi(idStr)

		if err != nil {
		    fmt.Println("Error converting string ID to integer:", err)
		    return
		}
		crud := modules.CrudGeneric[Teacher]{DB: h.DB}
		responseModel, err := crud.Read(uint(idInt))

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
		        http.Error(w, "Record not found", http.StatusNotFound)
		        return
		    }

		    fmt.Println("Database error:", err)
		    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		    return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseModel)
	}
}

func (h *TeacherHandlers) AddTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var responseModel []Teacher
	err := json.NewDecoder(r.Body).Decode(&responseModel)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	crud := modules.CrudGeneric[Teacher]{DB: h.DB}
	for i := range responseModel {
		err := crud.Create(&responseModel[i])
		if err != nil {
			http.Error(w, "Error inserting data into database", http.StatusBadRequest)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseModel)
}
