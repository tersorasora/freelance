package usecase

import (
	"errors"
	"strconv"
	"strings"

	"github.com/tersorasora/freelance/internal/entity"
	"github.com/tersorasora/freelance/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	RegisterUser(email string, name string, password string) (*entity.User, error)
	LoginUser(email, password string) (*entity.User, error)
	GetUser(userID string) (*entity.User, error)
	DeleteUser(userID string) error
	GetTotalUsers() (int64, error)
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

	lastID, err := uuc.repo.GetLastUserID()
	if err != nil {
		return nil, err
	}

	var newID int
	var newUserID string
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
	newUserID = "UID-" + strconv.Itoa(newID)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

    user := &entity.User{
        UserID: newUserID,
        Email:  email,
		Name:   name,
		Password: string(hashedPassword),
		Balance: 0.00,
		RoleID: "RL-2", // Default role as regular user
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
		return nil, errors.New("email belum terdaftar")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, errors.New("email atau password salah")
    }

	return user, nil
}

func (uuc *userUsecase) GetUser(userID string) (*entity.User, error) {
	user, err := uuc.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uuc *userUsecase) DeleteUser(userID string) error {
	err := uuc.repo.DeleteUser(userID)
	if err != nil {
		return err
	}

	return nil
}

func (uuc *userUsecase) GetTotalUsers() (int64, error) {
	total, err := uuc.repo.GetTotalUsers()
	if err != nil {
		return 0, err
	}
	return total, nil
}