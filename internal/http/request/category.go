package request

import "github.com/google/uuid"

type CreateCategoryRequest struct {
	Name *string `validate:"required,min=1,max=255" json:"name"`
}

type UpdateCategoryRequest struct {
	ID   *uuid.UUID `validate:"required"`
	Name *string    `validate:"required,min=1,max=255" json:"name"`
}
