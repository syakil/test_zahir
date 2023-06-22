package contact

import uuid "github.com/satori/go.uuid"

type Service interface {
	FindAll() ([]ContactData, error)
	FindById(ID string) (ContactData, error)
	Create(contact ContactRequest) (ContactData, error)
	Delete(ID string) (ContactData, error)
	Update(ID string, contact ContactUpdate) (ContactData, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
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
		Id:     uuid.NewV4(),
		Name:   request.Name,
		Gender: request.Gender,
		Phone:  request.Phone,
		Email:  request.Email,
	}
	newContact, err := s.repository.Create(contact)
	return newContact, err
}

func (s *service) Delete(ID string) (ContactData, error) {
	contact, err := s.repository.Delete(ID)
	return contact, err
}

func (s *service) Update(ID string, request ContactUpdate) (ContactData, error) {
	contact, err := s.repository.FindById(ID)
	contact.Name = request.Name
	contact.Gender = request.Gender
	contact.Phone = request.Phone
	contact.Email = request.Email

	newContact, err := s.repository.Update(contact)
	return newContact, err
}
