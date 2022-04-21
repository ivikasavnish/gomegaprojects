package models

import (
	"golang.org/x/crypto/bcrypt"
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