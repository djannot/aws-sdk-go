// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudtrail_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/djannot/aws-sdk-go/aws"
	"github.com/djannot/aws-sdk-go/aws/session"
	"github.com/djannot/aws-sdk-go/service/cloudtrail"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudTrail_AddTags() {
	svc := cloudtrail.New(session.New())

	params := &cloudtrail.AddTagsInput{
		ResourceId: aws.String("String"), // Required
		TagsList: []*cloudtrail.Tag{
			{ // Required
				Key:   aws.String("String"), // Required
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.AddTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudTrail_CreateTrail() {
	svc := cloudtrail.New(session.New())

	params := &cloudtrail.CreateTrailInput{
		Name:                       aws.String("String"), // Required
		S3BucketName:               aws.String("String"), // Required
		CloudWatchLogsLogGroupArn:  aws.String("String"),
		CloudWatchLogsRoleArn:      aws.String("String"),
		EnableLogFileValidation:    aws.Bool(true),
		IncludeGlobalServiceEvents: aws.Bool(true),
		IsMultiRegionTrail:         aws.Bool(true),
		KmsKeyId:                   aws.String("String"),
		S3KeyPrefix:                aws.String("String"),
		SnsTopicName:               aws.String("String"),
	}
	resp, err := svc.CreateTrail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudTrail_DeleteTrail() {
	svc := cloudtrail.New(session.New())

	params := &cloudtrail.DeleteTrailInput{
		Name: aws.String("String"), // Required
	}
	resp, err := svc.DeleteTrail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudTrail_DescribeTrails() {
	svc := cloudtrail.New(session.New())

	params := &cloudtrail.DescribeTrailsInput{
		IncludeShadowTrails: aws.Bool(true),
		TrailNameList: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTrails(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudTrail_GetTrailStatus() {
	svc := cloudtrail.New(session.New())

	params := &cloudtrail.GetTrailStatusInput{
		Name: aws.String("String"), // Required
	}
	resp, err := svc.GetTrailStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudTrail_ListPublicKeys() {
	svc := cloudtrail.New(session.New())

	params := &cloudtrail.ListPublicKeysInput{
		EndTime:   aws.Time(time.Now()),
		NextToken: aws.String("String"),
		StartTime: aws.Time(time.Now()),
	}
	resp, err := svc.ListPublicKeys(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudTrail_ListTags() {
	svc := cloudtrail.New(session.New())

	params := &cloudtrail.ListTagsInput{
		ResourceIdList: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		NextToken: aws.String("String"),
	}
	resp, err := svc.ListTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudTrail_LookupEvents() {
	svc := cloudtrail.New(session.New())

	params := &cloudtrail.LookupEventsInput{
		EndTime: aws.Time(time.Now()),
		LookupAttributes: []*cloudtrail.LookupAttribute{
			{ // Required
				AttributeKey:   aws.String("LookupAttributeKey"), // Required
				AttributeValue: aws.String("String"),             // Required
			},
			// More values...
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
		StartTime:  aws.Time(time.Now()),
	}
	resp, err := svc.LookupEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudTrail_RemoveTags() {
	svc := cloudtrail.New(session.New())

	params := &cloudtrail.RemoveTagsInput{
		ResourceId: aws.String("String"), // Required
		TagsList: []*cloudtrail.Tag{
			{ // Required
				Key:   aws.String("String"), // Required
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.RemoveTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudTrail_StartLogging() {
	svc := cloudtrail.New(session.New())

	params := &cloudtrail.StartLoggingInput{
		Name: aws.String("String"), // Required
	}
	resp, err := svc.StartLogging(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudTrail_StopLogging() {
	svc := cloudtrail.New(session.New())

	params := &cloudtrail.StopLoggingInput{
		Name: aws.String("String"), // Required
	}
	resp, err := svc.StopLogging(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudTrail_UpdateTrail() {
	svc := cloudtrail.New(session.New())

	params := &cloudtrail.UpdateTrailInput{
		Name: aws.String("String"), // Required
		CloudWatchLogsLogGroupArn:  aws.String("String"),
		CloudWatchLogsRoleArn:      aws.String("String"),
		EnableLogFileValidation:    aws.Bool(true),
		IncludeGlobalServiceEvents: aws.Bool(true),
		IsMultiRegionTrail:         aws.Bool(true),
		KmsKeyId:                   aws.String("String"),
		S3BucketName:               aws.String("String"),
		S3KeyPrefix:                aws.String("String"),
		SnsTopicName:               aws.String("String"),
	}
	resp, err := svc.UpdateTrail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
