package main

import (
    "os"
    "github.com/gin-gonic/gin"
    "net/http"
)
    
type satellites struct {
   Satellites []struct {
        Name     string   `json:"name"`
        Distance float64  `json:"distance"`
        Message  []string `json:"message"`
    } `json:"satellites"`
}

type resistenciamensaje struct {
    Name     string  `json:"name"`
    Distance  float64  `json:"distance"`
    Message string  `json:"message"`
}

var mensajes = []resistenciamensaje{
    {Name: "kenobi", Distance: 100.0, Message: "['','']"},
    {Name: "skywalker", Distance: 115.5, Message:"['','']"},
    {Name: "skywalker", Distance: 115.5, Message:"['','']"},
}


func main() {
   port := os.Getenv("PORT")
   app := gin.New()

    app.GET("/topsecret_split/:satellite_name", func(c *gin.Context) {
        satellite_name := c.Param("satellite_name")
        c.String(http.StatusOK, "Hello %s", satellite_name)
         for _, a := range mensajes {
        if a.Name == satellite_name {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Tus ojos pueden engañarte, no confíes en ellos - Obi-Wan \n"})
    })

   app.POST("/topsecret/",func(context *gin.Context){
    body:=satellites{}
    if err:=context.BindJSON(&body);err!=nil{
         context.AbortWithError(http.StatusBadRequest,err)
         return
      }
      message := string(`RESPONSE CODE: 200{"position": { "x": -100.0,"y": 75.5 },"message": "este es un mensaje secreto" }`)
      context.JSON(http.StatusAccepted,&message)

   })
   app.Run(":"+port)

}


