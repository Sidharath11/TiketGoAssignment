package univeristycontroller

import (
	"net/http"
	"strconv"
	"university/domain/dto"
	"university/services"
	"university/utils"

	"github.com/gin-gonic/gin"
)

func CreateUniversity(c *gin.Context) {
	var info dto.University
	if err := c.ShouldBindJSON(&info); err != nil {
		restErr := utils.BadRequest("Invalid json")
		c.JSON(restErr.Status, restErr)
	}

	newuser, restErr := services.CreateUniversity(&info)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, newuser)
}

func FindUniversity(c *gin.Context) {
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
	user, restErr := services.FindUniversity(i)
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
	restErr := services.DeleteUniversity(i)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"isRemoved": true})
}

func UpdateInfo(c *gin.Context) {
	var info dto.University
	if err := c.ShouldBindJSON(&info); err != nil {
		restErr := utils.BadRequest("Invalid json")
		c.JSON(restErr.Status, restErr)
	}

	newuser, restErr := services.UpdateUniversity(&info)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, newuser)
}
