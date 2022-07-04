package univeristyinfocontroller

import (
	"net/http"
	"strconv"
	"university/domain/dto"
	"university/services"
	"university/utils"

	"github.com/gin-gonic/gin"
)

func CreateInfo(c *gin.Context) {
	var info dto.UniversityInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		restErr := utils.BadRequest("Invalid json")
		c.JSON(restErr.Status, restErr)
	}

	newuser, restErr := services.CreateUniversityInfo(&info)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, newuser)
}

func FindInfo(c *gin.Context) {
	id := c.Query("universityid")
	if id == "" {
		restErr := utils.BadRequest("university id is blank")
		c.JSON(restErr.Status, restErr)
		return
	}
	i, e := strconv.Atoi(id)
	if e != nil {
		restErr := utils.BadRequest("university id is not int type")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := services.FindUniversityInfo(i)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)

}

func DeleteInfo(c *gin.Context) {
	id := c.Query("universityid")
	if id == "" {
		restErr := utils.BadRequest("university id is blank")
		c.JSON(restErr.Status, restErr)
		return
	}
	i, e := strconv.Atoi(id)
	if e != nil {
		restErr := utils.BadRequest("university id is not int type")
		c.JSON(restErr.Status, restErr)
		return
	}
	restErr := services.DeleteUniversityInfo(i)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"isRemoved": true})
}

func UpdateInfo(c *gin.Context) {
	var info dto.UniversityInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		restErr := utils.BadRequest("Invalid json")
		c.JSON(restErr.Status, restErr)
	}

	newuser, restErr := services.UpdateUniversityInfo(&info)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, newuser)
}
