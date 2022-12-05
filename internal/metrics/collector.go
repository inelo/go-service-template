package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

const (
	namespace = "go-service-template"
)

type Collector struct {
}

// NewCollector creates metrics collector with metrics
func NewCollector() *Collector {
	c := &Collector{}

	// Init metrics

	// Disable default Go metrics

	prometheus.Unregister(collectors.NewGoCollector())
	prometheus.Unregister(collectors.NewBuildInfoCollector())
	prometheus.Unregister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	return c
}
