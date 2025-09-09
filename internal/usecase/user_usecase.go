package usecase

import (
	"errors"

	"github.com/tersorasora/freelance/internal/entity"
	"github.com/tersorasora/freelance/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	RegisterUser(email string, name string, password string) (*entity.User, error)
	LoginUser(email, password string) (*entity.User, error)
}

type userUsecase struct {
    repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
    return &userUsecase{r}
}

func (uuc *userUsecase) RegisterUser(email string, name string, password string) (*entity.User, error) {
	// Check if user already exists
    existing, _ := uuc.repo.GetUserByEmail(email)
    if existing != nil {
        return nil, errors.New("email sudah terdaftar")
    }

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

    user := &entity.User{
        UserID: uuid.NewString(),
        Email:  email,
		Name:   name,
		Password: string(hashedPassword),
    }

    err = uuc.repo.CreateUser(user)
    if err != nil {
        return nil, err
    }

    return user, nil
}

func (uuc *userUsecase) LoginUser(email, password string) (*entity.User, error) {
	user, err := uuc.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, errors.New("email atau password salah")
    }

	return user, nil
}