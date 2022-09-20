package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//All routes related to users comes here
	router.POST("/user", controllers.CreateUser())            //add this
	router.GET("/user/:userId", controllers.GetAUser())       //add this
	router.PUT("/user/:userId", controllers.EditAUser())      //add this
	router.DELETE("/user/:userId", controllers.DeleteAUser()) //add this
	router.GET("/users", controllers.GetAllUsers())           //add this

}
