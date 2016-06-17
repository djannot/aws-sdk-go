//Package cloudhsm provides gucumber integration tests support.
package cloudhsm

import (
	"github.com/djannot/aws-sdk-go/awstesting/integration/smoke"
	"github.com/djannot/aws-sdk-go/service/cloudhsm"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cloudhsm", func() {
		World["client"] = cloudhsm.New(smoke.Session)
	})
}
