package entity

import (
	"time"

	"gorm.io/gorm"
)

type TicketCheck struct {
	gorm.Model
	TicketID  uint `gorm:"uniqueIndex"` // Use uniqueIndex to enforce uniqueness
	TimeStamp time.Time
	Status    string

	// Member    Member       `gorm:"foreignKey:MemberID"`
	Ticket Ticket `gorm:"foreignKey:TicketID"`
}
