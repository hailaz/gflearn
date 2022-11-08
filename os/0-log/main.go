package main

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

// JsonOutputsForLogger is for JSON marshaling in sequence.
type JsonOutputsForLogger struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Content string `json:"content"`
}

// LoggingJsonHandler is a example handler for logging JSON format content.
var LoggingJsonHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
	str := ""
	// TODO: something
	config := in.Logger.GetConfig()
	if in.TraceId != "" {
		str += "{" + in.TraceId + "}"
	} else {
		str += "{}"
	}
	if len(config.CtxKeys) > 0 {
		for _, ctxKey := range config.CtxKeys {
			var ctxValue interface{}
			if ctxValue = ctx.Value(ctxKey); ctxValue == nil {
				ctxValue = ctx.Value(gctx.StrKey(gconv.String(ctxKey)))
			}
			if ctxValue != nil {
				str += "{" + gconv.String(ctxValue) + "}"
			} else {

				str += "{}"
			}
		}
	}
	in.Buffer.WriteString(str)
	in.Buffer.WriteString(in.Content)
	in.Buffer.WriteString("\n")
	in.Next(ctx)
}

func main() {
	g.Log().SetHandlers(LoggingJsonHandler)
	ctx := context.TODO()
	g.Log().SetCtxKeys("Trace-Id", "Span-Id", "Test")
	ctx = context.WithValue(ctx, "Trace-Id", "1234567890")
	ctx = context.WithValue(ctx, "Span-Id", "abcdefg")
	g.Log().Debug(ctx, "Debugging...")
	g.Log().Warning(ctx, "It is warning info")
	g.Log().Error(ctx, "Error occurs, please have a check")
}
