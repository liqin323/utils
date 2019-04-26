package opzk

import (
	"utils/slog"

	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

const (
	// Our service name.
	serviceName = "identity"

	// Host + port of our service.
	hostPort = "localhost"

	// Endpoint to send Zipkin spans to.
	zipkinHTTPEndpoint = "http://localhost:9411/api/v1/spans"

	// Debug mode.
	debug = false

	// same span can be set to true for RPC style spans (Zipkin V1) vs Node style (OpenTracing)
	sameSpan = false

	// make Tracer generate 128 bit traceID's for root spans.
	traceID128Bit = true
)

var tracer opentracing.Tracer

func Initialize() (err error) {
	// create collector.
	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	if err != nil {
		slog.Error("unable to create Zipkin HTTP collector: %+v\n", err)
		return
	}

	// create recorder.
	recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)

	// create tracer.
	t, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(sameSpan),
		zipkin.TraceID128Bit(traceID128Bit),
	)
	if err != nil {
		slog.Error("unable to create Zipkin tracer: %+v\n", err)
		return
	}

	// explicitly set our tracer to be the default tracer.
	opentracing.InitGlobalTracer(t)
	tracer = t

	return
}
