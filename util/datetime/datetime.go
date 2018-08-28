package datetime

import (
	"time"

	"github.com/araddon/dateparse"
)

type MetricTimeRange struct {
	Start time.Time
	End   time.Time
}

func CalculateMetricTimeRange(stateChangeTime string, duration time.Duration) (MetricTimeRange, error) {

	var metricTimeRange MetricTimeRange

	parsedTime, err := dateparse.ParseAny(stateChangeTime)
	if err != nil {
		return metricTimeRange, err
	}

	const aSecond string = "1s"
	aSecondDuration, err := time.ParseDuration(aSecond)
	if err != nil {
		return metricTimeRange, err
	}

	endTime := parsedTime.Add(aSecondDuration)
	startTime := endTime.Add(-1 * duration)

	metricTimeRange = MetricTimeRange{
		Start: startTime,
		End:   endTime,
	}

	return metricTimeRange, nil
}
