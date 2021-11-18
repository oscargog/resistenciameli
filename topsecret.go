package main

import (
    "os"
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
)

type Body struct {
  // json tag to serialize json body
   Name string `json:"name"`
}

func main() {
   port := os.Getenv("PORT")
   app := gin.New()

   app.POST("/topsecret/",func(context *gin.Context){

    body:=Body{}
    if err:=context.BindJSON(&body);err!=nil{
         context.AbortWithError(http.StatusBadRequest,err)
         return
      }
      fmt.Println(body)
      context.JSON(http.StatusAccepted,&body)
    
   })
   
   app.Run(":"+port)

}