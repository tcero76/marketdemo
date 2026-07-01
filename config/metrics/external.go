package metrics

import "time"

func MeasureExternal(service string) func() {
	start := time.Now()
	return func() {
		ExternalDuration.
			WithLabelValues(service).
			Observe(time.Since(start).Seconds())
	}
}