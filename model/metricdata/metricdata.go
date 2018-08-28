package metricdata

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/zulhilmizainuddin/aws-lambda-error-rate/util/datetime"
)

type MetricData struct {
	AwsConfig *aws.Config
}

func (this *MetricData) GetMetricDataResults(functionName string, metricTimeRange datetime.MetricTimeRange, periodInSeconds int64) ([]cloudwatch.MetricDataResult, error) {
	var metricDataResults []cloudwatch.MetricDataResult

	metricDataInput := buildErrorRateMetricDataInput(functionName, metricTimeRange, periodInSeconds)

	service := cloudwatch.New(*this.AwsConfig)
	request := service.GetMetricDataRequest(&metricDataInput)

	output, err := request.Send()
	if err != nil {
		return metricDataResults, err
	}

	return output.MetricDataResults, nil
}

func buildErrorRateMetricDataInput(functionName string, metricTimeRange datetime.MetricTimeRange, periodInSeconds int64) cloudwatch.GetMetricDataInput {
	return cloudwatch.GetMetricDataInput{
		MetricDataQueries: []cloudwatch.MetricDataQuery{
			buildErrorRateMetricDataQuery(),
			buildErrorsMetricDataQuery(functionName, periodInSeconds),
			buildInvocationsMetricDataQuery(functionName, periodInSeconds),
		},
		StartTime: &metricTimeRange.Start,
		EndTime:   &metricTimeRange.End,
		ScanBy:    "TimestampAscending",
	}
}

func buildErrorRateMetricDataQuery() cloudwatch.MetricDataQuery {
	errorRateId := "errorrate"
	errorRateExpression := "errors / invocations * 100"

	return cloudwatch.MetricDataQuery{
		Id:         &errorRateId,
		Expression: &errorRateExpression,
	}
}

func buildErrorsMetricDataQuery(functionName string, periodInSeconds int64) cloudwatch.MetricDataQuery {

	dimensionName := "FunctionName"

	dimension := cloudwatch.Dimension{
		Name:  &dimensionName,
		Value: &functionName,
	}

	metricNamespace := "AWS/Lambda"
	metricName := "Errors"

	metric := cloudwatch.Metric{
		Namespace:  &metricNamespace,
		MetricName: &metricName,
		Dimensions: []cloudwatch.Dimension{
			dimension,
		},
	}

	stat := "Sum"

	metricStat := cloudwatch.MetricStat{
		Metric: &metric,
		Period: &periodInSeconds,
		Stat:   &stat,
		Unit:   cloudwatch.StandardUnitCount,
	}

	id := "errors"

	return cloudwatch.MetricDataQuery{
		Id:         &id,
		MetricStat: &metricStat,
	}
}

func buildInvocationsMetricDataQuery(functionName string, periodInSeconds int64) cloudwatch.MetricDataQuery {

	dimensionName := "FunctionName"

	dimension := cloudwatch.Dimension{
		Name:  &dimensionName,
		Value: &functionName,
	}

	metricNamespace := "AWS/Lambda"
	metricName := "Invocations"

	metric := cloudwatch.Metric{
		Namespace:  &metricNamespace,
		MetricName: &metricName,
		Dimensions: []cloudwatch.Dimension{
			dimension,
		},
	}

	stat := "Sum"

	metricStat := cloudwatch.MetricStat{
		Metric: &metric,
		Period: &periodInSeconds,
		Stat:   &stat,
		Unit:   cloudwatch.StandardUnitCount,
	}

	id := "invocations"

	return cloudwatch.MetricDataQuery{
		Id:         &id,
		MetricStat: &metricStat,
	}
}
