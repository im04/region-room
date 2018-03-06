package router

import (
	"github.com/gin-gonic/gin"
	"region-room/controller/user"
	"region-room/middleware"
)

func InitRouter (app *gin.Engine) {
	/*router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})*/
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	sign := app.Group("/sign")
	{
		sign.POST("/signUp", user.SignUp)
		sign.POST("/signIn", user.SignIn)
	}
	us := app.Group("/api")
	{
		us.Use(middleware.CheckToken())
		us.POST("/signOut", user.SignOut)
	}

}
