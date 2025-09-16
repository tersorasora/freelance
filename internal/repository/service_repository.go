package repository

import (
	"github.com/tersorasora/freelance/internal/entity"

	"gorm.io/gorm"
)

type ServiceRepository interface {
	CreateService(service *entity.Service) error
	DeleteService(serviceID string) error
	GetAllServices() ([]entity.Service, error)
	GetMyServices(userID string) ([]entity.Service, error)
	SearchServices(serviceName string, fieldID string) ([]entity.Service, error)
	GetLastServiceID() (string, error)
}

type serviceRepositoryData struct {
	db *gorm.DB
}
func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepositoryData{db}
}

func (srd *serviceRepositoryData) CreateService(service *entity.Service) error {
	return srd.db.Create(service).Error
}

func (srd *serviceRepositoryData) DeleteService(serviceID string) error {
	return srd.db.Where("service_id = ?", serviceID).Delete(&entity.Service{}).Error
}

func (srd *serviceRepositoryData) GetAllServices() ([]entity.Service, error) {
	var services []entity.Service
	result := srd.db.Find(&services)
	if result.Error != nil {
		return nil, result.Error
	}

	return services, nil
}

func (srd *serviceRepositoryData) GetMyServices(userID string) ([]entity.Service, error) {
	var myServices []entity.Service
	result := srd.db.Where("user_id = ?", userID).Find(&myServices)
	if result.Error != nil {
		return nil, result.Error
	}

	return myServices, nil
}

func (srd *serviceRepositoryData) SearchServices(serviceName string, fieldID string) ([]entity.Service, error) {
	var services []entity.Service
	query := srd.db.Model(&entity.Service{})

	if fieldID != "" {
        query = query.Where("field_id = ?", fieldID)
    }
    if serviceName != "" {
        query = query.Where("service_name ILIKE ?", "%"+serviceName+"%")
    }

    result := query.Find(&services)
    if result.Error != nil {
        return nil, result.Error
    }

    return services, nil
}

func (srd *serviceRepositoryData) GetLastServiceID() (string, error) {
	var lastID string
	err := srd.db.Model(&entity.Service{}).Order("service_id DESC").Limit(1).Pluck("service_id", &lastID).Error
	if err != nil{
		return "", err
	}
	return lastID, nil
}