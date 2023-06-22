package contact

import "github.com/google/uuid"

type ContactRequest struct {
	Id        uuid.UUID
	Name      string `binding:"required"`
	Gender    string `binding:"required"`
	Phone     string `binding:"required"`
	Email     string `binding:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
