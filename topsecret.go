package main

import (
    "os"
    "github.com/gin-gonic/gin"
    "net/http"
    "log"
)
    
type satellites struct {
   Satellites []struct {
        Name     string   `json:"name"`
        Distance float64  `json:"distance"`
        Message  []string `json:"message"`
    } `json:"satellites"`
}


func main() {
   port := os.Getenv("PORT")
   app := gin.New()

   app.GET("/topsecret_split/:satellite_name", func(c *gin.Context) {
        satellite_name := c.Param("satellite_name")
        c.String(http.StatusOK, "Hello %s", satellite_name)
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

func satellitesName(w http.ResponseWriter, r *http.Request) {

    keys, ok := r.URL.Query()["key"]
    
    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }

    // Query()["key"] will return an array of items, 
    // we only want the single item.
    key := keys[0]

    log.Println("Url Param 'key' is: " + string(key))
}

