package restful

import "github.com/joefitzgerald/standardlog"

// Copyright 2014 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

var trace bool = false
var traceLogger standardlog.Logger

// TraceLogger enables detailed logging of Http request matching and filter invocation. Default no logger is set.
func TraceLogger(logger standardlog.Logger) {
	traceLogger = logger
	trace = logger != nil
}
