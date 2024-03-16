package traceController

import (
	"context"
	"encoding/json"
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

func SaveTraces(ctx context.Context,spans []trace.Span) {
	db := databasech.SetupDatabase()
	
	for _, span := range spans {
		
		trace := trace.Trace{
			Name: 				  span.Name,
			TraceId:              span.TraceId,
			SpanId:               span.SpanId,
			ParentSpanId:         span.ParentSpanId,
			Kind:                 int64(span.Kind),
			StartTimeUnixNano:    span.StartTimeUnixNano,
			EndTimeUnixNano:      span.EndTimeUnixNano,
	
			DroppedAttributesCount: uint64(span.DroppedAttributesCount),

			DroppedEventsCount:   int64(span.DroppedEventsCount),

			DroppedLinksCount:    int64(span.DroppedLinksCount),
			ProjectId:            span.ProjectId,
			
		}
		attributes, err :=json.Marshal(span.Attributes)
		if err != nil {
			println("Error marshalling attributes", err)
		}
		trace.Attributes = string(attributes)
		events, err :=json.Marshal(span.Events)
		if err != nil {
			println("Error marshalling events", err)
		}
		trace.Events = string(events)
		links, err :=json.Marshal(span.Links)
		if err != nil {
			println("Error marshalling links", err)
		}
		trace.Links = string(links)



		_, ersr := db.NewInsert().Model(&trace).Exec(ctx)
		if ersr != nil {
			println("Error saving trace", err)
		}
	}
}

func MapTrace(c *gin.Context) {
	
	var createTraceDto trace.TraceRaw
	projectId := c.MustGet("project").(string)
	println("project", projectId)
	err := c.BindJSON(&createTraceDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
 	var spans []trace.Span

    //resourceAttributes := createTraceDto.ResourceSpans[0].Resource.Attributes
	for _,rs := range createTraceDto.ResourceSpans{
		for _,ss := range rs.ScopeSpans{
			if(ss.Spans == nil){
				continue
			}
			for _,s := range ss.Spans{
				s.ProjectId = projectId
				s.ScopeName = ss.Scope.Name
			
				var eventAttributes []trace.Attribute
				for _,e := range s.Events{
					eventAttributes = append(eventAttributes, e.(trace.Attribute))
				}
				var totalAttributes []trace.Attribute
			
				totalAttributes = append(totalAttributes, eventAttributes...)
				totalAttributes = append(totalAttributes, s.Attributes...)
				totalAttributes = append(totalAttributes, rs.Resource.Attributes...)
				s.Attributes = totalAttributes
				spans = append(spans, s)
			}
		}
	}
	SaveTraces(c.Request.Context(),spans)
	// map function in go 
	// https://golangdocs.com/map-function-in-golang
	c.IndentedJSON(http.StatusOK, spans)
}