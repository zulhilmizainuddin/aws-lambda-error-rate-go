package util

import (
	"fmt"
	"testing"

	"github.com/araddon/dateparse"
)

func TestCalculateStateChangePeriod(t *testing.T) {
	stateChangePeriod, err := CalculateStateChangePeriod("2018-08-25T02:07:48.443+0000", 300)
	if err != nil {
		t.Error(err.Error())
	}

	expectedStartTime, _ := dateparse.ParseAny("2018-08-25T02:02:49.443+0000")
	if stateChangePeriod.Start.Sub(expectedStartTime).String() != "0s" {
		t.Error(fmt.Sprintf("Wrong start time: %v != %v", stateChangePeriod.Start.String(), expectedStartTime.String()))
	}

	expectedEndTime, _ := dateparse.ParseAny("2018-08-25T02:07:49.443+0000")
	if stateChangePeriod.End.Sub(expectedEndTime).String() != "0s" {
		t.Error(fmt.Sprintf("Wrong end time: %v != %v", stateChangePeriod.End.String(), expectedEndTime.String()))
	}
}
