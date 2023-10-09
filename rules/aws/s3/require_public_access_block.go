package s3

import (
	"fmt"

	"github.com/khulnasoft-lab/tfsecurity/pkg/defsec/infra"
	"github.com/khulnasoft-lab/tfsecurity/pkg/result"
)

func CheckBucketsHavePublicAccessBlocks(context *infra.Context) []*result.Result {
	var results []*result.Result
	for _, bucket := range context.AWS.S3.Buckets {
		if bucket.PublicAccessBlock == nil {
			results = append(results, &result.Result{
				Description: fmt.Sprintf("Bucket '%s' does not have a corresponding public access block.", bucket.Reference),
				Location:    bucket.Range,
			})
		}
	}
	return results
}
