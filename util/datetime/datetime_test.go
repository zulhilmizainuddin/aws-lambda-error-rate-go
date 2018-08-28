package datetime

import (
	"fmt"
	"testing"
	"time"

	"github.com/araddon/dateparse"
)

func TestCalculateMetricTimeRange(t *testing.T) {
	metricTimeRange, err := CalculateMetricTimeRange("2018-08-25T02:07:48.443+0000", 300*time.Second)
	if err != nil {
		t.Error(err.Error())
	}

	expectedStartTime, _ := dateparse.ParseAny("2018-08-25T02:02:49.443+0000")
	if metricTimeRange.Start.Sub(expectedStartTime).String() != "0s" {
		t.Error(fmt.Sprintf("Wrong start time: %v != %v", metricTimeRange.Start.String(), expectedStartTime.String()))
	}

	expectedEndTime, _ := dateparse.ParseAny("2018-08-25T02:07:49.443+0000")
	if metricTimeRange.End.Sub(expectedEndTime).String() != "0s" {
		t.Error(fmt.Sprintf("Wrong end time: %v != %v", metricTimeRange.End.String(), expectedEndTime.String()))
	}
}
