// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package databasemigrationserviceiface provides an interface for the AWS Database Migration Service.
package databasemigrationserviceiface

import (
	"github.com/djannot/aws-sdk-go/aws/request"
	"github.com/djannot/aws-sdk-go/service/databasemigrationservice"
)

// DatabaseMigrationServiceAPI is the interface type for databasemigrationservice.DatabaseMigrationService.
type DatabaseMigrationServiceAPI interface {
	AddTagsToResourceRequest(*databasemigrationservice.AddTagsToResourceInput) (*request.Request, *databasemigrationservice.AddTagsToResourceOutput)

	AddTagsToResource(*databasemigrationservice.AddTagsToResourceInput) (*databasemigrationservice.AddTagsToResourceOutput, error)

	CreateEndpointRequest(*databasemigrationservice.CreateEndpointInput) (*request.Request, *databasemigrationservice.CreateEndpointOutput)

	CreateEndpoint(*databasemigrationservice.CreateEndpointInput) (*databasemigrationservice.CreateEndpointOutput, error)

	CreateReplicationInstanceRequest(*databasemigrationservice.CreateReplicationInstanceInput) (*request.Request, *databasemigrationservice.CreateReplicationInstanceOutput)

	CreateReplicationInstance(*databasemigrationservice.CreateReplicationInstanceInput) (*databasemigrationservice.CreateReplicationInstanceOutput, error)

	CreateReplicationSubnetGroupRequest(*databasemigrationservice.CreateReplicationSubnetGroupInput) (*request.Request, *databasemigrationservice.CreateReplicationSubnetGroupOutput)

	CreateReplicationSubnetGroup(*databasemigrationservice.CreateReplicationSubnetGroupInput) (*databasemigrationservice.CreateReplicationSubnetGroupOutput, error)

	CreateReplicationTaskRequest(*databasemigrationservice.CreateReplicationTaskInput) (*request.Request, *databasemigrationservice.CreateReplicationTaskOutput)

	CreateReplicationTask(*databasemigrationservice.CreateReplicationTaskInput) (*databasemigrationservice.CreateReplicationTaskOutput, error)

	DeleteEndpointRequest(*databasemigrationservice.DeleteEndpointInput) (*request.Request, *databasemigrationservice.DeleteEndpointOutput)

	DeleteEndpoint(*databasemigrationservice.DeleteEndpointInput) (*databasemigrationservice.DeleteEndpointOutput, error)

	DeleteReplicationInstanceRequest(*databasemigrationservice.DeleteReplicationInstanceInput) (*request.Request, *databasemigrationservice.DeleteReplicationInstanceOutput)

	DeleteReplicationInstance(*databasemigrationservice.DeleteReplicationInstanceInput) (*databasemigrationservice.DeleteReplicationInstanceOutput, error)

	DeleteReplicationSubnetGroupRequest(*databasemigrationservice.DeleteReplicationSubnetGroupInput) (*request.Request, *databasemigrationservice.DeleteReplicationSubnetGroupOutput)

	DeleteReplicationSubnetGroup(*databasemigrationservice.DeleteReplicationSubnetGroupInput) (*databasemigrationservice.DeleteReplicationSubnetGroupOutput, error)

	DeleteReplicationTaskRequest(*databasemigrationservice.DeleteReplicationTaskInput) (*request.Request, *databasemigrationservice.DeleteReplicationTaskOutput)

	DeleteReplicationTask(*databasemigrationservice.DeleteReplicationTaskInput) (*databasemigrationservice.DeleteReplicationTaskOutput, error)

	DescribeAccountAttributesRequest(*databasemigrationservice.DescribeAccountAttributesInput) (*request.Request, *databasemigrationservice.DescribeAccountAttributesOutput)

	DescribeAccountAttributes(*databasemigrationservice.DescribeAccountAttributesInput) (*databasemigrationservice.DescribeAccountAttributesOutput, error)

	DescribeConnectionsRequest(*databasemigrationservice.DescribeConnectionsInput) (*request.Request, *databasemigrationservice.DescribeConnectionsOutput)

	DescribeConnections(*databasemigrationservice.DescribeConnectionsInput) (*databasemigrationservice.DescribeConnectionsOutput, error)

	DescribeEndpointTypesRequest(*databasemigrationservice.DescribeEndpointTypesInput) (*request.Request, *databasemigrationservice.DescribeEndpointTypesOutput)

	DescribeEndpointTypes(*databasemigrationservice.DescribeEndpointTypesInput) (*databasemigrationservice.DescribeEndpointTypesOutput, error)

	DescribeEndpointsRequest(*databasemigrationservice.DescribeEndpointsInput) (*request.Request, *databasemigrationservice.DescribeEndpointsOutput)

	DescribeEndpoints(*databasemigrationservice.DescribeEndpointsInput) (*databasemigrationservice.DescribeEndpointsOutput, error)

	DescribeOrderableReplicationInstancesRequest(*databasemigrationservice.DescribeOrderableReplicationInstancesInput) (*request.Request, *databasemigrationservice.DescribeOrderableReplicationInstancesOutput)

	DescribeOrderableReplicationInstances(*databasemigrationservice.DescribeOrderableReplicationInstancesInput) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error)

	DescribeRefreshSchemasStatusRequest(*databasemigrationservice.DescribeRefreshSchemasStatusInput) (*request.Request, *databasemigrationservice.DescribeRefreshSchemasStatusOutput)

	DescribeRefreshSchemasStatus(*databasemigrationservice.DescribeRefreshSchemasStatusInput) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error)

	DescribeReplicationInstancesRequest(*databasemigrationservice.DescribeReplicationInstancesInput) (*request.Request, *databasemigrationservice.DescribeReplicationInstancesOutput)

	DescribeReplicationInstances(*databasemigrationservice.DescribeReplicationInstancesInput) (*databasemigrationservice.DescribeReplicationInstancesOutput, error)

	DescribeReplicationSubnetGroupsRequest(*databasemigrationservice.DescribeReplicationSubnetGroupsInput) (*request.Request, *databasemigrationservice.DescribeReplicationSubnetGroupsOutput)

	DescribeReplicationSubnetGroups(*databasemigrationservice.DescribeReplicationSubnetGroupsInput) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error)

	DescribeReplicationTasksRequest(*databasemigrationservice.DescribeReplicationTasksInput) (*request.Request, *databasemigrationservice.DescribeReplicationTasksOutput)

	DescribeReplicationTasks(*databasemigrationservice.DescribeReplicationTasksInput) (*databasemigrationservice.DescribeReplicationTasksOutput, error)

	DescribeSchemasRequest(*databasemigrationservice.DescribeSchemasInput) (*request.Request, *databasemigrationservice.DescribeSchemasOutput)

	DescribeSchemas(*databasemigrationservice.DescribeSchemasInput) (*databasemigrationservice.DescribeSchemasOutput, error)

	DescribeTableStatisticsRequest(*databasemigrationservice.DescribeTableStatisticsInput) (*request.Request, *databasemigrationservice.DescribeTableStatisticsOutput)

	DescribeTableStatistics(*databasemigrationservice.DescribeTableStatisticsInput) (*databasemigrationservice.DescribeTableStatisticsOutput, error)

	ListTagsForResourceRequest(*databasemigrationservice.ListTagsForResourceInput) (*request.Request, *databasemigrationservice.ListTagsForResourceOutput)

	ListTagsForResource(*databasemigrationservice.ListTagsForResourceInput) (*databasemigrationservice.ListTagsForResourceOutput, error)

	ModifyEndpointRequest(*databasemigrationservice.ModifyEndpointInput) (*request.Request, *databasemigrationservice.ModifyEndpointOutput)

	ModifyEndpoint(*databasemigrationservice.ModifyEndpointInput) (*databasemigrationservice.ModifyEndpointOutput, error)

	ModifyReplicationInstanceRequest(*databasemigrationservice.ModifyReplicationInstanceInput) (*request.Request, *databasemigrationservice.ModifyReplicationInstanceOutput)

	ModifyReplicationInstance(*databasemigrationservice.ModifyReplicationInstanceInput) (*databasemigrationservice.ModifyReplicationInstanceOutput, error)

	ModifyReplicationSubnetGroupRequest(*databasemigrationservice.ModifyReplicationSubnetGroupInput) (*request.Request, *databasemigrationservice.ModifyReplicationSubnetGroupOutput)

	ModifyReplicationSubnetGroup(*databasemigrationservice.ModifyReplicationSubnetGroupInput) (*databasemigrationservice.ModifyReplicationSubnetGroupOutput, error)

	RefreshSchemasRequest(*databasemigrationservice.RefreshSchemasInput) (*request.Request, *databasemigrationservice.RefreshSchemasOutput)

	RefreshSchemas(*databasemigrationservice.RefreshSchemasInput) (*databasemigrationservice.RefreshSchemasOutput, error)

	RemoveTagsFromResourceRequest(*databasemigrationservice.RemoveTagsFromResourceInput) (*request.Request, *databasemigrationservice.RemoveTagsFromResourceOutput)

	RemoveTagsFromResource(*databasemigrationservice.RemoveTagsFromResourceInput) (*databasemigrationservice.RemoveTagsFromResourceOutput, error)

	StartReplicationTaskRequest(*databasemigrationservice.StartReplicationTaskInput) (*request.Request, *databasemigrationservice.StartReplicationTaskOutput)

	StartReplicationTask(*databasemigrationservice.StartReplicationTaskInput) (*databasemigrationservice.StartReplicationTaskOutput, error)

	StopReplicationTaskRequest(*databasemigrationservice.StopReplicationTaskInput) (*request.Request, *databasemigrationservice.StopReplicationTaskOutput)

	StopReplicationTask(*databasemigrationservice.StopReplicationTaskInput) (*databasemigrationservice.StopReplicationTaskOutput, error)

	TestConnectionRequest(*databasemigrationservice.TestConnectionInput) (*request.Request, *databasemigrationservice.TestConnectionOutput)

	TestConnection(*databasemigrationservice.TestConnectionInput) (*databasemigrationservice.TestConnectionOutput, error)
}

var _ DatabaseMigrationServiceAPI = (*databasemigrationservice.DatabaseMigrationService)(nil)
