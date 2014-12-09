// Package sts provides a client for AWS Security Token Service.
package sts

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/bmizerany/aws4"
	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// STS is a client for AWS Security Token Service.
type STS struct {
	client aws.Client
}

// New returns a new STS client.
func New(key, secret, region string, client *http.Client) *STS {
	if client == nil {
		client = http.DefaultClient
	}

	return &STS{
		client: &aws.QueryClient{
			Client: &aws4.Client{
				Keys: &aws4.Keys{
					AccessKey: key,
					SecretKey: secret,
				},
				Client: client,
			},
			Endpoint:   endpoints.Lookup("sts", region),
			APIVersion: "2011-06-15",
		},
	}
}

// AssumeRole returns a set of temporary security credentials (consisting
// of an access key ID, a secret access key, and a security token) that you
// can use to access AWS resources that you might not normally have access
// to. Typically, you use AssumeRole for cross-account access or
// federation. Important: You cannot call AssumeRole by using AWS account
// credentials; access will be denied. You must use IAM user credentials or
// temporary security credentials to call AssumeRole . For cross-account
// access, imagine that you own multiple accounts and need to access
// resources in each account. You could create long-term credentials in
// each account to access those resources. However, managing all those
// credentials and remembering which one can access which account can be
// time consuming. Instead, you can create one set of long-term credentials
// in one account and then use temporary security credentials to access all
// the other accounts by assuming roles in those accounts. For more
// information about roles, see Roles in Using . For federation, you can,
// for example, grant single sign-on access to the AWS Management Console.
// If you already have an identity and authentication system in your
// corporate network, you don't have to recreate user identities in AWS in
// order to grant those user identities access to Instead, after a user has
// been authenticated, you call AssumeRole (and specify the role with the
// appropriate permissions) to get temporary security credentials for that
// user. With those temporary security credentials, you construct a sign-in
// URL that users can use to access the console. For more information, see
// Scenarios for Granting Temporary Access in Using Temporary Security
// Credentials . The temporary security credentials are valid for the
// duration that you specified when calling AssumeRole , which can be from
// 900 seconds (15 minutes) to 3600 seconds (1 hour). The default is 1
// hour. Optionally, you can pass an IAM access policy to this operation.
// If you choose not to pass a policy, the temporary security credentials
// that are returned by the operation have the permissions that are defined
// in the access policy of the role that is being assumed. If you pass a
// policy to this operation, the temporary security credentials that are
// returned by the operation have the permissions that are allowed by both
// the access policy of the role that is being assumed, and the policy that
// you pass. This gives you a way to further restrict the permissions for
// the resulting temporary security credentials. You cannot use the passed
// policy to grant permissions that are in excess of those allowed by the
// access policy of the role that is being assumed. For more information,
// see Permissions for AssumeRole in Using Temporary Security Credentials
// To assume a role, your AWS account must be trusted by the role. The
// trust relationship is defined in the role's trust policy when the role
// is created. You must also have a policy that allows you to call
// sts:AssumeRole . Using MFA with AssumeRole You can optionally include
// multi-factor authentication information when you call AssumeRole . This
// is useful for cross-account scenarios in which you want to make sure
// that the user who is assuming the role has been authenticated using an
// AWS MFA device. In that scenario, the trust policy of the role being
// assumed includes a condition that tests for MFA authentication; if the
// caller does not include valid MFA information, the request to assume the
// role is denied. The condition in a trust policy that tests for MFA
// authentication might look like the following example. "Condition":
// {"Null": {"aws:MultiFactorAuthAge": false}} For more information, see
// Configuring MFA-Protected API Access in the Using guide. To use MFA with
// AssumeRole , you pass values for the SerialNumber and TokenCode
// parameters. The SerialNumber value identifies the user's hardware or
// virtual MFA device. The TokenCode is the time-based one-time password
// that the MFA devices produces.
func (c *STS) AssumeRole(req AssumeRoleRequest) (resp *AssumeRoleResult, err error) {
	resp = &AssumeRoleResult{}
	err = c.client.Do("AssumeRole", "POST", "/", req, resp)
	return
}

// AssumeRoleWithSAML returns a set of temporary security credentials for
// users who have been authenticated via a authentication response. This
// operation provides a mechanism for tying an enterprise identity store or
// directory to role-based AWS access without user-specific credentials or
// configuration. The temporary security credentials returned by this
// operation consist of an access key ID, a secret access key, and a
// security token. Applications can use these temporary security
// credentials to sign calls to AWS services. The credentials are valid for
// the duration that you specified when calling AssumeRoleWithSAML , which
// can be up to 3600 seconds (1 hour) or until the time specified in the
// authentication response's NotOnOrAfter value, whichever is shorter.
// Optionally, you can pass an IAM access policy to this operation. If you
// choose not to pass a policy, the temporary security credentials that are
// returned by the operation have the permissions that are defined in the
// access policy of the role that is being assumed. If you pass a policy to
// this operation, the temporary security credentials that are returned by
// the operation have the permissions that are allowed by both the access
// policy of the role that is being assumed, and the policy that you pass.
// This gives you a way to further restrict the permissions for the
// resulting temporary security credentials. You cannot use the passed
// policy to grant permissions that are in excess of those allowed by the
// access policy of the role that is being assumed. For more information,
// see Permissions for AssumeRoleWithSAML in Using Temporary Security
// Credentials Before your application can call AssumeRoleWithSAML , you
// must configure your identity provider (IdP) to issue the claims required
// by Additionally, you must use AWS Identity and Access Management to
// create a provider entity in your AWS account that represents your
// identity provider, and create an IAM role that specifies this provider
// in its trust policy. Calling AssumeRoleWithSAML does not require the use
// of AWS security credentials. The identity of the caller is validated by
// using keys in the metadata document that is uploaded for the provider
// entity for your identity provider. For more information, see the
// following resources: Creating Temporary Security Credentials for
// Federation in Using Temporary Security Credentials . Providers in Using
// .
func (c *STS) AssumeRoleWithSAML(req AssumeRoleWithSAMLRequest) (resp *AssumeRoleWithSAMLResult, err error) {
	resp = &AssumeRoleWithSAMLResult{}
	err = c.client.Do("AssumeRoleWithSAML", "POST", "/", req, resp)
	return
}

// AssumeRoleWithWebIdentity returns a set of temporary security
// credentials for users who have been authenticated in a mobile or web
// application with a web identity provider, such as Login with Amazon,
// Amazon Cognito, Facebook, or Google. Calling AssumeRoleWithWebIdentity
// does not require the use of AWS security credentials. Therefore, you can
// distribute an application (for example, on mobile devices) that requests
// temporary security credentials without including long-term AWS
// credentials in the application, and without deploying server-based proxy
// services that use long-term AWS credentials. Instead, the identity of
// the caller is validated by using a token from the web identity provider.
// The temporary security credentials returned by this API consist of an
// access key ID, a secret access key, and a security token. Applications
// can use these temporary security credentials to sign calls to AWS
// service APIs. The credentials are valid for the duration that you
// specified when calling AssumeRoleWithWebIdentity , which can be from 900
// seconds (15 minutes) to 3600 seconds (1 hour). By default, the temporary
// security credentials are valid for 1 hour. Optionally, you can pass an
// IAM access policy to this operation. If you choose not to pass a policy,
// the temporary security credentials that are returned by the operation
// have the permissions that are defined in the access policy of the role
// that is being assumed. If you pass a policy to this operation, the
// temporary security credentials that are returned by the operation have
// the permissions that are allowed by both the access policy of the role
// that is being assumed, and the policy that you pass. This gives you a
// way to further restrict the permissions for the resulting temporary
// security credentials. You cannot use the passed policy to grant
// permissions that are in excess of those allowed by the access policy of
// the role that is being assumed. For more information, see Permissions
// for AssumeRoleWithWebIdentity in Using Temporary Security Credentials
// Before your application can call AssumeRoleWithWebIdentity , you must
// have an identity token from a supported identity provider and create a
// role that the application can assume. The role that your application
// assumes must trust the identity provider that is associated with the
// identity token. In other words, the identity provider must be specified
// in the role's trust policy. For more information about how to use web
// identity federation and the AssumeRoleWithWebIdentity , see the
// following resources: Web Identity Federation Playground . This
// interactive website lets you walk through the process of authenticating
// via Login with Amazon, Facebook, or Google, getting temporary security
// credentials, and then using those credentials to make a request to AWS
// SDK for iOS and AWS SDK for Android . These toolkits contain sample apps
// that show how to invoke the identity providers, and then how to use the
// information from these providers to get and use temporary security
// credentials. Web Identity Federation with Mobile Applications . This
// article discusses web identity federation and shows an example of how to
// use web identity federation to get access to content in Amazon S3.
func (c *STS) AssumeRoleWithWebIdentity(req AssumeRoleWithWebIdentityRequest) (resp *AssumeRoleWithWebIdentityResult, err error) {
	resp = &AssumeRoleWithWebIdentityResult{}
	err = c.client.Do("AssumeRoleWithWebIdentity", "POST", "/", req, resp)
	return
}

// DecodeAuthorizationMessage decodes additional information about the
// authorization status of a request from an encoded message returned in
// response to an AWS request. For example, if a user is not authorized to
// perform an action that he or she has requested, the request returns a
// Client.UnauthorizedOperation response (an 403 response). Some AWS
// actions additionally return an encoded message that can provide details
// about this authorization failure. The message is encoded because the
// details of the authorization status can constitute privileged
// information that the user who requested the action should not see. To
// decode an authorization status message, a user must be granted
// permissions via an IAM policy to request the DecodeAuthorizationMessage
// sts:DecodeAuthorizationMessage ) action. The decoded message includes
// the following type of information: Whether the request was denied due to
// an explicit deny or due to the absence of an explicit allow. For more
// information, see Determining Whether a Request is Allowed or Denied in
// Using . The principal who made the request. The requested resource. The
// values of condition keys in the context of the user's request.
func (c *STS) DecodeAuthorizationMessage(req DecodeAuthorizationMessageRequest) (resp *DecodeAuthorizationMessageResult, err error) {
	resp = &DecodeAuthorizationMessageResult{}
	err = c.client.Do("DecodeAuthorizationMessage", "POST", "/", req, resp)
	return
}

// GetFederationToken returns a set of temporary security credentials
// (consisting of an access key ID, a secret access key, and a security
// token) for a federated user. A typical use is in a proxy application
// that gets temporary security credentials on behalf of distributed
// applications inside a corporate network. Because you must call the
// GetFederationToken action using the long-term security credentials of an
// IAM user, this call is appropriate in contexts where those credentials
// can be safely stored, usually in a server-based application. Note: Do
// not use this call in mobile applications or client-based web
// applications that directly get temporary security credentials. For those
// types of applications, use AssumeRoleWithWebIdentity The
// GetFederationToken action must be called by using the long-term AWS
// security credentials of an IAM user. You can also call
// GetFederationToken using the security credentials of an AWS account
// (root), but this is not recommended. Instead, we recommend that you
// create an IAM user for the purpose of the proxy application and then
// attach a policy to the IAM user that limits federated users to only the
// actions and resources they need access to. For more information, see IAM
// Best Practices in Using . The temporary security credentials that are
// obtained by using the long-term credentials of an IAM user are valid for
// the specified duration, between 900 seconds (15 minutes) and 129600
// seconds (36 hours). Temporary credentials that are obtained by using AWS
// account (root) credentials have a maximum duration of 3600 seconds (1
// hour) The permissions for the temporary security credentials returned by
// GetFederationToken are determined by a combination of the following: The
// policy or policies that are attached to the IAM user whose credentials
// are used to call GetFederationToken The policy that is passed as a
// parameter in the call. The passed policy is attached to the temporary
// security credentials that result from the GetFederationToken API
// call--that is, to the federated user . When the federated user makes an
// AWS request, AWS evaluates the policy attached to the federated user in
// combination with the policy or policies attached to the IAM user whose
// credentials were used to call GetFederationToken . AWS allows the
// federated user's request only when both the federated user and the IAM
// user are explicitly allowed to perform the requested action. The passed
// policy cannot grant more permissions than those that are defined in the
// IAM user policy. A typical use case is that the permissions of the IAM
// user whose credentials are used to call GetFederationToken are designed
// to allow access to all the actions and resources that any federated user
// will need. Then, for individual users, you pass a policy to the
// operation that scopes down the permissions to a level that's appropriate
// to that individual user, using a policy that allows only a subset of
// permissions that are granted to the IAM user. If you do not pass a
// policy, the resulting temporary security credentials have no effective
// permissions. The only exception is when the temporary security
// credentials are used to access a resource that has a resource-based
// policy that specifically allows the federated user to access the
// resource. For more information about how permissions work, see
// Permissions for GetFederationToken in Using Temporary Security
// Credentials . For information about using GetFederationToken to create
// temporary security credentials, see Creating Temporary Credentials to
// Enable Access for Federated Users in Using Temporary Security
// Credentials .
func (c *STS) GetFederationToken(req GetFederationTokenRequest) (resp *GetFederationTokenResult, err error) {
	resp = &GetFederationTokenResult{}
	err = c.client.Do("GetFederationToken", "POST", "/", req, resp)
	return
}

// GetSessionToken returns a set of temporary credentials for an AWS
// account or IAM user. The credentials consist of an access key ID, a
// secret access key, and a security token. Typically, you use
// GetSessionToken if you want to use MFA to protect programmatic calls to
// specific AWS APIs like Amazon EC2 StopInstances . MFA-enabled IAM users
// would need to call GetSessionToken and submit an MFA code that is
// associated with their MFA device. Using the temporary security
// credentials that are returned from the call, IAM users can then make
// programmatic calls to APIs that require MFA authentication. The
// GetSessionToken action must be called by using the long-term AWS
// security credentials of the AWS account or an IAM user. Credentials that
// are created by IAM users are valid for the duration that you specify,
// between 900 seconds (15 minutes) and 129600 seconds (36 hours);
// credentials that are created by using account credentials have a maximum
// duration of 3600 seconds (1 hour). The permissions associated with the
// temporary security credentials returned by GetSessionToken are based on
// the permissions associated with account or IAM user whose credentials
// are used to call the action. If GetSessionToken is called using root
// account credentials, the temporary credentials have root account
// permissions. Similarly, if GetSessionToken is called using the
// credentials of an IAM user, the temporary credentials have the same
// permissions as the IAM user. For more information about using
// GetSessionToken to create temporary credentials, go to Creating
// Temporary Credentials to Enable Access for IAM Users in Using Temporary
// Security Credentials .
func (c *STS) GetSessionToken(req GetSessionTokenRequest) (resp *GetSessionTokenResult, err error) {
	resp = &GetSessionTokenResult{}
	err = c.client.Do("GetSessionToken", "POST", "/", req, resp)
	return
}

// AssumeRoleRequest is undocumented.
type AssumeRoleRequest struct {
	DurationSeconds int    `xml:"DurationSeconds"`
	ExternalID      string `xml:"ExternalId"`
	Policy          string `xml:"Policy"`
	RoleARN         string `xml:"RoleArn"`
	RoleSessionName string `xml:"RoleSessionName"`
	SerialNumber    string `xml:"SerialNumber"`
	TokenCode       string `xml:"TokenCode"`
}

// AssumeRoleResponse is undocumented.
type AssumeRoleResponse struct {
	AssumedRoleUser  AssumedRoleUser `xml:"AssumeRoleResult>AssumedRoleUser"`
	Credentials      Credentials     `xml:"AssumeRoleResult>Credentials"`
	PackedPolicySize int             `xml:"AssumeRoleResult>PackedPolicySize"`
}

// AssumeRoleWithSAMLRequest is undocumented.
type AssumeRoleWithSAMLRequest struct {
	DurationSeconds int    `xml:"DurationSeconds"`
	Policy          string `xml:"Policy"`
	PrincipalARN    string `xml:"PrincipalArn"`
	RoleARN         string `xml:"RoleArn"`
	SAMLAssertion   string `xml:"SAMLAssertion"`
}

// AssumeRoleWithSAMLResponse is undocumented.
type AssumeRoleWithSAMLResponse struct {
	AssumedRoleUser  AssumedRoleUser `xml:"AssumeRoleWithSAMLResult>AssumedRoleUser"`
	Audience         string          `xml:"AssumeRoleWithSAMLResult>Audience"`
	Credentials      Credentials     `xml:"AssumeRoleWithSAMLResult>Credentials"`
	Issuer           string          `xml:"AssumeRoleWithSAMLResult>Issuer"`
	NameQualifier    string          `xml:"AssumeRoleWithSAMLResult>NameQualifier"`
	PackedPolicySize int             `xml:"AssumeRoleWithSAMLResult>PackedPolicySize"`
	Subject          string          `xml:"AssumeRoleWithSAMLResult>Subject"`
	SubjectType      string          `xml:"AssumeRoleWithSAMLResult>SubjectType"`
}

// AssumeRoleWithWebIdentityRequest is undocumented.
type AssumeRoleWithWebIdentityRequest struct {
	DurationSeconds  int    `xml:"DurationSeconds"`
	Policy           string `xml:"Policy"`
	ProviderID       string `xml:"ProviderId"`
	RoleARN          string `xml:"RoleArn"`
	RoleSessionName  string `xml:"RoleSessionName"`
	WebIdentityToken string `xml:"WebIdentityToken"`
}

// AssumeRoleWithWebIdentityResponse is undocumented.
type AssumeRoleWithWebIdentityResponse struct {
	AssumedRoleUser             AssumedRoleUser `xml:"AssumeRoleWithWebIdentityResult>AssumedRoleUser"`
	Audience                    string          `xml:"AssumeRoleWithWebIdentityResult>Audience"`
	Credentials                 Credentials     `xml:"AssumeRoleWithWebIdentityResult>Credentials"`
	PackedPolicySize            int             `xml:"AssumeRoleWithWebIdentityResult>PackedPolicySize"`
	Provider                    string          `xml:"AssumeRoleWithWebIdentityResult>Provider"`
	SubjectFromWebIdentityToken string          `xml:"AssumeRoleWithWebIdentityResult>SubjectFromWebIdentityToken"`
}

// AssumedRoleUser is undocumented.
type AssumedRoleUser struct {
	ARN           string `xml:"Arn"`
	AssumedRoleID string `xml:"AssumedRoleId"`
}

// Credentials is undocumented.
type Credentials struct {
	AccessKeyID     string    `xml:"AccessKeyId"`
	Expiration      time.Time `xml:"Expiration"`
	SecretAccessKey string    `xml:"SecretAccessKey"`
	SessionToken    string    `xml:"SessionToken"`
}

// DecodeAuthorizationMessageRequest is undocumented.
type DecodeAuthorizationMessageRequest struct {
	EncodedMessage string `xml:"EncodedMessage"`
}

// DecodeAuthorizationMessageResponse is undocumented.
type DecodeAuthorizationMessageResponse struct {
	DecodedMessage string `xml:"DecodeAuthorizationMessageResult>DecodedMessage"`
}

// FederatedUser is undocumented.
type FederatedUser struct {
	ARN             string `xml:"Arn"`
	FederatedUserID string `xml:"FederatedUserId"`
}

// GetFederationTokenRequest is undocumented.
type GetFederationTokenRequest struct {
	DurationSeconds int    `xml:"DurationSeconds"`
	Name            string `xml:"Name"`
	Policy          string `xml:"Policy"`
}

// GetFederationTokenResponse is undocumented.
type GetFederationTokenResponse struct {
	Credentials      Credentials   `xml:"GetFederationTokenResult>Credentials"`
	FederatedUser    FederatedUser `xml:"GetFederationTokenResult>FederatedUser"`
	PackedPolicySize int           `xml:"GetFederationTokenResult>PackedPolicySize"`
}

// GetSessionTokenRequest is undocumented.
type GetSessionTokenRequest struct {
	DurationSeconds int    `xml:"DurationSeconds"`
	SerialNumber    string `xml:"SerialNumber"`
	TokenCode       string `xml:"TokenCode"`
}

// GetSessionTokenResponse is undocumented.
type GetSessionTokenResponse struct {
	Credentials Credentials `xml:"GetSessionTokenResult>Credentials"`
}

// AssumeRoleResult is a wrapper for AssumeRoleResponse.
type AssumeRoleResult struct {
	XMLName xml.Name `xml:"AssumeRoleResponse"`

	AssumedRoleUser  AssumedRoleUser `xml:"AssumeRoleResult>AssumedRoleUser"`
	Credentials      Credentials     `xml:"AssumeRoleResult>Credentials"`
	PackedPolicySize int             `xml:"AssumeRoleResult>PackedPolicySize"`
}

// AssumeRoleWithSAMLResult is a wrapper for AssumeRoleWithSAMLResponse.
type AssumeRoleWithSAMLResult struct {
	XMLName xml.Name `xml:"AssumeRoleWithSAMLResponse"`

	AssumedRoleUser  AssumedRoleUser `xml:"AssumeRoleWithSAMLResult>AssumedRoleUser"`
	Audience         string          `xml:"AssumeRoleWithSAMLResult>Audience"`
	Credentials      Credentials     `xml:"AssumeRoleWithSAMLResult>Credentials"`
	Issuer           string          `xml:"AssumeRoleWithSAMLResult>Issuer"`
	NameQualifier    string          `xml:"AssumeRoleWithSAMLResult>NameQualifier"`
	PackedPolicySize int             `xml:"AssumeRoleWithSAMLResult>PackedPolicySize"`
	Subject          string          `xml:"AssumeRoleWithSAMLResult>Subject"`
	SubjectType      string          `xml:"AssumeRoleWithSAMLResult>SubjectType"`
}

// AssumeRoleWithWebIdentityResult is a wrapper for AssumeRoleWithWebIdentityResponse.
type AssumeRoleWithWebIdentityResult struct {
	XMLName xml.Name `xml:"AssumeRoleWithWebIdentityResponse"`

	AssumedRoleUser             AssumedRoleUser `xml:"AssumeRoleWithWebIdentityResult>AssumedRoleUser"`
	Audience                    string          `xml:"AssumeRoleWithWebIdentityResult>Audience"`
	Credentials                 Credentials     `xml:"AssumeRoleWithWebIdentityResult>Credentials"`
	PackedPolicySize            int             `xml:"AssumeRoleWithWebIdentityResult>PackedPolicySize"`
	Provider                    string          `xml:"AssumeRoleWithWebIdentityResult>Provider"`
	SubjectFromWebIdentityToken string          `xml:"AssumeRoleWithWebIdentityResult>SubjectFromWebIdentityToken"`
}

// DecodeAuthorizationMessageResult is a wrapper for DecodeAuthorizationMessageResponse.
type DecodeAuthorizationMessageResult struct {
	XMLName xml.Name `xml:"DecodeAuthorizationMessageResponse"`

	DecodedMessage string `xml:"DecodeAuthorizationMessageResult>DecodedMessage"`
}

// GetFederationTokenResult is a wrapper for GetFederationTokenResponse.
type GetFederationTokenResult struct {
	XMLName xml.Name `xml:"GetFederationTokenResponse"`

	Credentials      Credentials   `xml:"GetFederationTokenResult>Credentials"`
	FederatedUser    FederatedUser `xml:"GetFederationTokenResult>FederatedUser"`
	PackedPolicySize int           `xml:"GetFederationTokenResult>PackedPolicySize"`
}

// GetSessionTokenResult is a wrapper for GetSessionTokenResponse.
type GetSessionTokenResult struct {
	XMLName xml.Name `xml:"GetSessionTokenResponse"`

	Credentials Credentials `xml:"GetSessionTokenResult>Credentials"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
