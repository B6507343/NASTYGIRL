package entity

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	TotalPrice int
	Status     string
	TimeStamp  time.Time

	//FK
	MemberID *uint
	Member   Member `gorm:"foreignKey:MemberID"`

	TicketID *uint
	Ticket   *Ticket `gorm:"foreignKey:TicketID"`
}