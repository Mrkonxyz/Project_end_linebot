package controller

import (
	"net/http"

	"github.com/Hitmepls/my-project/entity"
	"github.com/gin-gonic/gin"
)

func CreateAccidents(c *gin.Context) {
	var user entity.Accident
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /user/:id
func GetAccidents(c *gin.Context) {
	var user []entity.Accident
	id := c.Param("id")
	if err := entity.DB().Preload("Reporter").Preload("Level").Preload("ProcessStatus").Raw("SELECT * FROM Accidents WHERE reporter_id = ? ORDER by time DESC LIMIT 3", id).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /count acident
func GetCountAccidentsActive(c *gin.Context) {
	var count int64
	if err := entity.DB().Raw("SELECT count(id) FROM Accidents where process_status_id < 3").Find(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": count})
}

// GET /all acident
func ListAccidents(c *gin.Context) {
	var users []entity.Accident
	if err := entity.DB().Preload("Reporter").Preload("Level").Preload("ProcessStatus").Raw("SELECT * FROM Accidents").Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// func ListAccidentsDaysAgo(c *gin.Context) {
// 	var users []entity.Accident
// 	// days := c.Param("days")
// 	// .Add((time.Hour*-24)*7)
// 	if err := entity.DB().Preload("Reporter").Preload("Level").Preload("ProcessStatus").Raw("SELECT * FROM Accidents where date_time  BETWEEN ? and ?", time.Now(), time.Now()).Find(&users).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

//		c.JSON(http.StatusOK, gin.H{"data": users})
//	}
func ListAccidentsFromDate(c *gin.Context) {
	var users []entity.Accident
	date1 := c.Param("date1")
	date2 := c.Param("date2")
	if err := entity.DB().Preload("Reporter").Preload("Level").Preload("ProcessStatus").Raw("SELECT * FROM Accidents where time  BETWEEN ? and ?", date1, date2[1:]).Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /acident no status 3
func ListAccidentsActive(c *gin.Context) {
	var users []entity.Accident
	if err := entity.DB().Preload("Reporter").Preload("Level").Preload("ProcessStatus").Raw("SELECT * FROM Accidents WHERE process_status_id < 3").Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}
