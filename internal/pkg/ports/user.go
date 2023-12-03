package ports

import "github.com/maoudev/todo/internal/pkg/domain"

type UserRepository interface {
	Create(value interface{}) error
	First(out interface{}, conds ...interface{}) error
}

type UserService interface {
	Create(user *domain.User) error
	Login(credentials *domain.DefaultCredentials) (string, error)
	GetUser(userID string) (*domain.User, error)
}
