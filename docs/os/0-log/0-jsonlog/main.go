package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

// JsonOutputsForLogger is for JSON marshaling in sequence.
type JsonOutputsForLogger struct {
	TraceId string `json:"trace_id"`
	Time    string `json:"time"`
	Level   string `json:"level"`
	Content string `json:"content"`
}

// LoggingJsonHandler is a example handler for logging JSON format content.
var LoggingJsonHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
	jsonStruc := JsonOutputsForLogger{
		TraceId: in.TraceId,
		Time:    in.Time.Format("2006-01-02 15:04:05.000"),
		Level:   fmt.Sprintf("%d", in.Level),
		Content: in.Content,
	}
	b, err := json.Marshal(jsonStruc)
	if err != nil {
		in.Next(ctx)
	}
	in.Buffer.Write(b)
	in.Buffer.WriteString("\n")
	in.Next(ctx)
}

func main() {
	g.Log().SetHandlers(LoggingJsonHandler)
	ctx := gctx.New()
	g.Log().Debug(ctx, "Debugging...")
	g.Log().Warning(ctx, "It is warning info")
	g.Log().Error(ctx, "Error occurs, please have a check")
}
