package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

type User struct {
	ID        uint32    `json:"id" gorm:"primary_key;auto_increment"`
	Nickname  string    `json:"nickname" gorm:"size:255;not null;unique"`
	Email     string    `json:"email" gorm:"size:255;not null;unique"`
	Password  string    `json:"password" gorm:"size:255;not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP" `
}

func Hash(Password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
}
func VerifyPassWord(Password, HashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(Password), []byte(HashedPassword))
}
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil

}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		return nil
	case "login":
		return nil
	default:
		return nil

	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, nil
}
func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	user := User{}
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"nickname":  u.Nickname,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
	if db.Error != nil {
		return -1, db.Error
	}
	return db.RowsAffected, nil
}
