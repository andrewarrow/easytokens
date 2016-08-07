package main

import "fmt"
import "github.com/gin-gonic/gin"
//import "net/http"

func main() {
  fmt.Println("s")
  r := gin.Default()
  r.LoadHTMLGlob("templates/*.tmpl")
  //r.Static("/assets", "./assets")

  r.GET("/", func(c *gin.Context) {
    data := gin.H{"nav_dash": true}
    c.HTML(200, "index.tmpl", data)
  })  
  r.Run(":80") 
}
