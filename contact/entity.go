package contact

import "time"

type ContactData struct {
	Id        string
	Name      string
	Gender    string
	Phone     int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
