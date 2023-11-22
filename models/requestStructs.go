package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Users struct {
	Id          uint      `json:"user_id" orm:"pk;auto"`
	FirstName   string    `json:"first_name" orm:"null"`
	LastName    string    `json:"last_name" orm:"null"`
	Email       string    `json:"email" orm:"unique"`
	PhoneNumber string    `json:"phone_number" orm:"null"`
	Country     string    `json:"country"`
	Role        string    `json:"role"`
	Age         int       `json:"age" orm:"size(3)"`
	Password    string    `json:"password" orm:"notnull"`
	CreatedAt   time.Time `orm:"null"`
	UpdatedAt   time.Time `orm:"null"`
	DeletedAt   time.Time `orm:"null"`
}

type Book struct {
	Id          int    `orm:"auto;pk"`
	Name        string `orm:"size(100)"`
	Author      string
	Description string
	CreatedAt   time.Time `orm:"null"`
	UpdatedAt   time.Time `orm:"null"`
}

type Students struct {
	Id         uint      `json:"student_id" orm:"pk;auto"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email" orm:"unique;notnull"`
	Country    string    `json:"country" orm:"notnull"`
	RollNumber string    `json:"roll_number" orm:"notnull;unique"`
	Age        int       `json:"age" orm:"notnull;size(3)"`
	ClassId    uint      `json:"class_id"`
	Password   string    `json:"password" orm:"notnull"`
	CreatedAt  time.Time `orm:"null"`
	UpdatedAt  time.Time `orm:"null"`
	DeletedAt  time.Time `orm:"null"`
}

type InsertBook struct {
	Name   string `json:"book_name"`
	Author string `json:"author_name"`
}

type GetBooks struct {
	Id int `json:"book_id"`
}

type UpdateBooks struct {
	Id     int    `json:"book_id"`
	Name   string `json:"book_name"`
	Author string `json:"author_name"`
}

type ClassWiseStudent struct {
	ClassId int `json:"class_id"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewUserRequest struct {
	FirstName string `json:"first_name" orm:"null"`
	LastName  string `json:"last_name" orm:"null"`
	Email     string `json:"email" orm:"unique"`
	Country   string `json:"country"`
	Role      string `json:"role"`
	Age       int    `json:"age" orm:"size(3)"`
	Password  string `json:"password" orm:"notnull"`
}
type UpdateUserRequest struct {
	Id        uint   `json:"user_id"`
	FirstName string `json:"first_name" orm:"null"`
	LastName  string `json:"last_name" orm:"null"`
	Email     string `json:"email" orm:"unique"`
	Country   string `json:"country"`
	Role      string `json:"role"`
	Age       int    `json:"age" orm:"size(3)"`
}

type ResetUserPassword struct {
	CurrentPass string `json:"current_password"`
	NewPass     string `json:"new_password"`
	ConfirmPass string `json:"confirm_password"`
}

type JwtClaim struct {
	Email string
	ID    int
	jwt.StandardClaims
}

type UserDetailsRequest struct {
	Id        uint   `json:"user_id"`
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name" `
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Country   string `json:"country"`
}

type SendOtpData struct {
	Email string `json:"email"`
}

type ResetUserPasswordOtp struct {
	Email   string `json:"email"`
	Otp     string `json:"otp"`
	NewPass string `json:"new_password"`
}
