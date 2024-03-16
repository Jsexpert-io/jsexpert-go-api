package trace

import (
	"time"

	"github.com/uptrace/go-clickhouse/ch"
)

// Trace represents the structure of your table data in Go.
type Trace struct {
	ch.CHModel `ch:"table:jsexpertdb.trace_data,partition:toYYYYMM(time)"`
	ID                      string    `ch:"id"`
	Name                    string    `db:"name"`
	scopeName               string    `ch:"scopeName"`
	TraceId                 string    `ch:"traceId"`
	SpanId                  string    `ch:"spanId"`
	ParentSpanId            string    `ch:"parentSpanId"`
	Kind                    int64     `ch:"kind"`
	StartTimeUnixNano       string    `ch:"startTimeUnixNano"`
	EndTimeUnixNano         string    `ch:"endTimeUnixNano"`
	Attributes              string    `ch:"attributes"`
	DroppedAttributesCount  uint64    `ch:"droppedAttributesCount"`
	Events                  string    `ch:"events"`
	DroppedEventsCount      int64     `ch:"droppedEventsCount"`
	Links                   string    `ch:"links"`
	DroppedLinksCount       int64     `ch:"droppedLinksCount"`
	ProjectId               string    `ch:"projectId"`
	CreatedAt               time.Time `ch:"createdAt"`
	UpdatedAt               time.Time `ch:"updatedAt"`
	Status                  string    `ch:"status"`
	

	// TraceId string `json:"traceId"`
	// SpanId string `json:"spanId"`
	// ParentSpanId string `json:"parentSpanId"`
	// Name string `json:"name"`
	// Kind int `json:"kind"`
	// StartTimeUnixNano string `json:"startTimeUnixNano"`
	// EndTimeUnixNano string `json:"endTimeUnixNano"`
	// DroppedAttributesCount int `json:"droppedAttributesCount"`
	// Events []interface{} `json:"events"`
	// DroppedEventsCount int `json:"droppedEventsCount"`
	// Attributes []Attribute `json:"attributes"`
	// Status interface{} `json:"status"`
	// Links []interface{} `json:"links"`
	// ProjectId string `json:"projectId"`
	// ScopeName string `json:"scopeName"`
	// DroppedLinksCount int `json:"droppedLinksCount"`

}


