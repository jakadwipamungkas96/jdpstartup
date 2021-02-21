package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *repo) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email  = ? ", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
