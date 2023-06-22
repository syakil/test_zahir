package contact

import "github.com/google/uuid"

type Service interface {
	FindAll() ([]ContactData, error)
	FindById(ID string) (ContactData, error)
	Create(contact ContactRequest) (ContactData, error)
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]ContactData, error) {
	contacts, err := s.repository.FindAll()
	return contacts, err
}

func (s *service) FindById(ID string) (ContactData, error) {
	contact, err := s.repository.FindById(ID)
	return contact, err
}

func (s *service) Create(request ContactRequest) (ContactData, error) {
	contact := ContactData{
		Id:     uuid.NewString(),
		Name:   request.Name,
		Gender: request.Gender,
		Phone:  request.Phone,
		Email:  request.Email,
	}
	newContact, err := s.repository.Create(contact)
	return newContact, err
}
