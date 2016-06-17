// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package gameliftiface provides an interface for the Amazon GameLift.
package gameliftiface

import (
	"github.com/djannot/aws-sdk-go/aws/request"
	"github.com/djannot/aws-sdk-go/service/gamelift"
)

// GameLiftAPI is the interface type for gamelift.GameLift.
type GameLiftAPI interface {
	CreateAliasRequest(*gamelift.CreateAliasInput) (*request.Request, *gamelift.CreateAliasOutput)

	CreateAlias(*gamelift.CreateAliasInput) (*gamelift.CreateAliasOutput, error)

	CreateBuildRequest(*gamelift.CreateBuildInput) (*request.Request, *gamelift.CreateBuildOutput)

	CreateBuild(*gamelift.CreateBuildInput) (*gamelift.CreateBuildOutput, error)

	CreateFleetRequest(*gamelift.CreateFleetInput) (*request.Request, *gamelift.CreateFleetOutput)

	CreateFleet(*gamelift.CreateFleetInput) (*gamelift.CreateFleetOutput, error)

	CreateGameSessionRequest(*gamelift.CreateGameSessionInput) (*request.Request, *gamelift.CreateGameSessionOutput)

	CreateGameSession(*gamelift.CreateGameSessionInput) (*gamelift.CreateGameSessionOutput, error)

	CreatePlayerSessionRequest(*gamelift.CreatePlayerSessionInput) (*request.Request, *gamelift.CreatePlayerSessionOutput)

	CreatePlayerSession(*gamelift.CreatePlayerSessionInput) (*gamelift.CreatePlayerSessionOutput, error)

	CreatePlayerSessionsRequest(*gamelift.CreatePlayerSessionsInput) (*request.Request, *gamelift.CreatePlayerSessionsOutput)

	CreatePlayerSessions(*gamelift.CreatePlayerSessionsInput) (*gamelift.CreatePlayerSessionsOutput, error)

	DeleteAliasRequest(*gamelift.DeleteAliasInput) (*request.Request, *gamelift.DeleteAliasOutput)

	DeleteAlias(*gamelift.DeleteAliasInput) (*gamelift.DeleteAliasOutput, error)

	DeleteBuildRequest(*gamelift.DeleteBuildInput) (*request.Request, *gamelift.DeleteBuildOutput)

	DeleteBuild(*gamelift.DeleteBuildInput) (*gamelift.DeleteBuildOutput, error)

	DeleteFleetRequest(*gamelift.DeleteFleetInput) (*request.Request, *gamelift.DeleteFleetOutput)

	DeleteFleet(*gamelift.DeleteFleetInput) (*gamelift.DeleteFleetOutput, error)

	DeleteScalingPolicyRequest(*gamelift.DeleteScalingPolicyInput) (*request.Request, *gamelift.DeleteScalingPolicyOutput)

	DeleteScalingPolicy(*gamelift.DeleteScalingPolicyInput) (*gamelift.DeleteScalingPolicyOutput, error)

	DescribeAliasRequest(*gamelift.DescribeAliasInput) (*request.Request, *gamelift.DescribeAliasOutput)

	DescribeAlias(*gamelift.DescribeAliasInput) (*gamelift.DescribeAliasOutput, error)

	DescribeBuildRequest(*gamelift.DescribeBuildInput) (*request.Request, *gamelift.DescribeBuildOutput)

	DescribeBuild(*gamelift.DescribeBuildInput) (*gamelift.DescribeBuildOutput, error)

	DescribeEC2InstanceLimitsRequest(*gamelift.DescribeEC2InstanceLimitsInput) (*request.Request, *gamelift.DescribeEC2InstanceLimitsOutput)

	DescribeEC2InstanceLimits(*gamelift.DescribeEC2InstanceLimitsInput) (*gamelift.DescribeEC2InstanceLimitsOutput, error)

	DescribeFleetAttributesRequest(*gamelift.DescribeFleetAttributesInput) (*request.Request, *gamelift.DescribeFleetAttributesOutput)

	DescribeFleetAttributes(*gamelift.DescribeFleetAttributesInput) (*gamelift.DescribeFleetAttributesOutput, error)

	DescribeFleetCapacityRequest(*gamelift.DescribeFleetCapacityInput) (*request.Request, *gamelift.DescribeFleetCapacityOutput)

	DescribeFleetCapacity(*gamelift.DescribeFleetCapacityInput) (*gamelift.DescribeFleetCapacityOutput, error)

	DescribeFleetEventsRequest(*gamelift.DescribeFleetEventsInput) (*request.Request, *gamelift.DescribeFleetEventsOutput)

	DescribeFleetEvents(*gamelift.DescribeFleetEventsInput) (*gamelift.DescribeFleetEventsOutput, error)

	DescribeFleetPortSettingsRequest(*gamelift.DescribeFleetPortSettingsInput) (*request.Request, *gamelift.DescribeFleetPortSettingsOutput)

	DescribeFleetPortSettings(*gamelift.DescribeFleetPortSettingsInput) (*gamelift.DescribeFleetPortSettingsOutput, error)

	DescribeFleetUtilizationRequest(*gamelift.DescribeFleetUtilizationInput) (*request.Request, *gamelift.DescribeFleetUtilizationOutput)

	DescribeFleetUtilization(*gamelift.DescribeFleetUtilizationInput) (*gamelift.DescribeFleetUtilizationOutput, error)

	DescribeGameSessionDetailsRequest(*gamelift.DescribeGameSessionDetailsInput) (*request.Request, *gamelift.DescribeGameSessionDetailsOutput)

	DescribeGameSessionDetails(*gamelift.DescribeGameSessionDetailsInput) (*gamelift.DescribeGameSessionDetailsOutput, error)

	DescribeGameSessionsRequest(*gamelift.DescribeGameSessionsInput) (*request.Request, *gamelift.DescribeGameSessionsOutput)

	DescribeGameSessions(*gamelift.DescribeGameSessionsInput) (*gamelift.DescribeGameSessionsOutput, error)

	DescribePlayerSessionsRequest(*gamelift.DescribePlayerSessionsInput) (*request.Request, *gamelift.DescribePlayerSessionsOutput)

	DescribePlayerSessions(*gamelift.DescribePlayerSessionsInput) (*gamelift.DescribePlayerSessionsOutput, error)

	DescribeScalingPoliciesRequest(*gamelift.DescribeScalingPoliciesInput) (*request.Request, *gamelift.DescribeScalingPoliciesOutput)

	DescribeScalingPolicies(*gamelift.DescribeScalingPoliciesInput) (*gamelift.DescribeScalingPoliciesOutput, error)

	GetGameSessionLogUrlRequest(*gamelift.GetGameSessionLogUrlInput) (*request.Request, *gamelift.GetGameSessionLogUrlOutput)

	GetGameSessionLogUrl(*gamelift.GetGameSessionLogUrlInput) (*gamelift.GetGameSessionLogUrlOutput, error)

	ListAliasesRequest(*gamelift.ListAliasesInput) (*request.Request, *gamelift.ListAliasesOutput)

	ListAliases(*gamelift.ListAliasesInput) (*gamelift.ListAliasesOutput, error)

	ListBuildsRequest(*gamelift.ListBuildsInput) (*request.Request, *gamelift.ListBuildsOutput)

	ListBuilds(*gamelift.ListBuildsInput) (*gamelift.ListBuildsOutput, error)

	ListFleetsRequest(*gamelift.ListFleetsInput) (*request.Request, *gamelift.ListFleetsOutput)

	ListFleets(*gamelift.ListFleetsInput) (*gamelift.ListFleetsOutput, error)

	PutScalingPolicyRequest(*gamelift.PutScalingPolicyInput) (*request.Request, *gamelift.PutScalingPolicyOutput)

	PutScalingPolicy(*gamelift.PutScalingPolicyInput) (*gamelift.PutScalingPolicyOutput, error)

	RequestUploadCredentialsRequest(*gamelift.RequestUploadCredentialsInput) (*request.Request, *gamelift.RequestUploadCredentialsOutput)

	RequestUploadCredentials(*gamelift.RequestUploadCredentialsInput) (*gamelift.RequestUploadCredentialsOutput, error)

	ResolveAliasRequest(*gamelift.ResolveAliasInput) (*request.Request, *gamelift.ResolveAliasOutput)

	ResolveAlias(*gamelift.ResolveAliasInput) (*gamelift.ResolveAliasOutput, error)

	UpdateAliasRequest(*gamelift.UpdateAliasInput) (*request.Request, *gamelift.UpdateAliasOutput)

	UpdateAlias(*gamelift.UpdateAliasInput) (*gamelift.UpdateAliasOutput, error)

	UpdateBuildRequest(*gamelift.UpdateBuildInput) (*request.Request, *gamelift.UpdateBuildOutput)

	UpdateBuild(*gamelift.UpdateBuildInput) (*gamelift.UpdateBuildOutput, error)

	UpdateFleetAttributesRequest(*gamelift.UpdateFleetAttributesInput) (*request.Request, *gamelift.UpdateFleetAttributesOutput)

	UpdateFleetAttributes(*gamelift.UpdateFleetAttributesInput) (*gamelift.UpdateFleetAttributesOutput, error)

	UpdateFleetCapacityRequest(*gamelift.UpdateFleetCapacityInput) (*request.Request, *gamelift.UpdateFleetCapacityOutput)

	UpdateFleetCapacity(*gamelift.UpdateFleetCapacityInput) (*gamelift.UpdateFleetCapacityOutput, error)

	UpdateFleetPortSettingsRequest(*gamelift.UpdateFleetPortSettingsInput) (*request.Request, *gamelift.UpdateFleetPortSettingsOutput)

	UpdateFleetPortSettings(*gamelift.UpdateFleetPortSettingsInput) (*gamelift.UpdateFleetPortSettingsOutput, error)

	UpdateGameSessionRequest(*gamelift.UpdateGameSessionInput) (*request.Request, *gamelift.UpdateGameSessionOutput)

	UpdateGameSession(*gamelift.UpdateGameSessionInput) (*gamelift.UpdateGameSessionOutput, error)
}

var _ GameLiftAPI = (*gamelift.GameLift)(nil)
