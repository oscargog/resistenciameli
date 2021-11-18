package main

import (
    "os"
    "github.com/gin-gonic/gin"
    "net/http"
    
)

    
type Body struct {
    Satellites []struct {
        Name     string   `json:"name"`
        Distance float64  `json:"distance"`
        Message  []string `json:"message"`
    } `json:"satellites"`
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
      //fmt.Println(body)
      message := []byte(`{"position": { "x": -100.0,"y": 75.5 },"message": "este es un mensaje secreto" }`)
      context.JSON(http.StatusAccepted,&message)

   })
   
   app.Run(":"+port)

}