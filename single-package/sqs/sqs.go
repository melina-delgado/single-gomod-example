package sqs

import "github.com/aws/aws-sdk-go-v2/service/sqs"

type ReceiveMessageAPI struct {
	input sqs.ReceiveMessageInput
	output sqs.ReceiveMessageOutput
}