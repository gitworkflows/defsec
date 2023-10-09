package s3

import (
	"fmt"

	"github.com/khulnasoft-lab/khulnasoft-lab/pkg/defsec/infra"
	"github.com/khulnasoft-lab/khulnasoft-lab/pkg/result"
)

func CheckVersioningIsEnabled(context *infra.Context) []*result.Result {

	var results []*result.Result

	for _, bucket := range context.AWS.S3.Buckets {
		if !bucket.Versioning.Enabled.IsTrue() {
			results = append(results, &result.Result{
				Description: fmt.Sprintf("Resource '%s' does not have versioning enabled", bucket.Reference),
				Location:    bucket.Versioning.Enabled.Range,
			})
		}
	}
	return results
}
