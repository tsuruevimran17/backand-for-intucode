package handlers

import (
	"strconv"
	"tsuruev/database"
	"tsuruev/models"

	"github.com/gin-gonic/gin"
)

func GetNotesByStudentID(c *gin.Context){
	var notes []models.Note
	id := c.Param("id")
	ID , err := strconv.Atoi(id)
	if err != nil{
		c.JSON(400,gin.H{
			"error" : err,
		})
		return
	}
	if result := database.DB.Preload("Student").Model(&models.Note{}).Where("student_id = ?",ID).Find(&notes);result.Error != nil{
		c.JSON(404, "not found")
	}

	if len(notes)==0{
		c.JSON(404,"No notes found for this student")
	}

	c.JSON(200,notes)

}

func CreateNote(c *gin.Context){
	var post models.NotePost

	if err := c.ShouldBindJSON(&post);err != nil{
		c.JSON(400,gin.H{
			"error":err.Error(),
		})
		return
	}
	if post.Text ==""{
		c.JSON(400,"Текст не должен быть пустым")
	}
	
	 note := models.Note{
		StudentId: post.StudentId,
		Author: post.Author,
		Text: post.Text,
	 }
	 if result := database.DB.Preload("Student").Create(&note); result.Error != nil{
		c.JSON(404,gin.H{
			"error": result.Error.Error(),
		})
	 }

	 c.JSON(200,note)
}

func UpdateNoteID(c *gin.Context){
	var note models.Note
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil{
		c.JSON(400,gin.H{
			"error": err.Error(),
		})
	}
	if result := database.DB.First(&note, ID); result.Error != nil{
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
	}

	var l models.NotePatch
	if err := c.ShouldBindJSON(&l);err != nil{
		c.JSON(400,gin.H{
			"Error": err.Error(),
		})
	}

	if result := database.DB.Model(&note).Updates(&l); result.Error != nil{
		c.JSON(400,gin.H{
			"error": result.Error.Error(),
		})
	}
	
	c.JSON(200,note)
}

func DeleteNoteID(c *gin.Context){
	var note models.Note
	id := c.Param("id")
	ID , err := strconv.Atoi(id)
	if err != nil{
		c.JSON(400,gin.H{
			"error": err.Error(),
		})
	}

	if result := database.DB.Delete(&note,ID);result.Error != nil{
		c.JSON(404,gin.H{
			"error" : result.Error.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "note deleted",
		"deleted": note,
	})

}
func GetNotes(c *gin.Context) {
    var notes []models.Note

    result := database.DB.Find(&notes)
    if result.Error != nil {
        c.JSON(500, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(200, notes)
}