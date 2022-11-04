package tracing

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	bplog "github.com/reddit/baseplate.go/log"
)

const (
	ctxKey = "context"

	// Note that this must be the same as
	// https://github.com/opentracing/opentracing-go/blob/v1.2.0/log/field.go#L128
	errorKey = "error.object"
)

// FinishOptions are the options to be converted into opentracing.FinishOptions.
//
// All fields are optional.
type FinishOptions struct {
	Ctx context.Context
	Err error
}

// Convert converts FinishOptions into opentracing.FinishOptions which can be
// used in Span.FinishWithOptions().
func (fo FinishOptions) Convert() opentracing.FinishOptions {
	var opts opentracing.FinishOptions
	var lr opentracing.LogRecord
	if fo.Ctx != nil {
		lr.Fields = append(lr.Fields, log.Object(ctxKey, fo.Ctx))
	}
	if fo.Err != nil {
		bplog.Warnw("in Convert err, fo.Err is: " + fmt.Sprintf("%s\n", fo.Err.Error()))
		bplog.Warnw("in Convert err, appending: " + fmt.Sprintf("%+v\n", log.Error(fo.Err)))
		lr.Fields = append(lr.Fields, log.Error(fo.Err))
	}
	if len(lr.Fields) > 0 {
		opts.LogRecords = append(opts.LogRecords, lr)
		bplog.Warnw("in Convert, log records is " + fmt.Sprintf("%+v\n", lr))
	}
	return opts
}
