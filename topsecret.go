package main

import (
    "os"
    "github.com/gin-gonic/gin"
)
func main() {
   port := os.Getenv("PORT")
   app := gin.New()
   app.GET("/",func (ctx *gin.Context){
    ctx.String(200, "Esto esta funcionando")

   })
   app.Run(":"+port)
}