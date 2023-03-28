package log

import "context"

type TraceFunc = func(ctx context.Context) []Field
