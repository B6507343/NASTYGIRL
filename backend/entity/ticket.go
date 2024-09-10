package entity

import (
	// "time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Point     int
	Status 	string
	MemberID uint
	
	Member   Member  `gorm:"foreignKey:MemberID"`

	// //FK
	// ShowTimesID *uint
	// ShowTime   ShowTimes `gorm:"foreignKey:ShowTimeID"`

	// MemberID *uint
	

	// PaymentID *uint
	// Payment   *Payment `gorm:"foreignKey:PaymentID"`

	// TicketCheckID *uint
	// TicketCheck   TicketCheck `gorm:"foreignKey:TicketCheckID"`

	//onetomany
	// Seat []Seat `gorm:"foreignKey:TicketID"`
}
