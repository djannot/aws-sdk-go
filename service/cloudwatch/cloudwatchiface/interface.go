// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudwatchiface provides an interface for the Amazon CloudWatch.
package cloudwatchiface

import (
	"github.com/djannot/aws-sdk-go/aws/request"
	"github.com/djannot/aws-sdk-go/service/cloudwatch"
)

// CloudWatchAPI is the interface type for cloudwatch.CloudWatch.
type CloudWatchAPI interface {
	DeleteAlarmsRequest(*cloudwatch.DeleteAlarmsInput) (*request.Request, *cloudwatch.DeleteAlarmsOutput)

	DeleteAlarms(*cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error)

	DescribeAlarmHistoryRequest(*cloudwatch.DescribeAlarmHistoryInput) (*request.Request, *cloudwatch.DescribeAlarmHistoryOutput)

	DescribeAlarmHistory(*cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error)

	DescribeAlarmHistoryPages(*cloudwatch.DescribeAlarmHistoryInput, func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool) error

	DescribeAlarmsRequest(*cloudwatch.DescribeAlarmsInput) (*request.Request, *cloudwatch.DescribeAlarmsOutput)

	DescribeAlarms(*cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error)

	DescribeAlarmsPages(*cloudwatch.DescribeAlarmsInput, func(*cloudwatch.DescribeAlarmsOutput, bool) bool) error

	DescribeAlarmsForMetricRequest(*cloudwatch.DescribeAlarmsForMetricInput) (*request.Request, *cloudwatch.DescribeAlarmsForMetricOutput)

	DescribeAlarmsForMetric(*cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error)

	DisableAlarmActionsRequest(*cloudwatch.DisableAlarmActionsInput) (*request.Request, *cloudwatch.DisableAlarmActionsOutput)

	DisableAlarmActions(*cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error)

	EnableAlarmActionsRequest(*cloudwatch.EnableAlarmActionsInput) (*request.Request, *cloudwatch.EnableAlarmActionsOutput)

	EnableAlarmActions(*cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error)

	GetMetricStatisticsRequest(*cloudwatch.GetMetricStatisticsInput) (*request.Request, *cloudwatch.GetMetricStatisticsOutput)

	GetMetricStatistics(*cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error)

	ListMetricsRequest(*cloudwatch.ListMetricsInput) (*request.Request, *cloudwatch.ListMetricsOutput)

	ListMetrics(*cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error)

	ListMetricsPages(*cloudwatch.ListMetricsInput, func(*cloudwatch.ListMetricsOutput, bool) bool) error

	PutMetricAlarmRequest(*cloudwatch.PutMetricAlarmInput) (*request.Request, *cloudwatch.PutMetricAlarmOutput)

	PutMetricAlarm(*cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error)

	PutMetricDataRequest(*cloudwatch.PutMetricDataInput) (*request.Request, *cloudwatch.PutMetricDataOutput)

	PutMetricData(*cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error)

	SetAlarmStateRequest(*cloudwatch.SetAlarmStateInput) (*request.Request, *cloudwatch.SetAlarmStateOutput)

	SetAlarmState(*cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error)
}

var _ CloudWatchAPI = (*cloudwatch.CloudWatch)(nil)
