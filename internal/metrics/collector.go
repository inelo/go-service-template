package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

const (
	namespace = "corridorman"
)

type Collector struct {
	mapMatchSuccess      *prometheus.CounterVec // Count of successfully mapmatched
	mapMatchWithWarnings *prometheus.CounterVec // Count of mapmatched with warnings
	mapMatchFail         *prometheus.CounterVec // Count of failed matchmaps
	mapMatchRetries      *prometheus.CounterVec // Count of connection errors

	mapMatchDuration prometheus.Histogram
}

// NewCollector creates metrics collector with metrics
func NewCollector() *Collector {
	c := &Collector{
		mapMatchSuccess: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "map_match_success",
				Help:      "map matched corridors without errors",
			},
			[]string{},
		),
		mapMatchWithWarnings: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "map_match_with_warnings",
				Help:      "map matched corridors with minor warnings",
			},
			[]string{},
		),
		mapMatchFail: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "map_match_fail",
				Help:      "map matching impossible, used source corridor instead",
			},
			[]string{},
		),
		mapMatchRetries: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Name:      "map_match_retries",
				Help:      "counts times of retries of all map match requests",
			},
			[]string{},
		),
		mapMatchDuration: promauto.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Name:      "map_match_duration",
				Help:      "duration of mapmatching corridor",
				Buckets:   []float64{250, 500, 1000, 5000, 15000, 30000},
			}),
	}

	// Init metrics

	c.MapMatchSuccessAdd(0)
	c.MapMatchWithWarningsAdd(0)
	c.MapMatchRetriesAdd(0)
	c.MapMatchFailAdd(0)

	c.mapMatchDuration.Observe(0)

	// Disable default Go metrics

	prometheus.Unregister(collectors.NewGoCollector())
	prometheus.Unregister(collectors.NewBuildInfoCollector())
	prometheus.Unregister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	return c
}

func (c *Collector) MapMatchSuccessAdd(value float64) {
	c.mapMatchSuccess.With(prometheus.Labels{}).Add(value)
}

func (c *Collector) MapMatchWithWarningsAdd(value float64) {
	c.mapMatchWithWarnings.With(prometheus.Labels{}).Add(value)
}

func (c *Collector) MapMatchFailAdd(value float64) {
	c.mapMatchFail.With(prometheus.Labels{}).Add(value)
}

func (c *Collector) MapMatchRetriesAdd(value float64) {
	c.mapMatchRetries.With(prometheus.Labels{}).Add(value)
}

func (c *Collector) MapMatchDurationObserve(duration time.Duration) {
	c.mapMatchDuration.Observe(float64(duration))
}
