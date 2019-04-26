package opjg

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/goadesign/goa"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func Tracing() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {

			var span opentracing.Span

			spanName := fmt.Sprintf("%v %v", req.Method, req.URL.Path)

			// Try to join to a trace propagated in `req`.
			wireContext, err := tracer.Extract(
				opentracing.TextMap,
				opentracing.HTTPHeadersCarrier(req.Header),
			)

			if err != nil {
				span = tracer.StartSpan(spanName)
			} else {
				span = tracer.StartSpan(spanName, ext.RPCServerOption(wireContext))
			}

			span.SetTag("User-Agent", strings.Join(req.Header["User-Agent"], ", "))
			span.SetTag("http.method", req.Method)
			span.SetTag("http.path", req.URL.Path)

			// create span
			defer span.Finish()

			// store span in context
			ctx = opentracing.ContextWithSpan(ctx, span)

			// update request context to include our new span
			req = req.WithContext(ctx)

			return h(ctx, rw, req)
		}
	}
}
