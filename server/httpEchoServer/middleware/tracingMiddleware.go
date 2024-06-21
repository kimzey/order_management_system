package customMiddleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

var Tracer = otel.GetTracerProvider().Tracer("echo-server")

func TracingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Create a carrier to extract context
		carrier := propagation.HeaderCarrier(c.Request().Header)
		ctx := otel.GetTextMapPropagator().Extract(c.Request().Context(), carrier)

		// Define span options
		spanOptions := []trace.SpanStartOption{
			trace.WithAttributes(semconv.HTTPMethodKey.String(c.Request().Method)),
			trace.WithAttributes(semconv.HTTPTargetKey.String(c.Path())),
			trace.WithAttributes(semconv.HTTPRouteKey.String(c.Path())),
			trace.WithAttributes(semconv.HTTPURLKey.String(fmt.Sprintf("%s://%s%s", c.Scheme(), c.Request().Host, c.Request().RequestURI))),
			trace.WithAttributes(semconv.UserAgentOriginal(c.Request().UserAgent())),
			trace.WithAttributes(semconv.HTTPRequestContentLengthKey.Int64(c.Request().ContentLength)),
			trace.WithAttributes(semconv.HTTPSchemeKey.String(c.Scheme())),
			trace.WithAttributes(semconv.NetTransportTCP),
			trace.WithSpanKind(trace.SpanKindServer),
		}

		// Start a new span with the extracted context and options
		ctx, span := Tracer.Start(ctx, c.Request().Method+" "+c.Path(), spanOptions...)
		defer span.End()

		// Set the context with the new span in the request
		c.SetRequest(c.Request().WithContext(ctx))

		// Debugging - Print the Trace ID in middleware
		traceID := trace.SpanContextFromContext(ctx).TraceID()
		fmt.Println("Middleware Trace ID: ", traceID)

		// Proceed to the next middleware/handler
		if err := next(c); err != nil {
			c.Error(err)
		}

		// Inject headers back into response
		propagator := otel.GetTextMapPropagator()
		carrier = propagation.HeaderCarrier{}
		propagator.Inject(ctx, carrier)

		for _, k := range carrier.Keys() {
			c.Response().Header().Set(k, carrier.Get(k))
		}

		// Set HTTP status code in span attributes
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(c.Response().Status))
		return nil
	}
}

//package customMiddleware
//
//import (
//"fmt"
//"github.com/labstack/echo/v4"
//"go.opentelemetry.io/otel"
//"go.opentelemetry.io/otel/propagation"
//semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
//"go.opentelemetry.io/otel/trace"
//)
//
//var Tracer = otel.GetTracerProvider().Tracer("echo-server")
//
//func TracingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
//		carrier := propagation.HeaderCarrier{}
//
//		// Extract headers from Echo Context and inject into carrier
//		contentType := c.Request().Header.Get(echo.HeaderContentType)
//		carrier.Set(echo.HeaderContentType, contentType)
//
//		// Inject propagated headers into context
//		propagator.Inject(c.Request().Context(), carrier)
//
//		// Define span options
//		spanOptions := []trace.SpanStartOption{
//			trace.WithAttributes(semconv.HTTPMethodKey.String(c.Request().Method)),
//			trace.WithAttributes(semconv.HTTPTargetKey.String(c.Path())),
//			trace.WithAttributes(semconv.HTTPRouteKey.String(c.Path())),
//			trace.WithAttributes(semconv.HTTPURLKey.String(fmt.Sprintf("%s://%s%s", c.Scheme(), c.Request().Host, c.Request().RequestURI))),
//			trace.WithAttributes(semconv.UserAgentOriginal(c.Request().UserAgent())),
//			trace.WithAttributes(semconv.HTTPRequestContentLengthKey.Int64(c.Request().ContentLength)),
//			trace.WithAttributes(semconv.HTTPSchemeKey.String(c.Scheme())),
//			trace.WithAttributes(semconv.NetTransportTCP),
//			trace.WithSpanKind(trace.SpanKindServer),
//		}
//
//		// Start a new span for tracing
//		ctx, span := Tracer.Start(c.Request().Context(), fmt.Sprintf("%s %s", c.Request().Method, c.Path()), spanOptions...)
//		defer span.End()
//
//		// Inject headers back into response
//		{
//			propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
//			carrier := propagation.HeaderCarrier{}
//			propagator.Inject(ctx, carrier)
//
//			for _, k := range carrier.Keys() {
//				c.Response().Header().Set(k, carrier.Get(k))
//			}
//		}
//
//		// Debugging - Print the Trace ID in middleware
//		traceID := trace.SpanContextFromContext(c.Request().Context()).TraceID()
//		fmt.Println("Trace ID: ", traceID)
//
//		// Proceed to the next middleware/handler
//		if err := next(c); err != nil {
//			c.Error(err)
//		}
//
//		// Set HTTP status code in span attributes
//		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(c.Response().Status))
//		return nil
//	}
//}
