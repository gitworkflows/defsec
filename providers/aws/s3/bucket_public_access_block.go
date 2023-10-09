package s3

import "github.com/khulnasoft-lab/defsec/definition"

type PublicAccessBlock struct {
	*definition.Metadata
	Bucket *Bucket
}
