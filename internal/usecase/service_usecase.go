package usecase

import (
	"errors"
	"strconv"
	"strings"

	"github.com/tersorasora/freelance/internal/entity"
	"github.com/tersorasora/freelance/internal/repository"
)

type ServiceUseCase interface {
	CreateService(serviceName, description string, price float64, period string, fieldID string, userID string) (*entity.Service, error)
	DeleteService(serviceID string) error
	SearchServices(serviceName string, fieldID string) ([]entity.Service, error)
	GetAllServices() ([]entity.Service, error)
	GetMyServices(userID string) ([]entity.Service, error)
}

type serviceUseCase struct {
	serviceRepo     repository.ServiceRepository
	userRepo   		repository.UserRepository
	fieldRepo  		repository.FieldRepository
}

func NewServiceUseCase(sR repository.ServiceRepository, uR repository.UserRepository, fR repository.FieldRepository) ServiceUseCase {
	return &serviceUseCase{sR, uR, fR}
}

func (suc *serviceUseCase) CreateService(serviceName, description string, price float64, period string, fieldID string, userID string) (*entity.Service, error) {
	lastID, err := suc.serviceRepo.GetLastServiceID()
	if err != nil {
		return nil, err
	}

	var newID int
	var newServiceID string
	if lastID == "" {
		newID = 1
	}else{
		parts := strings.Split(lastID, "-")
		if len(parts) == 2 {
			lastNumber, _ := strconv.Atoi(parts[1])
			newID = lastNumber + 1
		}else{
			newID = 1
		}
	}
	newServiceID = "SID-" + strconv.Itoa(newID)

	newService := &entity.Service{
		ServiceID:   newServiceID,
		ServiceName: serviceName,
		Description: description,
		Price:       price,
		Period:      period,
		FieldID:     fieldID,
		UserID:      userID,
	}
	err = suc.serviceRepo.CreateService(newService)
	if err != nil {
		return nil, err
	}
	return newService, nil
}

func (suc *serviceUseCase) DeleteService(serviceID string) error {
	err := suc.serviceRepo.DeleteService(serviceID)
	if err != nil {
		return err
	}
	return nil
}

func (suc *serviceUseCase) SearchServices(serviceName string, fieldID string) ([]entity.Service, error) {
	services, err := suc.serviceRepo.SearchServices(serviceName, fieldID)
	if err != nil {
		return nil, errors.New("field not found")
	}
	return services, nil
}

func (suc *serviceUseCase) GetAllServices() ([]entity.Service, error) {
	services, err := suc.serviceRepo.GetAllServices()
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (suc *serviceUseCase) GetMyServices(userID string) ([]entity.Service, error) {
	myServices, err := suc.serviceRepo.GetMyServices(userID)
	if err != nil {
		return nil, err
	}
	return myServices, nil
}