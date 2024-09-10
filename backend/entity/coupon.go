package entity

import "gorm.io/gorm"

type Coupon struct {
	gorm.Model
	CouponName string
	Discount int
	Status string

	//FK
	MemberID *uint
	Member   Member `gorm:"foreignKey:MembersID"`

	//onetomany
	Payment []Payment `gorm:"foreignKey:CouponID"`
}