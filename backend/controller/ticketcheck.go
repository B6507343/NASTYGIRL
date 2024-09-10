package controller

import (
	"net/http"
	"strconv"
	"time"

	"example.com/project/config"
	"example.com/project/entity"
	"github.com/gin-gonic/gin"
)

// CreateTicketCheck - ฟังก์ชันสำหรับสร้าง ticket check
func CreateTicketCheck(c *gin.Context) {
	ticketIDStr := c.Param("ticket_id")
	ticketID, err := strconv.Atoi(ticketIDStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	// สร้างข้อมูล ticket check
	ticketCheck := entity.TicketCheck{
		TicketID:  uint(ticketID),
		TimeStamp: time.Now(),
		Status:    "Checked",
	}

	// บันทึกข้อมูลลงในฐานข้อมูล
	if err := config.DB().Create(&ticketCheck).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูลกลับไปยัง client
	c.JSON(http.StatusOK, gin.H{"success": true, "data": ticketCheck})
}
