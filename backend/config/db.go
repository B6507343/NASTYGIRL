package config

import (
	"fmt"
	

	"example.com/project/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// ฟังก์ชันสำหรับคืนค่าตัวแปร db
func DB() *gorm.DB {
	return db
}

// ฟังก์ชันสำหรับการเชื่อมต่อฐานข้อมูล
func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

// ฟังก์ชันสำหรับการตั้งค่าและสร้างตารางในฐานข้อมูล
func SetupDatabase() {

	// AutoMigrate สำหรับตาราง member, ticket, ticketcheck
	err := db.AutoMigrate(
		&entity.Member{},
		&entity.Ticket{},
		&entity.TicketCheck{},
	)
	if err != nil {
		fmt.Println("Error in AutoMigrate:", err)
	} else {
		fmt.Println("AutoMigrate completed successfully.")
	}

	// สร้างข้อมูลสมาชิกคนที่หนึ่ง (sa@gmail.com)
	hashedPassword, _ := HashPassword("123456")
	Member1 := &entity.Member{
		UserName:   "sa1",
		FirstName:  "Software1",
		LastName:   "Analysis1",
		Email:      "sa@gmail.com",
		Password:   hashedPassword,
		TotalPoint: 10,
		Role:       "customer",
	}
	db.FirstOrCreate(Member1, &entity.Member{
		Email: "sa@gmail.com",
	})

	// สร้างข้อมูลสมาชิกคนที่สอง (sa2@gmail.com)
	hashedPassword2, _ := HashPassword("123456")
	Member2 := &entity.Member{
		UserName:   "sa2",
		FirstName:  "Software2",
		LastName:   "Analysis2",
		Email:      "sa2@gmail.com",
		Password:   hashedPassword2,
		TotalPoint: 5,
		Role:       "customer",
	}
	db.FirstOrCreate(Member2, &entity.Member{
		Email: "sa2@gmail.com",
	})

	// สร้าง tickets สำหรับสมาชิกที่สอง (Member2)
	tickets := []entity.Ticket{
		{Point: 10, Status: "Booked", MemberID: Member2.ID},
		{Point: 15, Status: "Booked", MemberID: Member2.ID},
	}

	for i := range tickets {
		if err := db.Create(&tickets[i]).Error; err != nil {
			fmt.Printf("Error creating ticket: %s\n", err)
		}
	}

	

	fmt.Println("Database setup complete")
}
