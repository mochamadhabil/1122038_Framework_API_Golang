package controller

import (
	m "latihan_gin/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	db, err := connectGorm()

	if err != nil {
		SendErrorResponse(c, 500, "Internal Server Error")
	}

	var user m.User

	query := db.Find(&user)
	if query.Error != nil {
		SendErrorResponse(c, 500, "Internal Server Error")
	}

	c.JSON(http.StatusOK, user)
}

func InsertUser(c *gin.Context) {
	db, err := connectGorm()
	if err != nil {
		SendErrorResponse(c, 500, "internal server error")
		return
	}

	name := c.Query("Name")
	address := c.Query("Address")
	ageStr := c.Query("Age")
	email := c.Query("Email")
	passw := c.Query("Passw")

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		SendErrorResponse(c, 400, "Usia Tidak Valid")
		return
	}

	user := m.User{
		Name:    name,
		Address: address,
		Age:     age,
		Email:   email,
		Passw:   passw,
	}

	result := db.Create(&user)
	err = result.Error
	if err != nil {
		SendErrorResponse(c, 500, "internal server error")
		return
	}
	SendSuccessResponse(c, 200, "Success")
}

func UpdateUser(c *gin.Context) {
	db, err := connectGorm()

	if err = c.Request.ParseForm(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Parsing Data"})
		return
	}

	userID := c.Param("id")

	var user m.User

	user.Name = c.PostForm("Name")

	ageStr := c.PostForm("Age") // bisa diganti dengan c.Query tapi isi parameter di postman nya di bagian params
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usia Tidak Valid"})
		return
	}

	user.Address = c.PostForm("Address")
	user.Email = c.PostForm("Email")
	user.Passw = c.PostForm("Passw")

	result := db.Model(&user).Where("id = ?", userID).Update("Age", age)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saat update User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func DeleteUser(c *gin.Context) {
	db, err := connectGorm()
	if err != nil {
		SendErrorResponse(c, 500, "Internal Server Error")
		return
	}

	userID := c.Param("id")

	var user m.User
	if err := db.Delete(&user, &userID).Error; err != nil {
		SendErrorResponse(c, 500, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, user)
}
