package request

import "github.com/google/uuid"

type CreateProductRequest struct {
	Name       *string    `validate:"required,min=1,max=255" json:"name"`
	CategoryID *uuid.UUID `validate:"required" json:"categoryId"`
}

type UpdateProductRequest struct {
	ID         *uuid.UUID `validate:"required"`
	Name       *string    `validate:"required,min=1,max=255" json:"name"`
	CategoryID uuid.UUID  `validate:"required" json:"categoryId"`
}
