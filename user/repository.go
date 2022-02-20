package user

import "gorm.io/gorm"

type Repository interface {
	FindByEmail(email string) (User, error)
	FindById(ID int) (User, error)
	Save(user User) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

// Function to populate a repository struct
func UserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

//Create a new user
func (r *repository) Save(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

//Find user by email
func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindById(ID int) (User, error) {
	var user User
	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
