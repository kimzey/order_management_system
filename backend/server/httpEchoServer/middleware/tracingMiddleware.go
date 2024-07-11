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

		//if c.Path() == "/metricsx" {
		//	return nil
		//}

		carrier := propagation.HeaderCarrier(c.Request().Header)
		ctx := otel.GetTextMapPropagator().Extract(c.Request().Context(), carrier)

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

		ctx, span := Tracer.Start(ctx, c.Request().Method+" "+c.Path(), spanOptions...)
		defer span.End()

		c.SetRequest(c.Request().WithContext(ctx))

		if err := next(c); err != nil {
			c.Error(err)
		}

		propagator := otel.GetTextMapPropagator()
		carrier = propagation.HeaderCarrier{}
		propagator.Inject(ctx, carrier)

		for _, k := range carrier.Keys() {
			c.Response().Header().Set(k, carrier.Get(k))
		}

		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(c.Response().Status))
		return nil
	}
}
