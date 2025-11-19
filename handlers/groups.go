package handlers

import (
	"strconv"
	"tsuruev/database"
	"tsuruev/models"

	"github.com/gin-gonic/gin"
)

func GetGroups(c *gin.Context) {
	var groups []models.Group
	finished := c.Query("finished")
	week := c.Query("week")
	
	query := database.DB.Model(&models.Group{})

	if week != "" {
		Week,err := strconv.Atoi(week)
		if err != nil{
			c.JSON(400,gin.H{
				"error" : err.Error(),
			})
		}
		query = query.Where("current_week",Week)
	}
	if finished != "" {
		query = query.Where("in_Finished=?",finished)
	}

	 result := query.Find(&groups)
	 if result.Error != nil {
		c.JSON(404, gin.H{
			"Error": "NOT fOUND",
		})
		return
	}

	c.JSON(200, groups)

}

func CreateGroup(c *gin.Context) {
	var post models.GroupForPost

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var group models.Group
	group.Title = post.Title
	group.CurrentWeek = post.CurrentWeek
	group.TotalWeeks = post.TotalWeeks
	group.InFinished = post.InFinished

	if result := database.DB.Create(&group); result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(201, group)
}

func GetGroupsByFinished(c *gin.Context) {
	var groups []models.Group

	result := database.DB.Where("finished = ?", c.Query("finished")).Find(&groups)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"error": "Groups not found",
		})
		return
	}

	c.JSON(200, groups)
}

func UpdateGroupID(c *gin.Context) {
	var group models.Group

	id := c.Param("id")
	ID,err := strconv.Atoi(id)
	if err != nil{
		c.JSON(400,gin.H{
			"error" : err,
		})
	}

	if result := database.DB.First(&group, ID); result.Error != nil {
		c.JSON(400, gin.H{"error": "error"})
		return
	}

	var l models.GroupForPatch

	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := database.DB.Model(&group).Updates(l); result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, group)
}

func DeleteGroup(c *gin.Context) {
	var group models.Group

	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}


	if result := database.DB.First(&group, ID); result.Error != nil {
		c.JSON(404, gin.H{"error": "Group not found"})
		return
	}


	if result := database.DB.Delete(&group); result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Group deleted",
		"deleted": group,
	})
}

func GetGroupsID(c *gin.Context) {
	var group models.Group

	id := c.Param("id")

	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}


	if result:= database.DB.First(&group, ID); result.Error != nil {
		c.JSON(404, gin.H{"error": "Group not found"})
		return
	}

	c.JSON(200, group)
}