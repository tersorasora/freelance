package usecase

import (
	"errors"
	"strconv"
	"strings"
	"fmt"

	"github.com/tersorasora/freelance/internal/entity"
	"github.com/tersorasora/freelance/internal/repository"
)

type FieldUseCase interface {
	CreateField(name string) (*entity.Field, error)
	DeleteField(fieldID string) error
	GetAllFields() ([]entity.Field, error)
	GetFieldByID(fieldID string) (*entity.Field, error)
}

type fieldUseCase struct {
	repo repository.FieldRepository
}

func NewFieldUseCase(r repository.FieldRepository) FieldUseCase {
	return &fieldUseCase{r}
}

func (fuc *fieldUseCase) CreateField(name string) (*entity.Field, error) {
	fmt.Println("lewat3")
	lastID, err := fuc.repo.GetLastFieldID()
	if err != nil {
		return nil, err
	}

	if lastID != "" {
		nameFound, err := fuc.repo.GetFieldByName(name)
		if err == nil && nameFound != nil {
			return nil, errors.New("field sudah terdaftar")
		}
	}

	var newID int
	var newFieldID string
	if lastID == "" {
		newID = 1
	}else{
		parts := strings.Split(lastID, "-")
		fmt.Println("lewat4")
		if len(parts) == 2 {
			lastNumber, _ := strconv.Atoi(parts[1])
			newID = lastNumber + 1
		}else{
			newID = 1
		}
	}
	newFieldID = "FID-" + strconv.Itoa(newID)

	field := &entity.Field{
		FieldID:   newFieldID,
		FieldName: name,
	}
	
	err = fuc.repo.CreateField(field)
	if err != nil {
		return nil, err
	}

	return field, nil
}

func (fuc *fieldUseCase) DeleteField(fieldID string) error {
	err := fuc.repo.DeleteField(fieldID)
	if err != nil {
		return err
	}

	return nil
}

func (fuc *fieldUseCase) GetAllFields() ([]entity.Field, error) {
	fields, err := fuc.repo.GetAllFields()
	if err != nil {
		return nil, err
	}
	return fields, nil
}

func (fuc *fieldUseCase) GetFieldByID(fieldID string) (*entity.Field, error) {
	field, err := fuc.repo.GetFieldByID(fieldID)

	if err != nil {
		return nil, err
	}

	return field, nil
}