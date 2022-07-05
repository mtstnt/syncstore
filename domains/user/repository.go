package user

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db}
}

func (r Repository) GetByUsername(username string) (*User, error) {
	user := &User{}
	r.db.Find(user, "username = ?", username)

	if r.db.Error != nil {
		return nil, r.db.Error
	}

	return user, nil
}

func (r Repository) CreateUser(username string, hashedPassword string) (*User, error) {
	user := &User{
		Username: username,
		Password: hashedPassword,
	}

	r.db.Create(user)
	if r.db.Error != nil {
		return nil, r.db.Error
	}

	return user, nil
}
