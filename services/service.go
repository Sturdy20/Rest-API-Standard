package services

import (
	"company/models"
	"company/repositories"
	"log"
)

type IService interface {
	AddRegisterService(register models.RequestRegister) error
	GetAllEmployeesService() ([]models.ResponseEmployee, error)
	GetCompanyByIDService(cpnID string) (models.ResponseCompany, error)
	UpdateEmployeeService(employee models.RequestUpdateEmployee) error
	DeleteEmployeeService(epyID string) error
	
}

type service struct {
	r repositories.IRepositorie
}

func NewService(r repositories.IRepositorie) IService {
	return &service{r: r}
}

func (s *service) AddRegisterService(register models.RequestRegister) error {
	err := s.r.AddRegisterRepositorie(register)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (s *service) GetAllEmployeesService() ([]models.ResponseEmployee, error) {
	employees, err := s.r.GetAllEmployeesRepositorie()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return employees, nil
}

func (s *service) GetCompanyByIDService(cpnID string) (models.ResponseCompany, error) {
	company, err := s.r.GetCompanyByIDRepositorie(cpnID)
	if err != nil {
		log.Println(err.Error())
		return models.ResponseCompany{}, err
	}

	return company, nil
}

func (s *service) UpdateEmployeeService(employee models.RequestUpdateEmployee) error {
	err := s.r.UpdateEmployeeRepositorie(employee)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
func (s *service) DeleteEmployeeService(epyID string) error {
	err := s.r.DeleteEmployeeRepositorie(epyID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
