package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID
	FirstName      *string
	LastName       *string
	Password       *string
	Email          *string
	Phone          *string
	Token          *string
	RefreshToken   *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UserID         string
	UserCart       []ProductUser
	AddressDetails []Address
	OrderStatus    []Order
}

type Product struct{}

type ProductUser struct{}

type Address struct{}

type Order struct{}

type Payment struct{}
