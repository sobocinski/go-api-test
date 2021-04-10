package user

import (
	"github.com/go-pg/pg/v10"
)

type DbUserRepository struct {
	db *pg.DB
}

// UserRepositoryDb - returns DB(sql) UserRepository
func UserRepositoryDb(db *pg.DB) *DbUserRepository {
	return &DbUserRepository{
		db: db,
	}
}

func (r *DbUserRepository) GetById(id UserId) (User, error) {
	user := User{}
	err := r.db.Model(&user).Where("? = ?", pg.Ident("id"), id).Select()
	return user, err
}

func (r *DbUserRepository) GetAll() ([]User, error) {
	users := []User{}
	err := r.db.Model(&users).Select()
	return users, err
}

func (r *DbUserRepository) Create(user User) (*User, error) {
	_, err := r.db.Model(&user).Returning("id").Returning("created_at").Insert()
	return &user, err
}
