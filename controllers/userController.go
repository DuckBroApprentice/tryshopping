package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DuckBroApprentice/Shopping/database"
	"github.com/DuckBroApprentice/Shopping/models"
	"github.com/gin-gonic/gin"
)

var userList = []models.User{}

// Get User
func FindAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}

// Post User
func PostUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error:"+err.Error())
	}
	userList = append(userList, user)
	c.JSON(http.StatusOK, "Successfully posted")
}

// Delete User
func DeleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	for _, user := range userList {
		log.Println(user)
		if user.Id == userId {
			userList = append(userList[:userId], userList[userId+1:]...)
		}
	}
}

// Put User
func PutUser(c *gin.Context) {
	beforeUser := models.User{}
	err := c.BindJSON(&beforeUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
	}
	userId, _ := strconv.Atoi(c.Param("id"))
	for key, user := range userList {
		if userId == user.Id {
			userList[key] = beforeUser
			log.Println(userList[key])
			c.JSON(http.StatusOK, "Successfully")
			return
		}
	}
	c.JSON(http.StatusNotFound, "Error")
}

// DBConnect Controller
func FindAllUsersFromDB(c *gin.Context) {
	var users []models.User
	database.DBConnect.Find(&users)
	c.JSON(http.StatusOK, users)
}

func FindUserById(userId int) models.User {
	var user models.User
	database.DBConnect.Where("id = ?", userId).First(&user)
	return user
}

// Sign Up User(POST)
func PostUserToDB(c *gin.Context) {
	var user models.User
	database.DBConnect.Set("NewUser", user)
}
