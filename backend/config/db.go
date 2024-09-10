package config

import (
	"fmt"
	// "time"
	"example.com/project/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {

	db.AutoMigrate(
		&entity.Member{},
		&entity.Movie{},
		&entity.Theater{},
		&entity.ShowTimes{},
		&entity.Seat{},     // สร้างหลัง Theater เพราะมี foreign key เชื่อมโยง
		// &entity.Booking{},  // สร้างหลังจากที่ Member และ Movie ถูกสร้างแล้ว
		// &entity.BookSeat{}, // สร้างหลังจาก Seat และ Booking
		&entity.Payment{},
		&entity.Ticket{}, 
		&entity.Ticket{},
		&entity.TicketCheck{},
	)

}
