package main

import (
	traceController "golangproject/controller"
	"golangproject/database"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
//   gin.SetMode(gin.ReleaseMode) 
  database.ConnectDatabase()
  database.AutoMigrateModels()
  pprof.Register(router, "dev/pprof")
  
  traceRouteGroup := router.Group("/v1/traces") 
  {

	//traceRouteGroup.Use(middelwear.ProjectAuth())
	traceRouteGroup.GET("", traceController.GetTraces)
	traceRouteGroup.GET("/:id", traceController.GetTraceById)
	traceRouteGroup.POST("", traceController.MapTrace)
	// albumsRoute.PUT("/:id", albumController.UpdateAlbum)
	// albumsRoute.DELETE("/:id", albumController.DeleteAlbum)
  }
  router.GET("/ping",func(c *gin.Context) {
	c.JSON(200, gin.H{"status": "Healthy"})
  })
  router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}