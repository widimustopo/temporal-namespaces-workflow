package entities

import (
	"github.com/gofrs/uuid"
)

type Product struct {
	ProductID    uuid.UUID `json:"product_id" gorm:"primary_key; unique;type:uuid; column:id_request;default:uuid_generate_v4()"`
	ProductName  string    `json:"product_name" validate:"required"`
	Price        float64   `json:"price" validate:"required"`
	CounterOrder int32     `json:"counter_order"`
}

type Member struct {
	MemberID   uuid.UUID `json:"member_id" gorm:"primary_key; unique;type:uuid; column:id_request;default:uuid_generate_v4()"`
	MemberName string    `json:"member_name" validate:"required"`
}

type Payment struct {
	PaymentID     uuid.UUID `json:"payment_id" gorm:"primary_key; unique;type:uuid; column:id_request;default:uuid_generate_v4()"`
	ProductID     string    `json:"product_id" validate:"required"`
	ProductName   string    `json:"product_name" validate:"required"`
	MemberID      string    `json:"member_id" validate:"required"`
	MemberName    string    `json:"member_name" validate:"required"`
	Price         float64   `json:"price" validate:"required"`
	FullPrice     float64   `json:"full_price" validate:"required"`
	StatusPayment string    `json:"status_payment"`
}
