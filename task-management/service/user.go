package service

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"

	"task/database"
	"task/model"
)

var allUser []*model.User

func NewUser(id uint, name string, email string, password string) (*model.User, error) {
	//adding validation checks
	if id < 0 {
		return nil, errors.New("The id of the user cant be less than 0")
	}

	if name == "" {
		return nil, errors.New("The name of the user cant be empty")
	}

	if email == "" {
		return nil, errors.New("The email of the user cant be empty")
	}

	if password == "" {
		return nil, errors.New("The password of the user cant be empty")
	}

	createdAt := time.Now().Truncate(24 * time.Hour)
	updatedAt := createdAt

	var tempUser = &model.User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	db := database.Get_GormDB()
	log.Println(db)
	if err := db.Create(tempUser).Error; err != nil {
		return nil, err
	}
	allUser = append(allUser, tempUser)
	return tempUser, nil

}

func GetUser() ([]*model.User, error) {
	var user []*model.User

	db := database.Get_GormDB()
	log.Println("fetcheing data from the database")

	if err := db.Find(&user).Error; err != nil {
		log.Println("Error fetching users : %v", err)
		return nil, err
	}

	log.Printf("Users feteched : %d users", len(user))
	return user, nil
}

func GetUserById(userID uint) (*model.User, error) {
	var user model.User

	db := database.Get_GormDB() // database connection established

	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		log.Printf("Error in GetUserById: %v", err)
		return nil, err
	}

	log.Printf("Users feteched : %v", user)
	return &user, nil
}

func UpdateUserName(u *model.User, name string) (*model.User, error) {
	if u.Name == "" {
		return nil, errors.New("the name cant be empty.")
	}

	if name == "" {
		return nil, errors.New("The name cant be empty")
	}

	u.Name = name
	return &model.User{}, nil
}

func UpdateUserID(u *model.User, id uint) error {
	if u.ID < 0 {
		return errors.New("the id cant be null/less than zero.")
	}

	if id < 0 {
		return errors.New("the id cant be null.")
	}

	u.ID = id
	return nil
}

func DeleteUser(u *model.User) (*model.User, error) {
	if u.Name == "" {
		return nil, errors.New("the name cant be null.")
	}

	u.Name = ""

	return &model.User{}, nil
}

func DeleteUserID(userID uint) error {
	tx := database.Get_GormDB().Begin() // Start a transaction
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			log.Println("Error during user deletion:", err)
		} else if err := tx.Commit(); err != nil {
			log.Println("Error committing transaction:", err)
		}
	}()

	var user model.User
	if err := tx.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	if err := tx.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
