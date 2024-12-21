package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)
var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


func wsHandler(w http.ResponseWriter,r *http.Request){
	conn, err :=wsUpgrader.Upgrade(w,r,nil)
	if err!=nil{
		log.Println("connection:-",err)
        return
	}
	for{
		t, msg, err:=conn.ReadMessage()
		if err != nil{
			break
		}
		msg= []byte("pong")
		conn.WriteMessage(t,msg)
	}
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
		// c.String(http.StatusOK, "pong")
	})
	r.GET("/ws",func(c* gin.Context){
		wsHandler(c.Writer,c.Request)
	})
	r.Run("localhost:9000")
}
