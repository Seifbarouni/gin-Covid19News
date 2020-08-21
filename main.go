package main

import (
	"projects/GinFramework/gin-Covid/controllers"

	"github.com/gin-gonic/gin"
)


func main(){
	r:=gin.Default()
	r.LoadHTMLGlob("views/assets/*")
	r.Static("/static","./views/static")

	r.GET("/",controllers.Redirect)
	r.GET("/corona",controllers.RenderHTML)

	r.Run()
}