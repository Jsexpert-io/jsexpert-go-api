package trace

type  TraceRaw struct{
	ResourceSpans []ResourceSpan `json:"resourceSpans"`
}
type  ResourceSpan struct{
	Resource Resource `json:"resource"`
	ScopeSpans []ScopeSpan `json:"scopeSpans"`
}



type Attribute struct{

	Key string `json:"key" `
	Value interface{} `json:"value"` // Can be any type
}
type  Resource struct{
	Attributes []Attribute `json:"attributes"`
	droppedAttributesCount int

}
type Scope struct{
	Name string `json:"name"`
}
type  ScopeSpan struct{
	Scope Scope	`json:"scope"`// Can be any type
	Spans []Span `json:"spans"`
}
type  Span struct{
	
	TraceId string `json:"traceId"`
	SpanId string `json:"spanId"`
	ParentSpanId string `json:"parentSpanId"`
	Name string `json:"name"`
	Kind int `json:"kind"`
	StartTimeUnixNano string `json:"startTimeUnixNano"`
	EndTimeUnixNano string `json:"endTimeUnixNano"`
	DroppedAttributesCount int `json:"droppedAttributesCount"`
	Events []interface{} `json:"events"`
	DroppedEventsCount int `json:"droppedEventsCount"`
	Attributes []Attribute `json:"attributes"`
	Status interface{} `json:"status"`
	Links []interface{} `json:"links"`
	ProjectId string `json:"projectId"`
	ScopeName string `json:"scopeName"`
	DroppedLinksCount int `json:"droppedLinksCount"`

		
}