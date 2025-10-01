package models

import "gorm.io/gorm"

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type Company struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(100);unique;not null" validate:"required,min=2,max=100"`
	Address     string `json:"address" gorm:"type:text"`
	Phone       string `json:"phone" gorm:"type:varchar(20)"`
	Email       string `json:"email" gorm:"type:varchar(100);unique" validate:"email"`
	Website     string `json:"website" gorm:"type:varchar(255)"`
	CompanyType string `json:"company_type" gorm:"type:varchar(50);default:'General'"`
	IsActive    *bool  `json:"is_active" gorm:"default:true"`
}