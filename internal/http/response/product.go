package response

import "github.com/google/uuid"

type ProductResponse struct {
	ID         *uuid.UUID `json:"id"`
	Name       *string    `json:"name"`
	CategoryID *uuid.UUID `json:"categoryId"`
}
