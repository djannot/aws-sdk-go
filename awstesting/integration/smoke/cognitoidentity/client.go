//Package cognitoidentity provides gucumber integration tests support.
package cognitoidentity

import (
	"github.com/djannot/aws-sdk-go/awstesting/integration/smoke"
	"github.com/djannot/aws-sdk-go/service/cognitoidentity"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cognitoidentity", func() {
		World["client"] = cognitoidentity.New(smoke.Session)
	})
}
