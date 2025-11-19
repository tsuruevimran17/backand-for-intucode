package handlers

import (
	"strconv"
	"tsuruev/database"
	"tsuruev/models"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var post models.StudentForPost
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	student := models.Student{
		FullName:      post.FullName,
		Email:         post.Email,
		Telegram:      post.Telegram,
		GroupID:       post.GroupID,
		TuitionPaid:   post.TuitionPaid,
		TuitionTotal:  post.TuitionTotal,
		PaymentStatus: post.PaymentStatus,
		StudyStatus:   post.StudyStatus,
	}

	result := database.DB.Preload("Group").Create(&student)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, student)
}

func GetStudents(c *gin.Context) {
	var students []models.Student

	groupIDStr := c.Query("group_id")
	paymentStatus := c.Query("payment_status")

	query := database.DB.Model(&models.Student{})

	if groupIDStr != "" {
		groupID, err := strconv.Atoi(groupIDStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid group_id"})
			return
		}
		query = query.Where("group_id = ?", groupID)
	}

	if paymentStatus != "" {
		if paymentStatus != "paid" && paymentStatus != "unpaid" && paymentStatus != "partial" {
			c.JSON(400, gin.H{"error": "Invalid payment_status"})
			return
		}
		query = query.Where("payment_status = ?", paymentStatus)
	}
	
	result := query.Find(&students)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, students)
}

func GetStudentByID(c *gin.Context) {
	var student models.Student

	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if result := database.DB.Preload("Group").First(&student, ID); result.Error != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(200, student)
}

func UpdateStudentID(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if result := database.DB.First(&student, ID); result.Error != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}

	var l models.StudentForPatch

	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := database.DB.Model(&student).Updates(&l); result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if result := database.DB.First(&student, ID); result.Error != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}

	if result := database.DB.Delete(&student); result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Student deleted",
		"deleted": student,
	})
}
func GetStudentsByGroupID(c *gin.Context) {
	var students []models.Student

	groupIDStr := c.Param("id")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid group_id"})
		return
	}

	result := database.DB.Preload("Group").Model(&models.Student{}).Where("group_id = ?", groupID).Find(&students)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	if len(students) == 0 {
		c.JSON(404, gin.H{"message": "No students found for this group"})
		return
	}

	c.JSON(200, students)
}
