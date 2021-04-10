package user

import (
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/sobocinski/go-api-test/internal/car"
)

type UserId uint

type User struct {
	Id        uint      `json:"id,omitempty" sql:"id,pk"`
	Password  string    `json:"-" sql:"password"`
	Email     string    `json:"email" sql:"email" pg:"email,unique"`
	Cars      []car.Car `json:"cars" pg:"rel:has-many"`
	CreatedAt time.Time `json:"createdAt" sql:"created_at"`
	UpdatedAt time.Time `json:"-" sql:"updated_at"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(id UserId) (User, error)
	Create(user User) (*User, error)
}

func UserIdFromString(id string) UserId {
	userId64, _ := strconv.ParseUint(id, 10, 64)
	return UserId(uint(userId64))
}

func hashAndSaltPwd(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwd), err
}
