package main

import (
	"github.com/gin-gonic/gin"
	"go-dummyapp-moysklad/internal/server"
//        "encoding/json"
//        "fmt"
	"log"
//	"os"
//	"io"
)

const appBaseUrl = "/echo/api/moysklad/vendor/1.0/apps/:appId/:accountId"

func main() {
	appServer, err := server.NewServer()
	if err != nil {
		log.Fatalln(err)
	}
	//var f, _ = os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()

	router.PUT(appBaseUrl, appServer.ActivateHandler)
	router.DELETE(appBaseUrl, appServer.DeleteHandler)
	router.GET(appBaseUrl, appServer.StatusHandler)

	router.LoadHTMLFiles("internal/test2.sorochinsky/iframe.html")

	router.GET("/echo/iframe/:appUid", appServer.IframeHandler)
	router.POST("/echo/:appUid/update-settings", appServer.UpdateSettingsHandler)

	router.Run(":8002")
}
