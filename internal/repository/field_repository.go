package repository

import (
	"github.com/tersorasora/freelance/internal/entity"

	"gorm.io/gorm"
)

type FieldRepository interface {
	CreateField(field *entity.Field) error
	DeleteField(fieldID string) error
	GetAllFields() ([]entity.Field, error)
	GetFieldByName(fieldName string) (*entity.Field, error)
	GetFieldByID(fieldID string) (*entity.Field, error)
	GetLastFieldID() (string, error)
}

type fieldRepositoryData struct {
	db *gorm.DB
}

func NewFieldRepository(db *gorm.DB) FieldRepository {
	return  &fieldRepositoryData{db}
}

func (frd *fieldRepositoryData) CreateField(field *entity.Field) error {
	return frd.db.Create(field).Error
}

func (frd *fieldRepositoryData) DeleteField(fieldID string) error {
	return frd.db.Where("field_id = ?", fieldID).Delete(&entity.Field{}).Error
}

func (frd *fieldRepositoryData) GetAllFields() ([]entity.Field, error) {
	var fields []entity.Field
	result := frd.db.Find(&fields)
	if(result.Error != nil){
		return nil, result.Error
	}
	
	return fields, nil
}

func (frd *fieldRepositoryData) GetFieldByName(fieldName string) (*entity.Field, error) {
	var field entity.Field
	result := frd.db.Where(("field_name = ?"), fieldName).First(&field)
	if(result.Error != nil){
		return nil, result.Error
	}

	return &field, nil
}

func (frd *fieldRepositoryData) GetFieldByID(fieldID string) (*entity.Field, error) {
	var field entity.Field
	result := frd.db.Where(("field_id = ?"), fieldID).First(&field)
	if(result.Error != nil){
		return nil, result.Error
	}

	return &field, nil
}

func (frd *fieldRepositoryData) GetLastFieldID() (string, error) {
	var lastField string
	err := frd.db.Model(&entity.Field{}).Order("field_id DESC").Limit(1).Pluck("field_id", &lastField).Error
	if err != nil{
		return "", err
	}
	return lastField, nil
}