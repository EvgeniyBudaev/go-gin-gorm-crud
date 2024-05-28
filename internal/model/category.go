package model

import "github.com/google/uuid"

type Category struct {
	ID       *uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name     *string    `gorm:"type:varchar(255);not null" json:"name"`
	Products []*Product
}
