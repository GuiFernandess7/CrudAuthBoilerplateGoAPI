package modules

import (
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
	"fmt"
)

type CrudGeneric[T any] struct {
	DB *gorm.DB
}

func ValidateModel[T any](payload *T) error{
	return validate.Struct(payload)
}

func (c *CrudGeneric[T]) Create(item *T) error {
	if err := c.ValidateModel(item) != nil {
		fmt.Println("Invalid payload: ", err)
		return err
	}
	return c.DB.Create(item).Error
}

func (c *CrudGeneric[T]) Read(id any) (*T, error) {
	var model T

	if err := c.ValidateModel(item) != nil {
		fmt.Println("Invalid payload: ", err)
		return nil, err
	}

	if err := c.DB.First(&model, id).Error; err != nil {
		fmt.Println("Error reading object from database: ", err)
		return nil, err
	}
	return &model, nil
}

// func (c *CrudGeneric[T]) ReadAll() ([]*T, error) {
// 	var objects []T
// 	if err := c.DB.
// }

func (c *CrudGeneric[T]) Update(id any, updated *T) error {
	return c.DB.Model(new(T)).Where("id = ?", id).Updates(updated).Error
}

func (c *CrudGeneric[T]) Delete(id any) error {
	return c.DB.Delete(new(T), id).Error
}
