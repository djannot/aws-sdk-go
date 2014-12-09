// Package autoscaling provides a client for Auto Scaling.
package autoscaling

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/bmizerany/aws4"
	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// AutoScaling is a client for Auto Scaling.
type AutoScaling struct {
	client aws.Client
}

// New returns a new AutoScaling client.
func New(key, secret, region string, client *http.Client) *AutoScaling {
	if client == nil {
		client = http.DefaultClient
	}

	return &AutoScaling{
		client: &aws.QueryClient{
			Client: &aws4.Client{
				Keys: &aws4.Keys{
					AccessKey: key,
					SecretKey: secret,
				},
				Client: client,
			},
			Endpoint:   endpoints.Lookup("autoscaling", region),
			APIVersion: "2011-01-01",
		},
	}
}

// AttachInstances attaches one or more EC2 instances to the specified Auto
// Scaling group. For more information, see Attach Amazon EC2 Instances to
// Your Existing Auto Scaling Group in the Auto Scaling Developer Guide
func (c *AutoScaling) AttachInstances(req AttachInstancesQuery) (err error) {
	// NRE
	err = c.client.Do("AttachInstances", "POST", "/", req, nil)
	return
}

// CompleteLifecycleAction completes the lifecycle action for the
// associated token initiated under the given lifecycle hook with the
// specified result. This operation is a part of the basic sequence for
// adding a lifecycle hook to an Auto Scaling group: Create a notification
// target. A target can be either an Amazon SQS queue or an Amazon SNS
// topic. Create an IAM role. This role allows Auto Scaling to publish
// lifecycle notifications to the designated SQS queue or SNS topic. Create
// the lifecycle hook. You can create a hook that acts when instances
// launch or when instances terminate. If necessary, record the lifecycle
// action heartbeat to keep the instance in a pending state. Complete the
// lifecycle action For more information, see Auto Scaling Pending State
// and Auto Scaling Terminating State in the Auto Scaling Developer Guide
func (c *AutoScaling) CompleteLifecycleAction(req CompleteLifecycleActionType) (resp *CompleteLifecycleActionResult, err error) {
	resp = &CompleteLifecycleActionResult{}
	err = c.client.Do("CompleteLifecycleAction", "POST", "/", req, resp)
	return
}

// CreateAutoScalingGroup creates an Auto Scaling group with the specified
// name and attributes. If you exceed your maximum limit of Auto Scaling
// groups, which by default is 20 per region, the call fails. For
// information about viewing and updating these limits, see
// DescribeAccountLimits
func (c *AutoScaling) CreateAutoScalingGroup(req CreateAutoScalingGroupType) (err error) {
	// NRE
	err = c.client.Do("CreateAutoScalingGroup", "POST", "/", req, nil)
	return
}

// CreateLaunchConfiguration creates a launch configuration. If you exceed
// your maximum limit of launch configurations, which by default is 100 per
// region, the call fails. For information about viewing and updating these
// limits, see DescribeAccountLimits
func (c *AutoScaling) CreateLaunchConfiguration(req CreateLaunchConfigurationType) (err error) {
	// NRE
	err = c.client.Do("CreateLaunchConfiguration", "POST", "/", req, nil)
	return
}

// CreateOrUpdateTags creates or updates tags for the specified Auto
// Scaling group. A tag's definition is composed of a resource ID, resource
// type, key and value, and the propagate flag. Value and the propagate
// flag are optional parameters. See the Request Parameters for more
// information. For more information, see Add, Modify, or Remove Auto
// Scaling Group Tags in the Auto Scaling Developer Guide
func (c *AutoScaling) CreateOrUpdateTags(req CreateOrUpdateTagsType) (err error) {
	// NRE
	err = c.client.Do("CreateOrUpdateTags", "POST", "/", req, nil)
	return
}

// DeleteAutoScalingGroup deletes the specified Auto Scaling group. The
// group must have no instances and no scaling activities in progress. To
// remove all instances before calling DeleteAutoScalingGroup , you can
// call UpdateAutoScalingGroup to set the minimum and maximum size of the
// AutoScalingGroup to zero.
func (c *AutoScaling) DeleteAutoScalingGroup(req DeleteAutoScalingGroupType) (err error) {
	// NRE
	err = c.client.Do("DeleteAutoScalingGroup", "POST", "/", req, nil)
	return
}

// DeleteLaunchConfiguration deletes the specified launch configuration.
// The launch configuration must not be attached to an Auto Scaling group.
// When this call completes, the launch configuration is no longer
// available for use.
func (c *AutoScaling) DeleteLaunchConfiguration(req LaunchConfigurationNameType) (err error) {
	// NRE
	err = c.client.Do("DeleteLaunchConfiguration", "POST", "/", req, nil)
	return
}

// DeleteLifecycleHook deletes the specified lifecycle hook. If there are
// any outstanding lifecycle actions, they are completed first for
// launching instances, for terminating instances).
func (c *AutoScaling) DeleteLifecycleHook(req DeleteLifecycleHookType) (resp *DeleteLifecycleHookResult, err error) {
	resp = &DeleteLifecycleHookResult{}
	err = c.client.Do("DeleteLifecycleHook", "POST", "/", req, resp)
	return
}

// DeleteNotificationConfiguration is undocumented.
func (c *AutoScaling) DeleteNotificationConfiguration(req DeleteNotificationConfigurationType) (err error) {
	// NRE
	err = c.client.Do("DeleteNotificationConfiguration", "POST", "/", req, nil)
	return
}

// DeletePolicy is undocumented.
func (c *AutoScaling) DeletePolicy(req DeletePolicyType) (err error) {
	// NRE
	err = c.client.Do("DeletePolicy", "POST", "/", req, nil)
	return
}

// DeleteScheduledAction is undocumented.
func (c *AutoScaling) DeleteScheduledAction(req DeleteScheduledActionType) (err error) {
	// NRE
	err = c.client.Do("DeleteScheduledAction", "POST", "/", req, nil)
	return
}

// DeleteTags is undocumented.
func (c *AutoScaling) DeleteTags(req DeleteTagsType) (err error) {
	// NRE
	err = c.client.Do("DeleteTags", "POST", "/", req, nil)
	return
}

// DescribeAccountLimits describes the current Auto Scaling resource limits
// for your AWS account. For information about requesting an increase in
// these limits, see AWS Service Limits
func (c *AutoScaling) DescribeAccountLimits() (resp *DescribeAccountLimitsResult, err error) {
	resp = &DescribeAccountLimitsResult{}
	err = c.client.Do("DescribeAccountLimits", "POST", "/", nil, resp)
	return
}

// DescribeAdjustmentTypes lists the policy adjustment types for use with
// PutScalingPolicy
func (c *AutoScaling) DescribeAdjustmentTypes() (resp *DescribeAdjustmentTypesResult, err error) {
	resp = &DescribeAdjustmentTypesResult{}
	err = c.client.Do("DescribeAdjustmentTypes", "POST", "/", nil, resp)
	return
}

// DescribeAutoScalingGroups describes one or more Auto Scaling groups. If
// a list of names is not provided, the call describes all Auto Scaling
// groups. You can specify a maximum number of items to be returned with a
// single call. If there are more items to return, the call returns a
// token. To get the next set of items, repeat the call with the returned
// token in the NextToken parameter.
func (c *AutoScaling) DescribeAutoScalingGroups(req AutoScalingGroupNamesType) (resp *DescribeAutoScalingGroupsResult, err error) {
	resp = &DescribeAutoScalingGroupsResult{}
	err = c.client.Do("DescribeAutoScalingGroups", "POST", "/", req, resp)
	return
}

// DescribeAutoScalingInstances describes one or more Auto Scaling
// instances. If a list is not provided, the call describes all instances.
// You can describe up to a maximum of 50 instances with a single call. By
// default, a call returns up to 20 instances. If there are more items to
// return, the call returns a token. To get the next set of items, repeat
// the call with the returned token in the NextToken parameter.
func (c *AutoScaling) DescribeAutoScalingInstances(req DescribeAutoScalingInstancesType) (resp *DescribeAutoScalingInstancesResult, err error) {
	resp = &DescribeAutoScalingInstancesResult{}
	err = c.client.Do("DescribeAutoScalingInstances", "POST", "/", req, resp)
	return
}

// DescribeAutoScalingNotificationTypes lists the notification types that
// are supported by Auto Scaling.
func (c *AutoScaling) DescribeAutoScalingNotificationTypes() (resp *DescribeAutoScalingNotificationTypesResult, err error) {
	resp = &DescribeAutoScalingNotificationTypesResult{}
	err = c.client.Do("DescribeAutoScalingNotificationTypes", "POST", "/", nil, resp)
	return
}

// DescribeLaunchConfigurations describes one or more launch
// configurations. If you omit the list of names, then the call describes
// all launch configurations. You can specify a maximum number of items to
// be returned with a single call. If there are more items to return, the
// call returns a token. To get the next set of items, repeat the call with
// the returned token in the NextToken parameter.
func (c *AutoScaling) DescribeLaunchConfigurations(req LaunchConfigurationNamesType) (resp *DescribeLaunchConfigurationsResult, err error) {
	resp = &DescribeLaunchConfigurationsResult{}
	err = c.client.Do("DescribeLaunchConfigurations", "POST", "/", req, resp)
	return
}

// DescribeLifecycleHookTypes is undocumented.
func (c *AutoScaling) DescribeLifecycleHookTypes() (resp *DescribeLifecycleHookTypesResult, err error) {
	resp = &DescribeLifecycleHookTypesResult{}
	err = c.client.Do("DescribeLifecycleHookTypes", "POST", "/", nil, resp)
	return
}

// DescribeLifecycleHooks describes the lifecycle hooks for the specified
// Auto Scaling group.
func (c *AutoScaling) DescribeLifecycleHooks(req DescribeLifecycleHooksType) (resp *DescribeLifecycleHooksResult, err error) {
	resp = &DescribeLifecycleHooksResult{}
	err = c.client.Do("DescribeLifecycleHooks", "POST", "/", req, resp)
	return
}

// DescribeMetricCollectionTypes returns a list of metrics and a
// corresponding list of granularities for each metric. The
// GroupStandbyInstances metric is not returned by default. You must
// explicitly request it when calling EnableMetricsCollection
func (c *AutoScaling) DescribeMetricCollectionTypes() (resp *DescribeMetricCollectionTypesResult, err error) {
	resp = &DescribeMetricCollectionTypesResult{}
	err = c.client.Do("DescribeMetricCollectionTypes", "POST", "/", nil, resp)
	return
}

// DescribeNotificationConfigurations describes the notification actions
// associated with the specified Auto Scaling group.
func (c *AutoScaling) DescribeNotificationConfigurations(req DescribeNotificationConfigurationsType) (resp *DescribeNotificationConfigurationsResult, err error) {
	resp = &DescribeNotificationConfigurationsResult{}
	err = c.client.Do("DescribeNotificationConfigurations", "POST", "/", req, resp)
	return
}

// DescribePolicies describes the policies for the specified Auto Scaling
// group. You can specify a maximum number of items to be returned with a
// single call. If there are more items to return, the call returns a
// token. To get the next set of items, repeat the call with the returned
// token in the NextToken parameter.
func (c *AutoScaling) DescribePolicies(req DescribePoliciesType) (resp *DescribePoliciesResult, err error) {
	resp = &DescribePoliciesResult{}
	err = c.client.Do("DescribePolicies", "POST", "/", req, resp)
	return
}

// DescribeScalingActivities describes one or more scaling activities for
// the specified Auto Scaling group. If you omit the ActivityIds , the call
// returns all activities from the past six weeks. Activities are sorted by
// the start time. Activities still in progress appear first on the list.
// You can specify a maximum number of items to be returned with a single
// call. If there are more items to return, the call returns a token. To
// get the next set of items, repeat the call with the returned token in
// the NextToken parameter.
func (c *AutoScaling) DescribeScalingActivities(req DescribeScalingActivitiesType) (resp *DescribeScalingActivitiesResult, err error) {
	resp = &DescribeScalingActivitiesResult{}
	err = c.client.Do("DescribeScalingActivities", "POST", "/", req, resp)
	return
}

// DescribeScalingProcessTypes returns scaling process types for use in the
// ResumeProcesses and SuspendProcesses actions.
func (c *AutoScaling) DescribeScalingProcessTypes() (resp *DescribeScalingProcessTypesResult, err error) {
	resp = &DescribeScalingProcessTypesResult{}
	err = c.client.Do("DescribeScalingProcessTypes", "POST", "/", nil, resp)
	return
}

// DescribeScheduledActions lists the actions scheduled for your Auto
// Scaling group that haven't been executed. To list the actions that were
// already executed, use DescribeScalingActivities
func (c *AutoScaling) DescribeScheduledActions(req DescribeScheduledActionsType) (resp *DescribeScheduledActionsResult, err error) {
	resp = &DescribeScheduledActionsResult{}
	err = c.client.Do("DescribeScheduledActions", "POST", "/", req, resp)
	return
}

// DescribeTags describes the specified tags. You can use filters to limit
// the results. For example, you can query for the tags for a specific Auto
// Scaling group. You can specify multiple values for a filter. A tag must
// match at least one of the specified values for it to be included in the
// results. You can also specify multiple filters. The result includes
// information for a particular tag only if it matches all the filters. If
// there's no match, no special message is returned.
func (c *AutoScaling) DescribeTags(req DescribeTagsType) (resp *DescribeTagsResult, err error) {
	resp = &DescribeTagsResult{}
	err = c.client.Do("DescribeTags", "POST", "/", req, resp)
	return
}

// DescribeTerminationPolicyTypes lists the termination policies supported
// by Auto Scaling.
func (c *AutoScaling) DescribeTerminationPolicyTypes() (resp *DescribeTerminationPolicyTypesResult, err error) {
	resp = &DescribeTerminationPolicyTypesResult{}
	err = c.client.Do("DescribeTerminationPolicyTypes", "POST", "/", nil, resp)
	return
}

// DetachInstances removes one or more instances from the specified Auto
// Scaling group. After the instances are detached, you can manage them
// independently from the rest of the Auto Scaling group. For more
// information, see Detach EC2 Instances from Your Auto Scaling Group in
// the Auto Scaling Developer Guide
func (c *AutoScaling) DetachInstances(req DetachInstancesQuery) (resp *DetachInstancesResult, err error) {
	resp = &DetachInstancesResult{}
	err = c.client.Do("DetachInstances", "POST", "/", req, resp)
	return
}

// DisableMetricsCollection disables monitoring of the specified metrics
// for the specified Auto Scaling group.
func (c *AutoScaling) DisableMetricsCollection(req DisableMetricsCollectionQuery) (err error) {
	// NRE
	err = c.client.Do("DisableMetricsCollection", "POST", "/", req, nil)
	return
}

// EnableMetricsCollection enables monitoring of the specified metrics for
// the specified Auto Scaling group. You can only enable metrics collection
// if InstanceMonitoring in the launch configuration for the group is set
// to True
func (c *AutoScaling) EnableMetricsCollection(req EnableMetricsCollectionQuery) (err error) {
	// NRE
	err = c.client.Do("EnableMetricsCollection", "POST", "/", req, nil)
	return
}

// EnterStandby moves the specified instances into Standby mode. For more
// information, see Auto Scaling InService State in the Auto Scaling
// Developer Guide
func (c *AutoScaling) EnterStandby(req EnterStandbyQuery) (resp *EnterStandbyResult, err error) {
	resp = &EnterStandbyResult{}
	err = c.client.Do("EnterStandby", "POST", "/", req, resp)
	return
}

// ExecutePolicy is undocumented.
func (c *AutoScaling) ExecutePolicy(req ExecutePolicyType) (err error) {
	// NRE
	err = c.client.Do("ExecutePolicy", "POST", "/", req, nil)
	return
}

// ExitStandby moves the specified instances out of Standby mode. For more
// information, see Auto Scaling InService State in the Auto Scaling
// Developer Guide
func (c *AutoScaling) ExitStandby(req ExitStandbyQuery) (resp *ExitStandbyResult, err error) {
	resp = &ExitStandbyResult{}
	err = c.client.Do("ExitStandby", "POST", "/", req, resp)
	return
}

// PutLifecycleHook creates or updates a lifecycle hook for the specified
// Auto Scaling Group. A lifecycle hook tells Auto Scaling that you want to
// perform an action on an instance that is not actively in service; for
// example, either when the instance launches or before the instance
// terminates. This operation is a part of the basic sequence for adding a
// lifecycle hook to an Auto Scaling group: Create a notification target. A
// target can be either an Amazon SQS queue or an Amazon SNS topic. Create
// an IAM role. This role allows Auto Scaling to publish lifecycle
// notifications to the designated SQS queue or SNS topic. Create the
// lifecycle hook. You can create a hook that acts when instances launch or
// when instances terminate. If necessary, record the lifecycle action
// heartbeat to keep the instance in a pending state. Complete the
// lifecycle action. For more information, see Auto Scaling Pending State
// and Auto Scaling Terminating State in the Auto Scaling Developer Guide
func (c *AutoScaling) PutLifecycleHook(req PutLifecycleHookType) (resp *PutLifecycleHookResult, err error) {
	resp = &PutLifecycleHookResult{}
	err = c.client.Do("PutLifecycleHook", "POST", "/", req, resp)
	return
}

// PutNotificationConfiguration configures an Auto Scaling group to send
// notifications when specified events take place. Subscribers to this
// topic can have messages for events delivered to an endpoint such as a
// web server or email address. For more information see Getting
// Notifications When Your Auto Scaling Group Changes in the Auto Scaling
// Developer Guide This configuration overwrites an existing configuration.
func (c *AutoScaling) PutNotificationConfiguration(req PutNotificationConfigurationType) (err error) {
	// NRE
	err = c.client.Do("PutNotificationConfiguration", "POST", "/", req, nil)
	return
}

// PutScalingPolicy creates or updates a policy for an Auto Scaling group.
// To update an existing policy, use the existing policy name and set the
// parameters you want to change. Any existing parameter not changed in an
// update to an existing policy is not changed in this update request.
func (c *AutoScaling) PutScalingPolicy(req PutScalingPolicyType) (resp *PutScalingPolicyResult, err error) {
	resp = &PutScalingPolicyResult{}
	err = c.client.Do("PutScalingPolicy", "POST", "/", req, resp)
	return
}

// PutScheduledUpdateGroupAction creates or updates a scheduled scaling
// action for an Auto Scaling group. When updating a scheduled scaling
// action, if you leave a parameter unspecified, the corresponding value
// remains unchanged in the affected Auto Scaling group. For more
// information, see Scheduled Scaling in the Auto Scaling Developer Guide
// Auto Scaling supports the date and time expressed in
// "YYYY-MM-DDThh:mm:ssZ" format in only.
func (c *AutoScaling) PutScheduledUpdateGroupAction(req PutScheduledUpdateGroupActionType) (err error) {
	// NRE
	err = c.client.Do("PutScheduledUpdateGroupAction", "POST", "/", req, nil)
	return
}

// RecordLifecycleActionHeartbeat records a heartbeat for the lifecycle
// action associated with a specific token. This extends the timeout by the
// length of time defined by the HeartbeatTimeout parameter of
// PutLifecycleHook This operation is a part of the basic sequence for
// adding a lifecycle hook to an Auto Scaling group: Create a notification
// target. A target can be either an Amazon SQS queue or an Amazon SNS
// topic. Create an IAM role. This role allows Auto Scaling to publish
// lifecycle notifications to the designated SQS queue or SNS topic. Create
// the lifecycle hook. You can create a hook that acts when instances
// launch or when instances terminate. If necessary, record the lifecycle
// action heartbeat to keep the instance in a pending state. Complete the
// lifecycle action. For more information, see Auto Scaling Pending State
// and Auto Scaling Terminating State in the Auto Scaling Developer Guide
func (c *AutoScaling) RecordLifecycleActionHeartbeat(req RecordLifecycleActionHeartbeatType) (resp *RecordLifecycleActionHeartbeatResult, err error) {
	resp = &RecordLifecycleActionHeartbeatResult{}
	err = c.client.Do("RecordLifecycleActionHeartbeat", "POST", "/", req, resp)
	return
}

// ResumeProcesses resumes the specified suspended Auto Scaling processes
// for the specified Auto Scaling group. To resume specific processes, use
// the ScalingProcesses parameter. To resume all processes, omit the
// ScalingProcesses parameter. For more information, see Suspend and Resume
// Auto Scaling Processes in the Auto Scaling Developer Guide
func (c *AutoScaling) ResumeProcesses(req ScalingProcessQuery) (err error) {
	// NRE
	err = c.client.Do("ResumeProcesses", "POST", "/", req, nil)
	return
}

// SetDesiredCapacity is undocumented.
func (c *AutoScaling) SetDesiredCapacity(req SetDesiredCapacityType) (err error) {
	// NRE
	err = c.client.Do("SetDesiredCapacity", "POST", "/", req, nil)
	return
}

// SetInstanceHealth sets the health status of the specified instance. For
// more information, see Health Checks in the Auto Scaling Developer Guide
func (c *AutoScaling) SetInstanceHealth(req SetInstanceHealthQuery) (err error) {
	// NRE
	err = c.client.Do("SetInstanceHealth", "POST", "/", req, nil)
	return
}

// SuspendProcesses suspends the specified Auto Scaling processes for the
// specified Auto Scaling group. To suspend specific processes, use the
// ScalingProcesses parameter. To suspend all processes, omit the
// ScalingProcesses parameter. Note that if you suspend either the Launch
// or Terminate process types, it can prevent other process types from
// functioning properly. To resume processes that have been suspended, use
// ResumeProcesses For more information, see Suspend and Resume Auto
// Scaling Processes in the Auto Scaling Developer Guide
func (c *AutoScaling) SuspendProcesses(req ScalingProcessQuery) (err error) {
	// NRE
	err = c.client.Do("SuspendProcesses", "POST", "/", req, nil)
	return
}

// TerminateInstanceInAutoScalingGroup terminates the specified instance
// and optionally adjusts the desired group size. This call simply makes a
// termination request. The instances is not terminated immediately.
func (c *AutoScaling) TerminateInstanceInAutoScalingGroup(req TerminateInstanceInAutoScalingGroupType) (resp *TerminateInstanceInAutoScalingGroupResult, err error) {
	resp = &TerminateInstanceInAutoScalingGroupResult{}
	err = c.client.Do("TerminateInstanceInAutoScalingGroup", "POST", "/", req, resp)
	return
}

// UpdateAutoScalingGroup updates the configuration for the specified
// AutoScalingGroup . To update an Auto Scaling group with a launch
// configuration that has the InstanceMonitoring flag set to False , you
// must first ensure that collection of group metrics is disabled.
// Otherwise, calls to UpdateAutoScalingGroup will fail. If you have
// previously enabled group metrics collection, you can disable collection
// of all group metrics by calling DisableMetricsCollection . The new
// settings are registered upon the completion of this call. Any launch
// configuration settings take effect on any triggers after this call
// returns. Scaling activities that are currently in progress aren't
// affected. If a new value is specified for MinSize without specifying the
// value for DesiredCapacity , and if the new MinSize is larger than the
// current size of the Auto Scaling group, there will be an implicit call
// to SetDesiredCapacity to set the group to the new MinSize . If a new
// value is specified for MaxSize without specifying the value for
// DesiredCapacity , and the new MaxSize is smaller than the current size
// of the Auto Scaling group, there will be an implicit call to
// SetDesiredCapacity to set the group to the new MaxSize . All other
// optional parameters are left unchanged if not passed in the request.
func (c *AutoScaling) UpdateAutoScalingGroup(req UpdateAutoScalingGroupType) (err error) {
	// NRE
	err = c.client.Do("UpdateAutoScalingGroup", "POST", "/", req, nil)
	return
}

// ActivitiesType is undocumented.
type ActivitiesType struct {
	Activities []Activity `xml:"DescribeScalingActivitiesResult>Activities>member"`
	NextToken  string     `xml:"DescribeScalingActivitiesResult>NextToken"`
}

// Activity is undocumented.
type Activity struct {
	ActivityID           string    `xml:"ActivityId"`
	AutoScalingGroupName string    `xml:"AutoScalingGroupName"`
	Cause                string    `xml:"Cause"`
	Description          string    `xml:"Description"`
	Details              string    `xml:"Details"`
	EndTime              time.Time `xml:"EndTime"`
	Progress             int       `xml:"Progress"`
	StartTime            time.Time `xml:"StartTime"`
	StatusCode           string    `xml:"StatusCode"`
	StatusMessage        string    `xml:"StatusMessage"`
}

// ActivityType is undocumented.
type ActivityType struct {
	Activity Activity `xml:"TerminateInstanceInAutoScalingGroupResult>Activity"`
}

// AdjustmentType is undocumented.
type AdjustmentType struct {
	AdjustmentType string `xml:"AdjustmentType"`
}

// Alarm is undocumented.
type Alarm struct {
	AlarmARN  string `xml:"AlarmARN"`
	AlarmName string `xml:"AlarmName"`
}

// AttachInstancesQuery is undocumented.
type AttachInstancesQuery struct {
	AutoScalingGroupName string   `xml:"AutoScalingGroupName"`
	InstanceIds          []string `xml:"InstanceIds>member"`
}

// AutoScalingGroup is undocumented.
type AutoScalingGroup struct {
	AutoScalingGroupARN     string             `xml:"AutoScalingGroupARN"`
	AutoScalingGroupName    string             `xml:"AutoScalingGroupName"`
	AvailabilityZones       []string           `xml:"AvailabilityZones>member"`
	CreatedTime             time.Time          `xml:"CreatedTime"`
	DefaultCooldown         int                `xml:"DefaultCooldown"`
	DesiredCapacity         int                `xml:"DesiredCapacity"`
	EnabledMetrics          []EnabledMetric    `xml:"EnabledMetrics>member"`
	HealthCheckGracePeriod  int                `xml:"HealthCheckGracePeriod"`
	HealthCheckType         string             `xml:"HealthCheckType"`
	Instances               []Instance         `xml:"Instances>member"`
	LaunchConfigurationName string             `xml:"LaunchConfigurationName"`
	LoadBalancerNames       []string           `xml:"LoadBalancerNames>member"`
	MaxSize                 int                `xml:"MaxSize"`
	MinSize                 int                `xml:"MinSize"`
	PlacementGroup          string             `xml:"PlacementGroup"`
	Status                  string             `xml:"Status"`
	SuspendedProcesses      []SuspendedProcess `xml:"SuspendedProcesses>member"`
	Tags                    []TagDescription   `xml:"Tags>member"`
	TerminationPolicies     []string           `xml:"TerminationPolicies>member"`
	VPCZoneIdentifier       string             `xml:"VPCZoneIdentifier"`
}

// AutoScalingGroupNamesType is undocumented.
type AutoScalingGroupNamesType struct {
	AutoScalingGroupNames []string `xml:"AutoScalingGroupNames>member"`
	MaxRecords            int      `xml:"MaxRecords"`
	NextToken             string   `xml:"NextToken"`
}

// AutoScalingGroupsType is undocumented.
type AutoScalingGroupsType struct {
	AutoScalingGroups []AutoScalingGroup `xml:"DescribeAutoScalingGroupsResult>AutoScalingGroups>member"`
	NextToken         string             `xml:"DescribeAutoScalingGroupsResult>NextToken"`
}

// AutoScalingInstanceDetails is undocumented.
type AutoScalingInstanceDetails struct {
	AutoScalingGroupName    string `xml:"AutoScalingGroupName"`
	AvailabilityZone        string `xml:"AvailabilityZone"`
	HealthStatus            string `xml:"HealthStatus"`
	InstanceID              string `xml:"InstanceId"`
	LaunchConfigurationName string `xml:"LaunchConfigurationName"`
	LifecycleState          string `xml:"LifecycleState"`
}

// AutoScalingInstancesType is undocumented.
type AutoScalingInstancesType struct {
	AutoScalingInstances []AutoScalingInstanceDetails `xml:"DescribeAutoScalingInstancesResult>AutoScalingInstances>member"`
	NextToken            string                       `xml:"DescribeAutoScalingInstancesResult>NextToken"`
}

// BlockDeviceMapping is undocumented.
type BlockDeviceMapping struct {
	DeviceName  string `xml:"DeviceName"`
	Ebs         Ebs    `xml:"Ebs"`
	NoDevice    bool   `xml:"NoDevice"`
	VirtualName string `xml:"VirtualName"`
}

// CompleteLifecycleActionAnswer is undocumented.
type CompleteLifecycleActionAnswer struct {
}

// CompleteLifecycleActionType is undocumented.
type CompleteLifecycleActionType struct {
	AutoScalingGroupName  string `xml:"AutoScalingGroupName"`
	LifecycleActionResult string `xml:"LifecycleActionResult"`
	LifecycleActionToken  string `xml:"LifecycleActionToken"`
	LifecycleHookName     string `xml:"LifecycleHookName"`
}

// CreateAutoScalingGroupType is undocumented.
type CreateAutoScalingGroupType struct {
	AutoScalingGroupName    string   `xml:"AutoScalingGroupName"`
	AvailabilityZones       []string `xml:"AvailabilityZones>member"`
	DefaultCooldown         int      `xml:"DefaultCooldown"`
	DesiredCapacity         int      `xml:"DesiredCapacity"`
	HealthCheckGracePeriod  int      `xml:"HealthCheckGracePeriod"`
	HealthCheckType         string   `xml:"HealthCheckType"`
	InstanceID              string   `xml:"InstanceId"`
	LaunchConfigurationName string   `xml:"LaunchConfigurationName"`
	LoadBalancerNames       []string `xml:"LoadBalancerNames>member"`
	MaxSize                 int      `xml:"MaxSize"`
	MinSize                 int      `xml:"MinSize"`
	PlacementGroup          string   `xml:"PlacementGroup"`
	Tags                    []Tag    `xml:"Tags>member"`
	TerminationPolicies     []string `xml:"TerminationPolicies>member"`
	VPCZoneIdentifier       string   `xml:"VPCZoneIdentifier"`
}

// CreateLaunchConfigurationType is undocumented.
type CreateLaunchConfigurationType struct {
	AssociatePublicIPAddress bool                 `xml:"AssociatePublicIpAddress"`
	BlockDeviceMappings      []BlockDeviceMapping `xml:"BlockDeviceMappings>member"`
	EbsOptimized             bool                 `xml:"EbsOptimized"`
	IamInstanceProfile       string               `xml:"IamInstanceProfile"`
	ImageID                  string               `xml:"ImageId"`
	InstanceID               string               `xml:"InstanceId"`
	InstanceMonitoring       InstanceMonitoring   `xml:"InstanceMonitoring"`
	InstanceType             string               `xml:"InstanceType"`
	KernelID                 string               `xml:"KernelId"`
	KeyName                  string               `xml:"KeyName"`
	LaunchConfigurationName  string               `xml:"LaunchConfigurationName"`
	PlacementTenancy         string               `xml:"PlacementTenancy"`
	RamdiskID                string               `xml:"RamdiskId"`
	SecurityGroups           []string             `xml:"SecurityGroups>member"`
	SpotPrice                string               `xml:"SpotPrice"`
	UserData                 string               `xml:"UserData"`
}

// CreateOrUpdateTagsType is undocumented.
type CreateOrUpdateTagsType struct {
	Tags []Tag `xml:"Tags>member"`
}

// DeleteAutoScalingGroupType is undocumented.
type DeleteAutoScalingGroupType struct {
	AutoScalingGroupName string `xml:"AutoScalingGroupName"`
	ForceDelete          bool   `xml:"ForceDelete"`
}

// DeleteLifecycleHookAnswer is undocumented.
type DeleteLifecycleHookAnswer struct {
}

// DeleteLifecycleHookType is undocumented.
type DeleteLifecycleHookType struct {
	AutoScalingGroupName string `xml:"AutoScalingGroupName"`
	LifecycleHookName    string `xml:"LifecycleHookName"`
}

// DeleteNotificationConfigurationType is undocumented.
type DeleteNotificationConfigurationType struct {
	AutoScalingGroupName string `xml:"AutoScalingGroupName"`
	TopicARN             string `xml:"TopicARN"`
}

// DeletePolicyType is undocumented.
type DeletePolicyType struct {
	AutoScalingGroupName string `xml:"AutoScalingGroupName"`
	PolicyName           string `xml:"PolicyName"`
}

// DeleteScheduledActionType is undocumented.
type DeleteScheduledActionType struct {
	AutoScalingGroupName string `xml:"AutoScalingGroupName"`
	ScheduledActionName  string `xml:"ScheduledActionName"`
}

// DeleteTagsType is undocumented.
type DeleteTagsType struct {
	Tags []Tag `xml:"Tags>member"`
}

// DescribeAccountLimitsAnswer is undocumented.
type DescribeAccountLimitsAnswer struct {
	MaxNumberOfAutoScalingGroups    int `xml:"DescribeAccountLimitsResult>MaxNumberOfAutoScalingGroups"`
	MaxNumberOfLaunchConfigurations int `xml:"DescribeAccountLimitsResult>MaxNumberOfLaunchConfigurations"`
}

// DescribeAdjustmentTypesAnswer is undocumented.
type DescribeAdjustmentTypesAnswer struct {
	AdjustmentTypes []AdjustmentType `xml:"DescribeAdjustmentTypesResult>AdjustmentTypes>member"`
}

// DescribeAutoScalingInstancesType is undocumented.
type DescribeAutoScalingInstancesType struct {
	InstanceIds []string `xml:"InstanceIds>member"`
	MaxRecords  int      `xml:"MaxRecords"`
	NextToken   string   `xml:"NextToken"`
}

// DescribeAutoScalingNotificationTypesAnswer is undocumented.
type DescribeAutoScalingNotificationTypesAnswer struct {
	AutoScalingNotificationTypes []string `xml:"DescribeAutoScalingNotificationTypesResult>AutoScalingNotificationTypes>member"`
}

// DescribeLifecycleHookTypesAnswer is undocumented.
type DescribeLifecycleHookTypesAnswer struct {
	LifecycleHookTypes []string `xml:"DescribeLifecycleHookTypesResult>LifecycleHookTypes>member"`
}

// DescribeLifecycleHooksAnswer is undocumented.
type DescribeLifecycleHooksAnswer struct {
	LifecycleHooks []LifecycleHook `xml:"DescribeLifecycleHooksResult>LifecycleHooks>member"`
}

// DescribeLifecycleHooksType is undocumented.
type DescribeLifecycleHooksType struct {
	AutoScalingGroupName string   `xml:"AutoScalingGroupName"`
	LifecycleHookNames   []string `xml:"LifecycleHookNames>member"`
}

// DescribeMetricCollectionTypesAnswer is undocumented.
type DescribeMetricCollectionTypesAnswer struct {
	Granularities []MetricGranularityType `xml:"DescribeMetricCollectionTypesResult>Granularities>member"`
	Metrics       []MetricCollectionType  `xml:"DescribeMetricCollectionTypesResult>Metrics>member"`
}

// DescribeNotificationConfigurationsAnswer is undocumented.
type DescribeNotificationConfigurationsAnswer struct {
	NextToken                  string                      `xml:"DescribeNotificationConfigurationsResult>NextToken"`
	NotificationConfigurations []NotificationConfiguration `xml:"DescribeNotificationConfigurationsResult>NotificationConfigurations>member"`
}

// DescribeNotificationConfigurationsType is undocumented.
type DescribeNotificationConfigurationsType struct {
	AutoScalingGroupNames []string `xml:"AutoScalingGroupNames>member"`
	MaxRecords            int      `xml:"MaxRecords"`
	NextToken             string   `xml:"NextToken"`
}

// DescribePoliciesType is undocumented.
type DescribePoliciesType struct {
	AutoScalingGroupName string   `xml:"AutoScalingGroupName"`
	MaxRecords           int      `xml:"MaxRecords"`
	NextToken            string   `xml:"NextToken"`
	PolicyNames          []string `xml:"PolicyNames>member"`
}

// DescribeScalingActivitiesType is undocumented.
type DescribeScalingActivitiesType struct {
	ActivityIds          []string `xml:"ActivityIds>member"`
	AutoScalingGroupName string   `xml:"AutoScalingGroupName"`
	MaxRecords           int      `xml:"MaxRecords"`
	NextToken            string   `xml:"NextToken"`
}

// DescribeScheduledActionsType is undocumented.
type DescribeScheduledActionsType struct {
	AutoScalingGroupName string    `xml:"AutoScalingGroupName"`
	EndTime              time.Time `xml:"EndTime"`
	MaxRecords           int       `xml:"MaxRecords"`
	NextToken            string    `xml:"NextToken"`
	ScheduledActionNames []string  `xml:"ScheduledActionNames>member"`
	StartTime            time.Time `xml:"StartTime"`
}

// DescribeTagsType is undocumented.
type DescribeTagsType struct {
	Filters    []Filter `xml:"Filters>member"`
	MaxRecords int      `xml:"MaxRecords"`
	NextToken  string   `xml:"NextToken"`
}

// DescribeTerminationPolicyTypesAnswer is undocumented.
type DescribeTerminationPolicyTypesAnswer struct {
	TerminationPolicyTypes []string `xml:"DescribeTerminationPolicyTypesResult>TerminationPolicyTypes>member"`
}

// DetachInstancesAnswer is undocumented.
type DetachInstancesAnswer struct {
	Activities []Activity `xml:"DetachInstancesResult>Activities>member"`
}

// DetachInstancesQuery is undocumented.
type DetachInstancesQuery struct {
	AutoScalingGroupName           string   `xml:"AutoScalingGroupName"`
	InstanceIds                    []string `xml:"InstanceIds>member"`
	ShouldDecrementDesiredCapacity bool     `xml:"ShouldDecrementDesiredCapacity"`
}

// DisableMetricsCollectionQuery is undocumented.
type DisableMetricsCollectionQuery struct {
	AutoScalingGroupName string   `xml:"AutoScalingGroupName"`
	Metrics              []string `xml:"Metrics>member"`
}

// Ebs is undocumented.
type Ebs struct {
	DeleteOnTermination bool   `xml:"DeleteOnTermination"`
	Iops                int    `xml:"Iops"`
	SnapshotID          string `xml:"SnapshotId"`
	VolumeSize          int    `xml:"VolumeSize"`
	VolumeType          string `xml:"VolumeType"`
}

// EnableMetricsCollectionQuery is undocumented.
type EnableMetricsCollectionQuery struct {
	AutoScalingGroupName string   `xml:"AutoScalingGroupName"`
	Granularity          string   `xml:"Granularity"`
	Metrics              []string `xml:"Metrics>member"`
}

// EnabledMetric is undocumented.
type EnabledMetric struct {
	Granularity string `xml:"Granularity"`
	Metric      string `xml:"Metric"`
}

// EnterStandbyAnswer is undocumented.
type EnterStandbyAnswer struct {
	Activities []Activity `xml:"EnterStandbyResult>Activities>member"`
}

// EnterStandbyQuery is undocumented.
type EnterStandbyQuery struct {
	AutoScalingGroupName           string   `xml:"AutoScalingGroupName"`
	InstanceIds                    []string `xml:"InstanceIds>member"`
	ShouldDecrementDesiredCapacity bool     `xml:"ShouldDecrementDesiredCapacity"`
}

// ExecutePolicyType is undocumented.
type ExecutePolicyType struct {
	AutoScalingGroupName string `xml:"AutoScalingGroupName"`
	HonorCooldown        bool   `xml:"HonorCooldown"`
	PolicyName           string `xml:"PolicyName"`
}

// ExitStandbyAnswer is undocumented.
type ExitStandbyAnswer struct {
	Activities []Activity `xml:"ExitStandbyResult>Activities>member"`
}

// ExitStandbyQuery is undocumented.
type ExitStandbyQuery struct {
	AutoScalingGroupName string   `xml:"AutoScalingGroupName"`
	InstanceIds          []string `xml:"InstanceIds>member"`
}

// Filter is undocumented.
type Filter struct {
	Name   string   `xml:"Name"`
	Values []string `xml:"Values>member"`
}

// Instance is undocumented.
type Instance struct {
	AvailabilityZone        string `xml:"AvailabilityZone"`
	HealthStatus            string `xml:"HealthStatus"`
	InstanceID              string `xml:"InstanceId"`
	LaunchConfigurationName string `xml:"LaunchConfigurationName"`
	LifecycleState          string `xml:"LifecycleState"`
}

// InstanceMonitoring is undocumented.
type InstanceMonitoring struct {
	Enabled bool `xml:"Enabled"`
}

// LaunchConfiguration is undocumented.
type LaunchConfiguration struct {
	AssociatePublicIPAddress bool                 `xml:"AssociatePublicIpAddress"`
	BlockDeviceMappings      []BlockDeviceMapping `xml:"BlockDeviceMappings>member"`
	CreatedTime              time.Time            `xml:"CreatedTime"`
	EbsOptimized             bool                 `xml:"EbsOptimized"`
	IamInstanceProfile       string               `xml:"IamInstanceProfile"`
	ImageID                  string               `xml:"ImageId"`
	InstanceMonitoring       InstanceMonitoring   `xml:"InstanceMonitoring"`
	InstanceType             string               `xml:"InstanceType"`
	KernelID                 string               `xml:"KernelId"`
	KeyName                  string               `xml:"KeyName"`
	LaunchConfigurationARN   string               `xml:"LaunchConfigurationARN"`
	LaunchConfigurationName  string               `xml:"LaunchConfigurationName"`
	PlacementTenancy         string               `xml:"PlacementTenancy"`
	RamdiskID                string               `xml:"RamdiskId"`
	SecurityGroups           []string             `xml:"SecurityGroups>member"`
	SpotPrice                string               `xml:"SpotPrice"`
	UserData                 string               `xml:"UserData"`
}

// LaunchConfigurationNameType is undocumented.
type LaunchConfigurationNameType struct {
	LaunchConfigurationName string `xml:"LaunchConfigurationName"`
}

// LaunchConfigurationNamesType is undocumented.
type LaunchConfigurationNamesType struct {
	LaunchConfigurationNames []string `xml:"LaunchConfigurationNames>member"`
	MaxRecords               int      `xml:"MaxRecords"`
	NextToken                string   `xml:"NextToken"`
}

// LaunchConfigurationsType is undocumented.
type LaunchConfigurationsType struct {
	LaunchConfigurations []LaunchConfiguration `xml:"DescribeLaunchConfigurationsResult>LaunchConfigurations>member"`
	NextToken            string                `xml:"DescribeLaunchConfigurationsResult>NextToken"`
}

// LifecycleHook is undocumented.
type LifecycleHook struct {
	AutoScalingGroupName  string `xml:"AutoScalingGroupName"`
	DefaultResult         string `xml:"DefaultResult"`
	GlobalTimeout         int    `xml:"GlobalTimeout"`
	HeartbeatTimeout      int    `xml:"HeartbeatTimeout"`
	LifecycleHookName     string `xml:"LifecycleHookName"`
	LifecycleTransition   string `xml:"LifecycleTransition"`
	NotificationMetadata  string `xml:"NotificationMetadata"`
	NotificationTargetARN string `xml:"NotificationTargetARN"`
	RoleARN               string `xml:"RoleARN"`
}

// MetricCollectionType is undocumented.
type MetricCollectionType struct {
	Metric string `xml:"Metric"`
}

// MetricGranularityType is undocumented.
type MetricGranularityType struct {
	Granularity string `xml:"Granularity"`
}

// NotificationConfiguration is undocumented.
type NotificationConfiguration struct {
	AutoScalingGroupName string `xml:"AutoScalingGroupName"`
	NotificationType     string `xml:"NotificationType"`
	TopicARN             string `xml:"TopicARN"`
}

// PoliciesType is undocumented.
type PoliciesType struct {
	NextToken       string          `xml:"DescribePoliciesResult>NextToken"`
	ScalingPolicies []ScalingPolicy `xml:"DescribePoliciesResult>ScalingPolicies>member"`
}

// PolicyARNType is undocumented.
type PolicyARNType struct {
	PolicyARN string `xml:"PutScalingPolicyResult>PolicyARN"`
}

// ProcessType is undocumented.
type ProcessType struct {
	ProcessName string `xml:"ProcessName"`
}

// ProcessesType is undocumented.
type ProcessesType struct {
	Processes []ProcessType `xml:"DescribeScalingProcessTypesResult>Processes>member"`
}

// PutLifecycleHookAnswer is undocumented.
type PutLifecycleHookAnswer struct {
}

// PutLifecycleHookType is undocumented.
type PutLifecycleHookType struct {
	AutoScalingGroupName  string `xml:"AutoScalingGroupName"`
	DefaultResult         string `xml:"DefaultResult"`
	HeartbeatTimeout      int    `xml:"HeartbeatTimeout"`
	LifecycleHookName     string `xml:"LifecycleHookName"`
	LifecycleTransition   string `xml:"LifecycleTransition"`
	NotificationMetadata  string `xml:"NotificationMetadata"`
	NotificationTargetARN string `xml:"NotificationTargetARN"`
	RoleARN               string `xml:"RoleARN"`
}

// PutNotificationConfigurationType is undocumented.
type PutNotificationConfigurationType struct {
	AutoScalingGroupName string   `xml:"AutoScalingGroupName"`
	NotificationTypes    []string `xml:"NotificationTypes>member"`
	TopicARN             string   `xml:"TopicARN"`
}

// PutScalingPolicyType is undocumented.
type PutScalingPolicyType struct {
	AdjustmentType       string `xml:"AdjustmentType"`
	AutoScalingGroupName string `xml:"AutoScalingGroupName"`
	Cooldown             int    `xml:"Cooldown"`
	MinAdjustmentStep    int    `xml:"MinAdjustmentStep"`
	PolicyName           string `xml:"PolicyName"`
	ScalingAdjustment    int    `xml:"ScalingAdjustment"`
}

// PutScheduledUpdateGroupActionType is undocumented.
type PutScheduledUpdateGroupActionType struct {
	AutoScalingGroupName string    `xml:"AutoScalingGroupName"`
	DesiredCapacity      int       `xml:"DesiredCapacity"`
	EndTime              time.Time `xml:"EndTime"`
	MaxSize              int       `xml:"MaxSize"`
	MinSize              int       `xml:"MinSize"`
	Recurrence           string    `xml:"Recurrence"`
	ScheduledActionName  string    `xml:"ScheduledActionName"`
	StartTime            time.Time `xml:"StartTime"`
	Time                 time.Time `xml:"Time"`
}

// RecordLifecycleActionHeartbeatAnswer is undocumented.
type RecordLifecycleActionHeartbeatAnswer struct {
}

// RecordLifecycleActionHeartbeatType is undocumented.
type RecordLifecycleActionHeartbeatType struct {
	AutoScalingGroupName string `xml:"AutoScalingGroupName"`
	LifecycleActionToken string `xml:"LifecycleActionToken"`
	LifecycleHookName    string `xml:"LifecycleHookName"`
}

// ScalingPolicy is undocumented.
type ScalingPolicy struct {
	AdjustmentType       string  `xml:"AdjustmentType"`
	Alarms               []Alarm `xml:"Alarms>member"`
	AutoScalingGroupName string  `xml:"AutoScalingGroupName"`
	Cooldown             int     `xml:"Cooldown"`
	MinAdjustmentStep    int     `xml:"MinAdjustmentStep"`
	PolicyARN            string  `xml:"PolicyARN"`
	PolicyName           string  `xml:"PolicyName"`
	ScalingAdjustment    int     `xml:"ScalingAdjustment"`
}

// ScalingProcessQuery is undocumented.
type ScalingProcessQuery struct {
	AutoScalingGroupName string   `xml:"AutoScalingGroupName"`
	ScalingProcesses     []string `xml:"ScalingProcesses>member"`
}

// ScheduledActionsType is undocumented.
type ScheduledActionsType struct {
	NextToken                   string                       `xml:"DescribeScheduledActionsResult>NextToken"`
	ScheduledUpdateGroupActions []ScheduledUpdateGroupAction `xml:"DescribeScheduledActionsResult>ScheduledUpdateGroupActions>member"`
}

// ScheduledUpdateGroupAction is undocumented.
type ScheduledUpdateGroupAction struct {
	AutoScalingGroupName string    `xml:"AutoScalingGroupName"`
	DesiredCapacity      int       `xml:"DesiredCapacity"`
	EndTime              time.Time `xml:"EndTime"`
	MaxSize              int       `xml:"MaxSize"`
	MinSize              int       `xml:"MinSize"`
	Recurrence           string    `xml:"Recurrence"`
	ScheduledActionARN   string    `xml:"ScheduledActionARN"`
	ScheduledActionName  string    `xml:"ScheduledActionName"`
	StartTime            time.Time `xml:"StartTime"`
	Time                 time.Time `xml:"Time"`
}

// SetDesiredCapacityType is undocumented.
type SetDesiredCapacityType struct {
	AutoScalingGroupName string `xml:"AutoScalingGroupName"`
	DesiredCapacity      int    `xml:"DesiredCapacity"`
	HonorCooldown        bool   `xml:"HonorCooldown"`
}

// SetInstanceHealthQuery is undocumented.
type SetInstanceHealthQuery struct {
	HealthStatus             string `xml:"HealthStatus"`
	InstanceID               string `xml:"InstanceId"`
	ShouldRespectGracePeriod bool   `xml:"ShouldRespectGracePeriod"`
}

// SuspendedProcess is undocumented.
type SuspendedProcess struct {
	ProcessName      string `xml:"ProcessName"`
	SuspensionReason string `xml:"SuspensionReason"`
}

// Tag is undocumented.
type Tag struct {
	Key               string `xml:"Key"`
	PropagateAtLaunch bool   `xml:"PropagateAtLaunch"`
	ResourceID        string `xml:"ResourceId"`
	ResourceType      string `xml:"ResourceType"`
	Value             string `xml:"Value"`
}

// TagDescription is undocumented.
type TagDescription struct {
	Key               string `xml:"Key"`
	PropagateAtLaunch bool   `xml:"PropagateAtLaunch"`
	ResourceID        string `xml:"ResourceId"`
	ResourceType      string `xml:"ResourceType"`
	Value             string `xml:"Value"`
}

// TagsType is undocumented.
type TagsType struct {
	NextToken string           `xml:"DescribeTagsResult>NextToken"`
	Tags      []TagDescription `xml:"DescribeTagsResult>Tags>member"`
}

// TerminateInstanceInAutoScalingGroupType is undocumented.
type TerminateInstanceInAutoScalingGroupType struct {
	InstanceID                     string `xml:"InstanceId"`
	ShouldDecrementDesiredCapacity bool   `xml:"ShouldDecrementDesiredCapacity"`
}

// UpdateAutoScalingGroupType is undocumented.
type UpdateAutoScalingGroupType struct {
	AutoScalingGroupName    string   `xml:"AutoScalingGroupName"`
	AvailabilityZones       []string `xml:"AvailabilityZones>member"`
	DefaultCooldown         int      `xml:"DefaultCooldown"`
	DesiredCapacity         int      `xml:"DesiredCapacity"`
	HealthCheckGracePeriod  int      `xml:"HealthCheckGracePeriod"`
	HealthCheckType         string   `xml:"HealthCheckType"`
	LaunchConfigurationName string   `xml:"LaunchConfigurationName"`
	MaxSize                 int      `xml:"MaxSize"`
	MinSize                 int      `xml:"MinSize"`
	PlacementGroup          string   `xml:"PlacementGroup"`
	TerminationPolicies     []string `xml:"TerminationPolicies>member"`
	VPCZoneIdentifier       string   `xml:"VPCZoneIdentifier"`
}

// CompleteLifecycleActionResult is a wrapper for CompleteLifecycleActionAnswer.
type CompleteLifecycleActionResult struct {
	XMLName xml.Name `xml:"CompleteLifecycleActionResponse"`
}

// DeleteLifecycleHookResult is a wrapper for DeleteLifecycleHookAnswer.
type DeleteLifecycleHookResult struct {
	XMLName xml.Name `xml:"DeleteLifecycleHookResponse"`
}

// DescribeAccountLimitsResult is a wrapper for DescribeAccountLimitsAnswer.
type DescribeAccountLimitsResult struct {
	XMLName xml.Name `xml:"DescribeAccountLimitsResponse"`

	MaxNumberOfAutoScalingGroups    int `xml:"DescribeAccountLimitsResult>MaxNumberOfAutoScalingGroups"`
	MaxNumberOfLaunchConfigurations int `xml:"DescribeAccountLimitsResult>MaxNumberOfLaunchConfigurations"`
}

// DescribeAdjustmentTypesResult is a wrapper for DescribeAdjustmentTypesAnswer.
type DescribeAdjustmentTypesResult struct {
	XMLName xml.Name `xml:"DescribeAdjustmentTypesResponse"`

	AdjustmentTypes []AdjustmentType `xml:"DescribeAdjustmentTypesResult>AdjustmentTypes>member"`
}

// DescribeAutoScalingGroupsResult is a wrapper for AutoScalingGroupsType.
type DescribeAutoScalingGroupsResult struct {
	XMLName xml.Name `xml:"DescribeAutoScalingGroupsResponse"`

	AutoScalingGroups []AutoScalingGroup `xml:"DescribeAutoScalingGroupsResult>AutoScalingGroups>member"`
	NextToken         string             `xml:"DescribeAutoScalingGroupsResult>NextToken"`
}

// DescribeAutoScalingInstancesResult is a wrapper for AutoScalingInstancesType.
type DescribeAutoScalingInstancesResult struct {
	XMLName xml.Name `xml:"DescribeAutoScalingInstancesResponse"`

	AutoScalingInstances []AutoScalingInstanceDetails `xml:"DescribeAutoScalingInstancesResult>AutoScalingInstances>member"`
	NextToken            string                       `xml:"DescribeAutoScalingInstancesResult>NextToken"`
}

// DescribeAutoScalingNotificationTypesResult is a wrapper for DescribeAutoScalingNotificationTypesAnswer.
type DescribeAutoScalingNotificationTypesResult struct {
	XMLName xml.Name `xml:"DescribeAutoScalingNotificationTypesResponse"`

	AutoScalingNotificationTypes []string `xml:"DescribeAutoScalingNotificationTypesResult>AutoScalingNotificationTypes>member"`
}

// DescribeLaunchConfigurationsResult is a wrapper for LaunchConfigurationsType.
type DescribeLaunchConfigurationsResult struct {
	XMLName xml.Name `xml:"DescribeLaunchConfigurationsResponse"`

	LaunchConfigurations []LaunchConfiguration `xml:"DescribeLaunchConfigurationsResult>LaunchConfigurations>member"`
	NextToken            string                `xml:"DescribeLaunchConfigurationsResult>NextToken"`
}

// DescribeLifecycleHookTypesResult is a wrapper for DescribeLifecycleHookTypesAnswer.
type DescribeLifecycleHookTypesResult struct {
	XMLName xml.Name `xml:"DescribeLifecycleHookTypesResponse"`

	LifecycleHookTypes []string `xml:"DescribeLifecycleHookTypesResult>LifecycleHookTypes>member"`
}

// DescribeLifecycleHooksResult is a wrapper for DescribeLifecycleHooksAnswer.
type DescribeLifecycleHooksResult struct {
	XMLName xml.Name `xml:"DescribeLifecycleHooksResponse"`

	LifecycleHooks []LifecycleHook `xml:"DescribeLifecycleHooksResult>LifecycleHooks>member"`
}

// DescribeMetricCollectionTypesResult is a wrapper for DescribeMetricCollectionTypesAnswer.
type DescribeMetricCollectionTypesResult struct {
	XMLName xml.Name `xml:"DescribeMetricCollectionTypesResponse"`

	Granularities []MetricGranularityType `xml:"DescribeMetricCollectionTypesResult>Granularities>member"`
	Metrics       []MetricCollectionType  `xml:"DescribeMetricCollectionTypesResult>Metrics>member"`
}

// DescribeNotificationConfigurationsResult is a wrapper for DescribeNotificationConfigurationsAnswer.
type DescribeNotificationConfigurationsResult struct {
	XMLName xml.Name `xml:"DescribeNotificationConfigurationsResponse"`

	NextToken                  string                      `xml:"DescribeNotificationConfigurationsResult>NextToken"`
	NotificationConfigurations []NotificationConfiguration `xml:"DescribeNotificationConfigurationsResult>NotificationConfigurations>member"`
}

// DescribePoliciesResult is a wrapper for PoliciesType.
type DescribePoliciesResult struct {
	XMLName xml.Name `xml:"DescribePoliciesResponse"`

	NextToken       string          `xml:"DescribePoliciesResult>NextToken"`
	ScalingPolicies []ScalingPolicy `xml:"DescribePoliciesResult>ScalingPolicies>member"`
}

// DescribeScalingActivitiesResult is a wrapper for ActivitiesType.
type DescribeScalingActivitiesResult struct {
	XMLName xml.Name `xml:"DescribeScalingActivitiesResponse"`

	Activities []Activity `xml:"DescribeScalingActivitiesResult>Activities>member"`
	NextToken  string     `xml:"DescribeScalingActivitiesResult>NextToken"`
}

// DescribeScalingProcessTypesResult is a wrapper for ProcessesType.
type DescribeScalingProcessTypesResult struct {
	XMLName xml.Name `xml:"DescribeScalingProcessTypesResponse"`

	Processes []ProcessType `xml:"DescribeScalingProcessTypesResult>Processes>member"`
}

// DescribeScheduledActionsResult is a wrapper for ScheduledActionsType.
type DescribeScheduledActionsResult struct {
	XMLName xml.Name `xml:"DescribeScheduledActionsResponse"`

	NextToken                   string                       `xml:"DescribeScheduledActionsResult>NextToken"`
	ScheduledUpdateGroupActions []ScheduledUpdateGroupAction `xml:"DescribeScheduledActionsResult>ScheduledUpdateGroupActions>member"`
}

// DescribeTagsResult is a wrapper for TagsType.
type DescribeTagsResult struct {
	XMLName xml.Name `xml:"DescribeTagsResponse"`

	NextToken string           `xml:"DescribeTagsResult>NextToken"`
	Tags      []TagDescription `xml:"DescribeTagsResult>Tags>member"`
}

// DescribeTerminationPolicyTypesResult is a wrapper for DescribeTerminationPolicyTypesAnswer.
type DescribeTerminationPolicyTypesResult struct {
	XMLName xml.Name `xml:"DescribeTerminationPolicyTypesResponse"`

	TerminationPolicyTypes []string `xml:"DescribeTerminationPolicyTypesResult>TerminationPolicyTypes>member"`
}

// DetachInstancesResult is a wrapper for DetachInstancesAnswer.
type DetachInstancesResult struct {
	XMLName xml.Name `xml:"DetachInstancesResponse"`

	Activities []Activity `xml:"DetachInstancesResult>Activities>member"`
}

// EnterStandbyResult is a wrapper for EnterStandbyAnswer.
type EnterStandbyResult struct {
	XMLName xml.Name `xml:"EnterStandbyResponse"`

	Activities []Activity `xml:"EnterStandbyResult>Activities>member"`
}

// ExitStandbyResult is a wrapper for ExitStandbyAnswer.
type ExitStandbyResult struct {
	XMLName xml.Name `xml:"ExitStandbyResponse"`

	Activities []Activity `xml:"ExitStandbyResult>Activities>member"`
}

// PutLifecycleHookResult is a wrapper for PutLifecycleHookAnswer.
type PutLifecycleHookResult struct {
	XMLName xml.Name `xml:"PutLifecycleHookResponse"`
}

// PutScalingPolicyResult is a wrapper for PolicyARNType.
type PutScalingPolicyResult struct {
	XMLName xml.Name `xml:"PutScalingPolicyResponse"`

	PolicyARN string `xml:"PutScalingPolicyResult>PolicyARN"`
}

// RecordLifecycleActionHeartbeatResult is a wrapper for RecordLifecycleActionHeartbeatAnswer.
type RecordLifecycleActionHeartbeatResult struct {
	XMLName xml.Name `xml:"RecordLifecycleActionHeartbeatResponse"`
}

// TerminateInstanceInAutoScalingGroupResult is a wrapper for ActivityType.
type TerminateInstanceInAutoScalingGroupResult struct {
	XMLName xml.Name `xml:"TerminateInstanceInAutoScalingGroupResponse"`

	Activity Activity `xml:"TerminateInstanceInAutoScalingGroupResult>Activity"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
