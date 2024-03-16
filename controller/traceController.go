package traceController

import (
	databasech "golangproject/database"
	trace "golangproject/models/album"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTraces(c *gin.Context) {
	db := databasech.SetupDatabase()
	var traces []trace.Trace
	ctx := c.Request.Context()
	projectId := c.MustGet("project").(string)
	println("project", projectId)
	err := db.NewSelect().Model(&traces).Where("\"projectId\" = ?", projectId).Scan(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
    c.IndentedJSON(http.StatusOK, traces)
}

func GetTraceById(c *gin.Context) {
	db := databasech.SetupDatabase()
	var trace trace.Trace
	ctx := c.Request.Context()
	err := db.NewSelect().Model(&trace).Where("id = ?", c.Param("id")).Scan(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, trace)
}

func AddTrace(c *gin.Context) {
	db := databasech.SetupDatabase()
	var trace trace.Trace
	ctx := c.Request.Context()
	err := c.BindJSON(&trace)
	projectId := c.MustGet("project").(string)
	println("projectId", projectId)
	trace.ProjectId = projectId
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	println("trace", trace.ProjectId)
	res,err := db.NewInsert().Model(&trace).Exec(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}