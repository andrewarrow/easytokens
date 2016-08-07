package main

import "fmt"
import "github.com/gin-gonic/gin"
//import "net/http"
import "os"

func main() {
  fmt.Println("s")
  r := gin.Default()
  r.LoadHTMLGlob("templates/*.tmpl")
  //r.Static("/assets", "./assets")

  r.GET("/", func(c *gin.Context) {
    data := gin.H{"nav_dash": true}
    c.HTML(200, "index.tmpl", data)
  })  
  r.POST("/auth", func(c *gin.Context) {
    team := c.PostForm("team")
    c.Redirect(http.StatusMovedPermanently, "https://" + team + 
     ".slack.com/oauth/authorize?client_id=" + os.Getenv("EASY_CID") + "&scope=client")
  })  
  r.Run(":80") 
}
