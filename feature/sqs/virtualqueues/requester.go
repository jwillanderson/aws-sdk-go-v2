package virtualqueues

import (
	"context"
	"crypto/rand"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	smithyrand "github.com/aws/smithy-go/rand"
	"github.com/google/uuid"
)

var ErrTimeout = errors.New("unable to receive message response, timeout reached.")

type IdempotencyTokenProvider interface {
	GetIdempotencyToken() (string, error)
}

type QueueRequester struct {
	sqsClient       *sqs.Client
	queuePrefix     string
	queueAttributes map[string]string
	attributeNames  []types.QueueAttributeName

	queueTokenProvider IdempotencyTokenProvider

	responsePollInterval time.Duration
}

func NewDefaultQueueRequester(sqsClient *sqs.Client, queuePrefix string) (*QueueRequester, error) {
	if sqsClient == nil {
		return nil, errors.New("unable to create queue requester, missing sqs client")
	}

	if queuePrefix == "" {
		return nil, errors.New("unable to create queue request with empty queue prefix")
	}

	return &QueueRequester{
		sqsClient:            sqs.Client,
		queuePrefix:          queuePrefix,
		queueAttributes:      map[string]string{},
		attributeNames:       []types.QueueAttributeName{types.QueueAttributeNameAll},
		responsePollInterval: time.Second * 1,
		queueTokenProvider:   smithyrand.NewUUIDIdempotencyToken(rand.Reader),
	}, nil
}

func (qr *QueueRequester) WithQueueAttributes(attrs map[string]string) *QueueRequester {
	qr.queueAttributes = attrs
	return qr
}

func (qr *QueueRequester) WithAttributeName(attrNames []types.QueueAttributeName) *QueueRequester {
	qr.attributeNames = attrNames
	return qr
}

func (qr *QueueRequester) WithResponsePollInterval(interval time.Duration) *QueueRequester {
	qr.responsePollInterval = interval
	return qr
}

func (qr *QueueRequester) SendMessageAndGetResponse(ctx context.Context, input *sqs.SendMessageInput, timeout time.Duration) (*types.Message, error) {
	vQueueId, err := qr.queueTokenProvider.GetIdempotencyToken()
	if err != nil {
		return nil, err
	}
	virtQueueName := qr.queuePrefix + "_" + vQueueId
	createQueueRequest := &sqs.CreateQueueInput{
		QueueName:  &virtQueueName,
		Attributes: qr.queueAttributes,
	}

	out, err := qr.sqsClient.CreateQueue(ctx, createQueueRequest)
	if err != nil {
		return nil, err
	}

	input.MessageAttributes[RESPONSE_QUEUE_URL_ATTRIBUTE_NAME] = types.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: out.QueueUrl,
	}

	defer func() {
		_, _ = qr.sqsClient.DeleteQueue(ctx, &sqs.DeleteQueueInput{QueueUrl: out.QueueUrl})
	}()

	return qr.sendAndReceiveMessage(ctx, input, &sqs.ReceiveMessageInput{
		QueueUrl:              out.QueueUrl,
		AttributeNames:        qr.attributeNames,
		MessageAttributeNames: []string{"All"},
	})
}

func (qr *QueueRequester) sendAndReceiveMessage(ctx context.Context, sndInput *sqs.SendMessageInput, rcvInput *sqs.ReceiveMessageInput, timeout time.Duration) (*types.Message, error) {
	_, err := qr.sqsClient.SendMessage(ctx, sndInput)
	if err != nil {
		return nil, err
	}

	startTime := time.Now()

	for {
		select {
		case <-time.After(qr.responsePollInterval):
			if time.Since(startTime) > timeout {
				return nil, ErrTimeout
			}

			out, err := qr.sqsClient.ReceiveMessage(ctx, rcvInput)
			if err != nil {
				return nil, err
			}

			if len(out.Messages) > 0 {
				return &out.Messages[0], nil
			}
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
