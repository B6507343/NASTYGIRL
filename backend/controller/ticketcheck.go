package controller

import (
	"net/http"

	"example.com/project/config"
	"example.com/project/entity"
	"github.com/gin-gonic/gin"
	// "github.com/siriphobmean/gear-registry/entity"
)

// POST /students
func Checkin(c *gin.Context) {
	var ticketcheck entity.TicketCheck

	// bind เข้าตัวแปร student
	if err := c.ShouldBindJSON(&ticketcheck); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// สร้าง Student
	u := entity.TicketCheck{
		TicketID: ticketcheck.TicketID,
		TimeStamp: ticketcheck.TimeStamp,
		Status: ticketcheck.Status,
	}

	// บันทึก
	if err := config.DB().Create(&u).Error; err != nil {
		if err.Error() == "UNIQUE constraint failed: ticketcheck.TicketID" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Already use"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": u})
}

// GET /student/:id
func GetTicketCheck(c *gin.Context) {
	var ticketcheck entity.TicketCheck
	id := c.Param("id")
	if err := config.DB().Raw("SELECT * FROM ticketcheck WHERE id = ?", id).Find(&ticketcheck).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ticketcheck})
}

// GET /students
func ListStudents(c *gin.Context) {
	var ticketcheck[] entity.TicketCheck
	if err := config.DB().Raw("SELECT * FROM ticketcheck").Find(&ticketcheck).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ticketcheck})
}

// DELETE /students/:id
func DeleteTicketcheck(c *gin.Context) {
	id := c.Param("id")
	if tx := config.DB().Exec("DELETE FROM students WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /students
// func UpdateStudent(c *gin.Context) {
// 	var student entity.Student
// 	var result entity.Student

// 	if err := c.ShouldBindJSON(&student); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
	
//     if tx := entity.DB().Where("students_id = ?", student.StudentsID).First(&result); tx.RowsAffected == 0 {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่มีรหัสนักศึกษาในระบบ"})
//         return
//     }

//     // Update only the Status field
//     result.Status = student.Status
// 	if err := entity.DB().Save(&result).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": result})
// }