package user

import (
	"errors"
	"log/slog"
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/maoudev/todo/internal/config"
	"github.com/maoudev/todo/internal/pkg/domain"
	"github.com/maoudev/todo/internal/pkg/ports"
	"github.com/maoudev/todo/internal/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidUserEmail = errors.New("the user email is invalid")
	ErrPasswordTooShort = errors.New("the password is too short")
)

type userService struct {
	repository ports.UserRepository
}

func NewService(repository ports.UserRepository) *userService {
	return &userService{
		repository: repository,
	}
}

func (u *userService) Create(user *domain.User) error {
	user.ID = utils.CreateID()

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return ErrInvalidUserEmail
	}

	if !isPasswordValid(user.Password) {
		return ErrPasswordTooShort
	}

	user.Password = hashPassword(user.Password)

	return u.repository.Create(user)
}

func (u *userService) Login(credentials *domain.DefaultCredentials) (string, error) {
	_, err := mail.ParseAddress(credentials.Email)
	if err != nil {
		return "", ErrInvalidUserEmail
	}

	user := &domain.User{}

	if err := u.repository.First(user, "email = ?", credentials.Email); err != nil {
		return "", err
	}

	if err := tryMatchPassword(user.Password, credentials.Password); err != nil {
		return "", err
	}

	return createToken(user)
}

func isPasswordValid(password string) bool {
	return len(password) >= 8
}

func createToken(user *domain.User) (string, error) {
	claims := domain.Claims{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return jwtToken.SignedString([]byte(config.SECRET_KEY))
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), config.HASH_COST)
	if err != nil {
		slog.Error("Error when creating hash appears from the password entered")
		return ""
	}

	return string(hash)
}

func tryMatchPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
