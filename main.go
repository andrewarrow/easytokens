package main

import "fmt"
import "github.com/gin-gonic/gin"
import "github.com/gin-gonic/contrib/sessions"
import "encoding/json"
import "net/http"
import "os"
import "io/ioutil"

func main() {
  fmt.Println("s")
  r := gin.Default()
  store := sessions.NewCookieStore([]byte("secret"))
  r.Use(sessions.Sessions("mysession", store))
  r.LoadHTMLGlob("templates/*.tmpl")
  //r.Static("/assets", "./assets")

  r.GET("/", func(c *gin.Context) {
    data := gin.H{"nav_dash": true}
    c.HTML(200, "index.tmpl", data)
  })  
  r.GET("/its_so_easy/redirect", func(c *gin.Context) {
    code := c.Query("code")
    session := sessions.Default(c)
    var team string
    v := session.Get("team")
    team = v.(string)

    url := "https://" + team + ".slack.com/api/oauth.access?client_id=" + 
       os.Getenv("EASY_CID") + "&client_secret=" + os.Getenv("EASY_SEC") + 
       "&code=" + code
    res, _ := http.Get(url)
    data, _ := ioutil.ReadAll(res.Body) 
    var thing map[string]string
    json.Unmarshal(data, &thing)
    res.Body.Close()
    c.JSON(200, thing["access_token"])
  })  
  r.POST("/auth", func(c *gin.Context) {
    team := c.PostForm("team")
    session := sessions.Default(c)
    session.Set("team", team)
    session.Save()
    c.Redirect(http.StatusMovedPermanently, "https://" + team + 
     ".slack.com/oauth/authorize?client_id=" + os.Getenv("EASY_CID") + "&scope=client")
  })  
  r.Run(":80") 
}
