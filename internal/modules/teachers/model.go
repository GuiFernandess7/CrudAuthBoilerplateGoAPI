package teachers

import (
	"gorm.io/gorm"
)

type TeacherHandlers struct {
    DB *gorm.DB
}

type Teacher struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Class     string `json:"class" validate:"required"`
	Subject   string `json:"subject" validate:"required"`
}
