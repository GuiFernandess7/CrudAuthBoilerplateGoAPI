package teachers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	modules "apiproject02/internal/modules"
	sqlconnect "apiproject02/internal/repository/sqlconnect"
	"gorm.io/driver/mysql"
)

var (
	teachers 	= make(map[int]Teacher)
	mutex 		= &sync.Mutex{}
	nextID 		= 1
)

func init() {
	teachers[nextID] = Teacher {
		ID: nextID,
		FirstName: "John",
		LastName: "Cena",
		Class: "A",
		Subject: "Fight",
	}
	nextID++
	teachers[nextID] = Teacher {
		ID: nextID,
		FirstName: "Jake",
		LastName: "Peralta",
		Class: "B",
		Subject: "Investigation",
	}
	nextID++
	teachers[nextID] = Teacher {
		ID: nextID,
		FirstName: "Robert",
		LastName: "Greene",
		Class: "B",
		Subject: "Biology",
	}
	nextID++
}


func TeacherHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTeachersHandlers(w, r)
	case http.MethodPost:
		createTeacherHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTeachersHandlers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusForbidden)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	fmt.Println(idStr)

	if idStr == "" {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teacherList := make([]Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) &&
				(lastName == "" || teacher.LastName == lastName) {
				teacherList = append(teacherList, teacher)
			}
		}

		response := struct {
			Status string    `json:"status"`
			Count  int       `json:"count"`
			Data   []Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teacherList),
			Data:   teacherList,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return

	} else {
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid parameter", http.StatusForbidden)
			return
		}

		teacher, exists := teachers[idInt]
		if !exists {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(teacher)
		return
	}
}

func createTeacherHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
	    Conn: db,
	}), &gorm.Config{})

	if err != nil {
	    panic("failed to initialize gorm DB")
	}

	defer db.Close()
	var newTeachers []Teacher
	err = json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	crud := modules.CrudMaker[Teacher]{DB: db}
	for _, teacher := range newTeachers {
		err := crud.Create(teacher)
		if err != nil {
			http.Error(w, "Error inserting data into database", http.StatusBadRequest)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
