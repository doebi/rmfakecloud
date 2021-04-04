package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *App) registerRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		count := app.hub.ClientCount()
		c.String(http.StatusOK, "Working, %d clients", count)
	})
	// register  a new device
	router.POST("/token/json/2/device/new", app.newDevice)

	// renew device acces token
	router.POST("/token/json/2/user/new", app.newUserToken)

	//service locator
	router.GET("/service/json/1/:service", app.locateService)

	//some beta stuff from internal.cloud
	router.GET("/settings/v1/beta", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"enrolled": false, "available": true})
	})

	//some telemetry stuff from ping.
	router.POST("/v1/reports", func(c *gin.Context) {
		c.Status(http.StatusOK)
		/*
			TODO: reverse this protobuf thing

			body, err := ioutil.ReadAll(c.Request.Body)

			if err != nil {
				c.AbortWithStatus(500)
				return
			}
			name := uuid.New().String() + ".dump"
			ioutil.WriteFile(name, body, 0644)
			log.Info(hex.Dump(body))
			c.Status(400)
		*/
	})

	app.docStorer.RegisterRoutes(router)
	app.ui.RegisterRoutes(router)

	//routes needing api authentitcation
	authRoutes := router.Group("/")
	authRoutes.Use(app.authMiddleware())
	{

		//unregister device
		authRoutes.POST("/token/json/3/device/delete", func(c *gin.Context) {
			c.String(http.StatusNoContent, "")
		})

		// doucment notifications
		authRoutes.GET("/notifications/ws/json/1", app.connectWebSocket)

		authRoutes.PUT("/document-storage/json/2/upload/request", app.uploadRequest)

		authRoutes.PUT("/document-storage/json/2/upload/update-status", app.updateStatus)

		authRoutes.PUT("/document-storage/json/2/delete", app.deleteDocument)

		authRoutes.GET("/document-storage/json/2/docs", app.listDocuments)

		// send email
		authRoutes.POST("/api/v2/document", app.sendEmail)
		// hwr
		authRoutes.POST("/api/v1/page", app.handleHwr)
		//livesync
		authRoutes.GET("/livesync/ws/json/2/:authid/sub", func(c *gin.Context) {
			uid := c.GetString(userIDKey)
			deviceID := c.GetString(deviceIDKey)
			app.hub.ConnectWs(uid, deviceID, c.Writer, c.Request)
		})
	}
}
