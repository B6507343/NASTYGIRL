package entity

import (
	// "time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Point    int
	Status   string
	MemberID uint

	Member Member `gorm:"foreignKey:MemberID"`

	// //FK
	// ShowTimesID *uint
	// ShowTime    ShowTimes `gorm:"foreignKey:ShowTimeID"`

	// PaymentID *uint
	// Payment   *Payment `gorm:"foreignKey:PaymentID"`

	// //onetomany
	// Seat []Seat `gorm:"foreignKey:TicketID"`
}
