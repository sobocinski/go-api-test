package car

import (
	"time"

	"github.com/sobocinski/go-api-test/internal/user"
)

type CarId uint

type Car struct {
	Id                 CarId
	Name               string
	Manufacturer       string
	Model              string
	RegistrationNumber string
	CreatedAt          time.Time
	UserId             user.UserId
}

type CarRepository interface {
	GetAll() ([]Car, error)
	GetById(id CarId) (Car, error)
}
