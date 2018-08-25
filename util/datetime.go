package util

import (
	"time"

	"github.com/araddon/dateparse"
)

type StateChangePeriod struct {
	Start time.Time
	End   time.Time
}

func CalculateStateChangePeriod(stateChangeTime string, period time.Duration) (StateChangePeriod, error) {

	var stateChangePeriod StateChangePeriod

	parsedTime, err := dateparse.ParseAny(stateChangeTime)
	if err != nil {
		return stateChangePeriod, err
	}

	const aSecond string = "1s"
	aSecondDuration, err := time.ParseDuration(aSecond)
	if err != nil {
		return stateChangePeriod, err
	}

	endTime := parsedTime.Add(aSecondDuration)
	startTime := endTime.Add(-1 * period * time.Second)

	stateChangePeriod = StateChangePeriod{
		Start: startTime,
		End:   endTime,
	}

	return stateChangePeriod, nil
}
