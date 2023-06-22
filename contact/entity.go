package contact

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type ContactData struct {
	Id        uuid.UUID `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string
	Gender    string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
