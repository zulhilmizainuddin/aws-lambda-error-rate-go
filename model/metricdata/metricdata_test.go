package metricdata

import (
	"testing"
	"time"

	"github.com/araddon/dateparse"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/stretchr/testify/assert"
	"github.com/zulhilmizainuddin/aws-lambda-error-rate/util/datetime"
)

func TestBuildErrorRateMetricDataInput(t *testing.T) {
	metricTimeRange, _ := datetime.CalculateMetricTimeRange("2018-08-25T02:07:48.443+0000", 300*time.Second)

	metricDataInput := buildErrorRateMetricDataInput("dummyFunc", metricTimeRange, 60)

	assert.Equal(t, *metricDataInput.MetricDataQueries[0].Id, "errorrate")
	assert.Equal(t, *metricDataInput.MetricDataQueries[0].Expression, "errors / invocations * 100")

	assert.Equal(t, *metricDataInput.MetricDataQueries[1].Id, "errors")
	assert.Equal(t, *metricDataInput.MetricDataQueries[1].MetricStat.Metric.Namespace, "AWS/Lambda")
	assert.Equal(t, *metricDataInput.MetricDataQueries[1].MetricStat.Metric.MetricName, "Errors")
	assert.Equal(t, *metricDataInput.MetricDataQueries[1].MetricStat.Metric.Dimensions[0].Name, "FunctionName")
	assert.Equal(t, *metricDataInput.MetricDataQueries[1].MetricStat.Metric.Dimensions[0].Value, "dummyFunc")
	assert.Equal(t, *metricDataInput.MetricDataQueries[1].MetricStat.Period, int64(60))
	assert.Equal(t, *metricDataInput.MetricDataQueries[1].MetricStat.Stat, "Sum")
	assert.Equal(t, metricDataInput.MetricDataQueries[1].MetricStat.Unit, cloudwatch.StandardUnitCount)

	assert.Equal(t, *metricDataInput.MetricDataQueries[2].Id, "invocations")
	assert.Equal(t, *metricDataInput.MetricDataQueries[2].MetricStat.Metric.Namespace, "AWS/Lambda")
	assert.Equal(t, *metricDataInput.MetricDataQueries[2].MetricStat.Metric.MetricName, "Invocations")
	assert.Equal(t, *metricDataInput.MetricDataQueries[2].MetricStat.Metric.Dimensions[0].Name, "FunctionName")
	assert.Equal(t, *metricDataInput.MetricDataQueries[2].MetricStat.Metric.Dimensions[0].Value, "dummyFunc")
	assert.Equal(t, *metricDataInput.MetricDataQueries[2].MetricStat.Period, int64(60))
	assert.Equal(t, *metricDataInput.MetricDataQueries[2].MetricStat.Stat, "Sum")
	assert.Equal(t, metricDataInput.MetricDataQueries[2].MetricStat.Unit, cloudwatch.StandardUnitCount)

	expectedStartTime, _ := dateparse.ParseAny("2018-08-25T02:02:49.443+0000")
	assert.Equal(t, *metricDataInput.StartTime, expectedStartTime)

	expectedEndTime, _ := dateparse.ParseAny("2018-08-25T02:07:49.443+0000")
	assert.Equal(t, *metricDataInput.EndTime, expectedEndTime)

	assert.Equal(t, metricDataInput.ScanBy, cloudwatch.ScanBy("TimestampAscending"))
}

func TestBuildErrorRateMetricDataQuery(t *testing.T) {
	metricDataQuery := buildErrorRateMetricDataQuery()

	assert.Equal(t, *metricDataQuery.Id, "errorrate")
	assert.Equal(t, *metricDataQuery.Expression, "errors / invocations * 100")
}

func TestBuildErrorsMetricDataQuery(t *testing.T) {
	metricDataQuery := buildErrorsMetricDataQuery("dummyFunc", 60)

	assert.Equal(t, *metricDataQuery.Id, "errors")

	assert.Equal(t, *metricDataQuery.MetricStat.Metric.Namespace, "AWS/Lambda")
	assert.Equal(t, *metricDataQuery.MetricStat.Metric.MetricName, "Errors")
	assert.Equal(t, *metricDataQuery.MetricStat.Metric.Dimensions[0].Name, "FunctionName")
	assert.Equal(t, *metricDataQuery.MetricStat.Metric.Dimensions[0].Value, "dummyFunc")

	assert.Equal(t, *metricDataQuery.MetricStat.Period, int64(60))
	assert.Equal(t, *metricDataQuery.MetricStat.Stat, "Sum")
	assert.Equal(t, metricDataQuery.MetricStat.Unit, cloudwatch.StandardUnitCount)
}

func TestBuildInvocationsMetricDataQuery(t *testing.T) {
	metricDataQuery := buildInvocationsMetricDataQuery("dummyFunc", 60)

	assert.Equal(t, *metricDataQuery.Id, "invocations")

	assert.Equal(t, *metricDataQuery.MetricStat.Metric.Namespace, "AWS/Lambda")
	assert.Equal(t, *metricDataQuery.MetricStat.Metric.MetricName, "Invocations")
	assert.Equal(t, *metricDataQuery.MetricStat.Metric.Dimensions[0].Name, "FunctionName")
	assert.Equal(t, *metricDataQuery.MetricStat.Metric.Dimensions[0].Value, "dummyFunc")

	assert.Equal(t, *metricDataQuery.MetricStat.Period, int64(60))
	assert.Equal(t, *metricDataQuery.MetricStat.Stat, "Sum")
	assert.Equal(t, metricDataQuery.MetricStat.Unit, cloudwatch.StandardUnitCount)
}
