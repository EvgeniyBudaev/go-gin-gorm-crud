package model

import "github.com/google/uuid"

type Product struct {
	ID         *uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name       *string    `gorm:"type:varchar(255);not null" json:"name"`
	CategoryID *uuid.UUID `gorm:"type:uuid;references:categories.id" json:"categoryId"`
}
