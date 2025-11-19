package main

import (
	"tsuruev/database"
	"tsuruev/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connection()
	r := gin.Default()

	r.POST("/students", handlers.CreateStudent)
	r.GET("/students", handlers.GetStudents)
	r.GET("/groups/:id/students", handlers.GetStudentsByGroupID)
	r.GET("/students/:id", handlers.GetStudentByID)
	r.PATCH("/students/:id", handlers.UpdateStudentID)
	r.DELETE("/students/:id", handlers.DeleteStudent)

	r.GET("/groups", handlers.GetGroups)
	r.GET("/groups/:id", handlers.GetGroupsID)
	r.POST("/groups", handlers.CreateGroup)
	r.PATCH("/groups/:id", handlers.UpdateGroupID)
	r.DELETE("/groups/:id", handlers.DeleteGroup)

	r.GET("/students/:id/notes", handlers.GetNotesByStudentID)
	r.GET("/notes",handlers.GetNotes)
	r.POST("/notes", handlers.CreateNote)
	r.PATCH("/notes/:id", handlers.UpdateNoteID)
	r.DELETE("/notes/:id", handlers.DeleteNoteID)

	r.Run(":8081")
}
