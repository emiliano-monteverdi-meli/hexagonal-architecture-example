package telemetry

import (
	"time"

	"github.com/mercadolibre/fury_go-core/pkg/telemetry"
	"golang.org/x/net/context"
)

const appPrefix = "application.gateway.apicards.go."

func RecordSimpleMetric(c context.Context, metricName string, tags Tags) {
	localTags := telemetry.Tags(tags.getValues()...)
	telemetry.FromContext(c).Incr(appPrefix+metricName, localTags)
}

func RecordTimeMetric(c context.Context, metricName string, startTime time.Time, tags Tags) {
	localTags := telemetry.Tags(tags.getValues()...)
	elapsedTime := time.Since(startTime).Milliseconds()
	telemetry.TimeInMilliseconds(c, appPrefix+metricName, float64(elapsedTime), localTags)
}
