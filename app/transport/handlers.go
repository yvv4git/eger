package transport

import (
	"github.com/opentracing/opentracing-go"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// curl -X GET 'http://localhost:8080/api/v1/jaeger'
func jaegerHandler(w http.ResponseWriter, r *http.Request) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("long request handler")
	defer span.Finish()

	span.SetTag("http.long request", "some value before long request")

	rand.Seed(time.Now().UnixNano())
	waitSeconds := rand.Intn(5-1) + 1
	longRequest(time.Duration(waitSeconds)*time.Second, span)

	span.SetTag("http.long request", "some value after long request")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("success"))
	if err != nil {
		log.Println("Error: ", err)
		return
	}
}

// longRequest - is very long request
func longRequest(timeWait time.Duration, span opentracing.Span) {
	tracer := opentracing.GlobalTracer()
	longRequestSpan := tracer.StartSpan(
		"long request",
		opentracing.ChildOf(span.Context()),
	)

	time.Sleep(timeWait)

	longRequestSpan.Finish()
}
