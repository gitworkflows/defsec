package s3

import (
	"fmt"

	"github.com/khulnasoft-lab/defsec/infra"
	"github.com/khulnasoft-lab/defsec/provider"
	"github.com/khulnasoft-lab/defsec/result"
	"github.com/khulnasoft-lab/defsec/rules"
	"github.com/khulnasoft-lab/defsec/severity"
)

var CheckLoggingIsEnabled = rules.RuleDef{

	Provider:  provider.AWSProvider,
	Service:   "s3",
	ShortCode: "enable-bucket-logging",
	Summary:   "S3 Bucket does not have logging enabled.",
	Explanation: `
Buckets should have logging enabled so that access can be audited. 
`,
	Impact:     "There is no way to determine the access to this bucket",
	Resolution: "Add a logging block to the resource to enable access logging",

	Links: []string{
		"https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerLogs.html",
	},

	Severity: severity.Medium,
	CheckFunc: func(context *infra.Context) []*result.Result {
		var results []*result.Result
		for _, bucket := range context.AWS.S3.Buckets {
			if !bucket.Logging.Enabled.IsTrue() && bucket.ACL.NotEqualTo("log-delivery-write") {
				results = append(results, &result.Result{
					Description: fmt.Sprintf("Resource '%s' does not have logging enabled", bucket.Reference),
					Location:    bucket.Logging.Enabled.Range,
				})
			}
		}
		return results
	},
}
