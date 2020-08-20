package main

import (
	"projects/GinFramework/gin-Covid/controllers"

	"github.com/gin-gonic/gin"
)


func main(){
	r:=gin.Default()
	r.LoadHTMLGlob("assets/*")
	r.Static("/static","./static")

	r.GET("/",controllers.Redirect)
	r.GET("/corona",controllers.RenderHTML)

	r.Run()
}