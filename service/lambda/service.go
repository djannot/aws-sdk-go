// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package lambda

import (
	"github.com/djannot/aws-sdk-go/aws"
	"github.com/djannot/aws-sdk-go/aws/client"
	"github.com/djannot/aws-sdk-go/aws/client/metadata"
	"github.com/djannot/aws-sdk-go/aws/request"
	"github.com/djannot/aws-sdk-go/private/protocol/restjson"
	"github.com/djannot/aws-sdk-go/private/signer/v4"
)

// Overview
//
// This is the AWS Lambda API Reference. The AWS Lambda Developer Guide provides
// additional information. For the service overview, go to What is AWS Lambda
// (http://docs.aws.amazon.com/lambda/latest/dg/welcome.html), and for information
// about how the service works, go to AWS Lambda: How it Works (http://docs.aws.amazon.com/lambda/latest/dg/lambda-introduction.html)
// in the AWS Lambda Developer Guide.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type Lambda struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "lambda"

// New creates a new instance of the Lambda client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Lambda client from just a session.
//     svc := lambda.New(mySession)
//
//     // Create a Lambda client with additional configuration
//     svc := lambda.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Lambda {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *Lambda {
	svc := &Lambda{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-03-31",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Lambda operation and runs any
// custom request initialization.
func (c *Lambda) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
