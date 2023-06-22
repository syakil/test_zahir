package contact

type ContactRequest struct {
	Name   string `binding:"required"`
	Gender string `binding:"required"`
	Phone  string `binding:"required"`
	Email  string `binding:"required"`
}

type ContactUpdate struct {
	Name   string
	Gender string
	Phone  string
	Email  string
}
