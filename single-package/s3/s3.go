package s3

import "github.com/aws/aws-sdk-go-v2/service/s3"

type ListObjectsAPI struct {
	input s3.ListObjectsV2Input
	output s3.ListObjectsV2Output
}