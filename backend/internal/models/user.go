package models

import (
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    gorm.Model
    Email     string `json:"email" gorm:"uniqueIndex;not null"`
    Password  string `json:"-" gorm:"not null"` // "-" prevents password from being included in JSON
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
    Role      string `json:"role" gorm:"default:'user'"`
    Tasks     []Task `json:"tasks,omitempty" gorm:"foreignKey:UserID"`
}

type UserCreate struct {
    Email     string `json:"email" binding:"required,email"`
    Password  string `json:"password" binding:"required,min=8"`
    FirstName string `json:"firstName" binding:"required"`
    LastName  string `json:"lastName" binding:"required"`
}

type UserLogin struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type UserUpdate struct {
    FirstName *string `json:"firstName"`
    LastName  *string `json:"lastName"`
    Password  *string `json:"password"`
}

func (u *User) HashPassword() error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    return nil
}

func (u *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}