package virtualqueues

import (
  "github.com/aws/aws-sdk-go-v2/service/sqs"
  "github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type MessageContent struct {
	MessageContent    string
	MessageAttributes map[string]types.MessageAttributeValue
}

func (m *MessageContent) AsSendMessageRequest() *sqs.SendMessageInput {
	return &sqs.SendMessageInput{
		MessageBody:       &m.MessageContent,
		MessageAttributes: m.MessageAttributes,
	}
}

type QueueResponder struct {
  sqsClient *sqs.Client
}

func (s *QueueResponder) IsResponseMessageRequested(requestMessage MessageContent) bool {
	msgAttrs := requestMessage.MessageAttributes
	if len(msgAttrs) == 0 {
		return false
	}

	_, ok := msgAttrs[RESPONSE_QUEUE_URL_ATTRIBUTE_NAME]
	return ok
}

func (s *QueueResponder) SendResponseMessage(ctx context.Context, requestMessage types.Message, response MessageContent) error {
	msgAttrs := requestMessage.MessageAttributes
	if len(msgAttrs) == 0 {
		return errors.New("unable to send response message: no message attributes found")
	}

	attribute, ok := msgAttrs[RESPONSE_QUEUE_URL_ATTRIBUTE_NAME]
	if !ok {
		return errors.New("unable to send response message: no response queue url found")
	}

	replyQueueUrl := attribute.StringValue
	responseRequest := response.AsSendMessageRequest()
	responseRequest.QueueUrl = replyQueueUrl
	fmt.Println("sending response message to queue", *replyQueueUrl)
	_, err := s.sqs.SendMessage(ctx, responseRequest)

	return err
}


