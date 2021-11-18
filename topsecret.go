package main

import (
    "os"
    "github.com/gin-gonic/gin"
    "net/http"

)


func postSatellites(c *gin.Context) {
    var resistencia =""

    if err := c.BindJSON(&resistencia); err != nil {
        return
    }

    c.IndentedJSON(http.StatusCreated, resistencia)
}


func main() {
   port := os.Getenv("PORT")
   app := gin.New()
   app.POST("/topsecret/",postSatellites)
   app.GET("/topsecret/",func (ctx *gin.Context){
    message := ctx.PostForm("este es un mensaje de la resistencia")
    satelite := ctx.DefaultPostForm("Kenoby","Obi Wan")
    ctx.JSON(200,gin.H{
        "status": "posted",
        "message" : message,
        "satelite": satelite,
    })

   })

   app.Run(":"+port)

}