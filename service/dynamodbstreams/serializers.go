// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbstreams

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/encoding/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/encoding/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsAwsjson10_serializeOpDescribeStream struct {
}

func (*awsAwsjson10_serializeOpDescribeStream) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDescribeStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeStreamInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("DynamoDBStreams_20120810.DescribeStream")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDescribeStreamInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpGetRecords struct {
}

func (*awsAwsjson10_serializeOpGetRecords) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpGetRecords) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetRecordsInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("DynamoDBStreams_20120810.GetRecords")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentGetRecordsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpGetShardIterator struct {
}

func (*awsAwsjson10_serializeOpGetShardIterator) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpGetShardIterator) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetShardIteratorInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("DynamoDBStreams_20120810.GetShardIterator")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentGetShardIteratorInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpListStreams struct {
}

func (*awsAwsjson10_serializeOpListStreams) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpListStreams) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListStreamsInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("DynamoDBStreams_20120810.ListStreams")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentListStreamsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson10_serializeOpDocumentDescribeStreamInput(v *DescribeStreamInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExclusiveStartShardId != nil {
		ok := object.Key("ExclusiveStartShardId")
		ok.String(*v.ExclusiveStartShardId)
	}

	if v.Limit != nil {
		ok := object.Key("Limit")
		ok.Integer(*v.Limit)
	}

	if v.StreamArn != nil {
		ok := object.Key("StreamArn")
		ok.String(*v.StreamArn)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentGetRecordsInput(v *GetRecordsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Limit != nil {
		ok := object.Key("Limit")
		ok.Integer(*v.Limit)
	}

	if v.ShardIterator != nil {
		ok := object.Key("ShardIterator")
		ok.String(*v.ShardIterator)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentGetShardIteratorInput(v *GetShardIteratorInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.SequenceNumber != nil {
		ok := object.Key("SequenceNumber")
		ok.String(*v.SequenceNumber)
	}

	if v.ShardId != nil {
		ok := object.Key("ShardId")
		ok.String(*v.ShardId)
	}

	if len(v.ShardIteratorType) > 0 {
		ok := object.Key("ShardIteratorType")
		ok.String(string(v.ShardIteratorType))
	}

	if v.StreamArn != nil {
		ok := object.Key("StreamArn")
		ok.String(*v.StreamArn)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentListStreamsInput(v *ListStreamsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExclusiveStartStreamArn != nil {
		ok := object.Key("ExclusiveStartStreamArn")
		ok.String(*v.ExclusiveStartStreamArn)
	}

	if v.Limit != nil {
		ok := object.Key("Limit")
		ok.Integer(*v.Limit)
	}

	if v.TableName != nil {
		ok := object.Key("TableName")
		ok.String(*v.TableName)
	}

	return nil
}
