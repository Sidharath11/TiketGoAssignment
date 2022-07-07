package univeristyinfocontroller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"university/domain"
	"university/domain/dto"
	"university/services"
	"university/utils"

	"github.com/gin-gonic/gin"
)

const (
	UNIINFOKEY string = "uniinfo"
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
	json, err := json.Marshal(newuser)
	if err != nil {
		rsterr := utils.InternalErr("while setting cache json")
		c.JSON(rsterr.Status, rsterr)
		return
	}
	err = domain.RedisClient.Set(UNIINFOKEY+strconv.Itoa(newuser.University_id), json, 0).Err()
	if err != nil {
		rsterr := utils.InternalErr("while setting redis cache")
		c.JSON(rsterr.Status, rsterr)
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
	val, err := domain.RedisClient.Get(UNIINFOKEY + string(id)).Bytes()
	if err != nil {
		rsterr := utils.InternalErr("while getting redis cache")
		c.JSON(rsterr.Status, rsterr)

		user, restErr := services.FindUniversityInfo(i)
		if restErr != nil {
			c.JSON(restErr.Status, restErr)
			return
		}
		c.JSON(http.StatusOK, user)

		return
	}
	c.JSON(http.StatusOK, utils.ToInfoJson(val))

}

func FindAllInfo(c *gin.Context) {
	var cursor uint64
	var info []dto.UniversityInfo

	iter := domain.RedisClient.Scan(cursor, "prefix:"+UNIINFOKEY, 0).Iterator()
	for iter.Next() {
		info = append(info, utils.ToInfoJson([]byte(iter.Val())))
	}
	if err := iter.Err(); err != nil {
		user, restErr := services.FindAllUniversityInfo()
		if restErr != nil {
			c.JSON(restErr.Status, restErr)
			return
		}
		c.JSON(http.StatusOK, user)
		return
	}

	c.JSON(http.StatusOK, info)

}

func DeleteInfo(c *gin.Context) {
	id := c.Query("universityid")
	_, err := domain.RedisClient.Del(UNIINFOKEY + string(id)).Result()

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.InternalErr("Error while deleting from cache"))
		return
	}
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
	_, err := domain.RedisClient.Del(UNIINFOKEY + strconv.Itoa(info.University_id)).Result()

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.InternalErr("Error while deleting from cache for update"))
		return
	}
	newuser, restErr := services.UpdateUniversityInfo(&info)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	json, err := json.Marshal(newuser)
	if err != nil {
		rsterr := utils.InternalErr("while setting cache json after update")
		c.JSON(rsterr.Status, rsterr)
		return
	}
	err = domain.RedisClient.Set(UNIINFOKEY+strconv.Itoa(newuser.University_id), json, 0).Err()
	if err != nil {
		rsterr := utils.InternalErr("while setting redis cache after update")
		c.JSON(rsterr.Status, rsterr)
		return
	}
	c.JSON(http.StatusCreated, newuser)
	return
}
