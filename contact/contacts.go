package contact

type Contact struct {
	Id        string
	Name      string `binding:"required"`
	Gender    string `binding:"required"`
	Phone     string `binding:"required"`
	Email     string `binding:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
