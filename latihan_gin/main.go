package main

import (
	"latihan_gin/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/getUser", controller.GetUser)
	r.POST("/insertUser", controller.InsertUser)
	r.PUT("/updateUser/:id", controller.UpdateUser)
	r.DELETE("/deleteUser/:id", controller.DeleteUser)

	r.Run(":8888")
}
