package response

import "github.com/google/uuid"

type CategoryResponse struct {
	ID   *uuid.UUID `json:"id,omitempty"`
	Name *string    `json:"name,omitempty"`
}
