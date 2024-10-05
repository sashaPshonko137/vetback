package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"strconv"
	"time"
)

var (
	requestMetrics = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: "ads",
		Subsystem: "http",
		Name:      "request_duration_seconds",
		Help:      "Duration of HTTP requests in seconds.",
		Objectives: map[float64]float64{
			0.5:  0.05,
			0.9:  0.01,
			0.99: 0.001,
		},
	}, []string{"status"})

	requestCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "ads",
		Subsystem: "http",
		Name:      "requests_total",
		Help:      "Total number of HTTP requests.",
	}, []string{"status"})

	requestHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "ads",
		Subsystem: "http",
		Name:      "request_duration_histogram_seconds",
		Help:      "Histogram of response latency (seconds) of HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"status"})
)

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Используем ResponseWriter с поддержкой отслеживания статуса
		ww := &responseWriterWithStatus{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(ww, r)

		duration := time.Since(start).Seconds()
		statusCode := ww.statusCode

		statusStr := strconv.Itoa(statusCode)

		// Наблюдаем за временем выполнения через Summary
		requestMetrics.WithLabelValues(statusStr).Observe(duration)

		// Увеличиваем счетчик запросов
		requestCounter.WithLabelValues(statusStr).Inc()

		// Наблюдаем за временем выполнения через Histogram
		requestHistogram.WithLabelValues(statusStr).Observe(duration)
	})
}

// Обёртка для отслеживания статуса ответа
type responseWriterWithStatus struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWithStatus) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
