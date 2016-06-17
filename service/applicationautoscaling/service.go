// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package applicationautoscaling

import (
	"github.com/djannot/aws-sdk-go/aws"
	"github.com/djannot/aws-sdk-go/aws/client"
	"github.com/djannot/aws-sdk-go/aws/client/metadata"
	"github.com/djannot/aws-sdk-go/aws/request"
	"github.com/djannot/aws-sdk-go/private/protocol/jsonrpc"
	"github.com/djannot/aws-sdk-go/private/signer/v4"
)

// Application Auto Scaling is a general purpose Auto Scaling service for supported
// elastic AWS resources. With Application Auto Scaling, you can automatically
// scale your AWS resources, with an experience similar to that of Auto Scaling.
//
//  At this time, Application Auto Scaling only supports scaling Amazon ECS
// services.
//
//  For example, you can use Application Auto Scaling to accomplish the following
// tasks:
//
//   Define scaling policies for automatically adjusting your application’s
// resources
//
//   Scale your resources in response to CloudWatch alarms
//
//   View history of your scaling events
//
//   Application Auto Scaling is available in the following regions:
//
//    us-east-1
//
//    us-west-2
//
//    eu-west-1
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type ApplicationAutoScaling struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "autoscaling"

// New creates a new instance of the ApplicationAutoScaling client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ApplicationAutoScaling client from just a session.
//     svc := applicationautoscaling.New(mySession)
//
//     // Create a ApplicationAutoScaling client with additional configuration
//     svc := applicationautoscaling.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ApplicationAutoScaling {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *ApplicationAutoScaling {
	svc := &ApplicationAutoScaling{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   "application-autoscaling",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2016-02-06",
				JSONVersion:   "1.1",
				TargetPrefix:  "AnyScaleFrontendService",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ApplicationAutoScaling operation and runs any
// custom request initialization.
func (c *ApplicationAutoScaling) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
