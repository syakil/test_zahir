package contact

import "time"

type ContactData struct {
	Id        string
	Name      string
	Gender    string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
