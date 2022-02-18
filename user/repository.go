package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

// Function to populate the repository struct
func UserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

//Save user data
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

//Find user by email
func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
