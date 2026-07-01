package metrics

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			HTTPInFlight.Inc()
			defer HTTPInFlight.Dec()
			start := time.Now()
			err := next(c)
			duration := time.Since(start).Seconds()
			method := c.Request().Method
			route := c.Path()
			status := strconv.Itoa(c.Response().Status)
			HTTPDuration.WithLabelValues(method, route, status).
				Observe(duration)
			HTTPRequests.WithLabelValues(method, route, status).
				Inc()
			if c.Response().Status >= 500 {
				HTTPErrors.WithLabelValues(method, route, status).
					Inc()
			}
			return err
		}
	}
}