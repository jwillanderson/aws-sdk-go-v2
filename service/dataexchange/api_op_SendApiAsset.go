// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation invokes an API Gateway API asset. The request is proxied to the
// provider’s API Gateway API.
func (c *Client) SendApiAsset(ctx context.Context, params *SendApiAssetInput, optFns ...func(*Options)) (*SendApiAssetOutput, error) {
	if params == nil {
		params = &SendApiAssetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendApiAsset", params, optFns, c.addOperationSendApiAssetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendApiAssetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendApiAssetInput struct {

	// Asset ID value for the API request.
	//
	// This member is required.
	AssetId *string

	// Data set ID value for the API request.
	//
	// This member is required.
	DataSetId *string

	// Revision ID value for the API request.
	//
	// This member is required.
	RevisionId *string

	// The request body.
	Body *string

	// HTTP method value for the API request. Alternatively, you can use the
	// appropriate verb in your request.
	Method *string

	// URI path value for the API request. Alternatively, you can set the URI path
	// directly by invoking /v1/{pathValue}.
	Path *string

	// Attach query string parameters to the end of the URI (for example,
	// /v1/examplePath?exampleParam=exampleValue).
	QueryStringParameters map[string]string

	// Any header value prefixed with x-amzn-dataexchange-header- will have that
	// stripped before sending the Asset API request. Use this when you want to
	// override a header that AWS Data Exchange uses. Alternatively, you can use the
	// header without a prefix to the HTTP request.
	RequestHeaders map[string]string

	noSmithyDocumentSerde
}

type SendApiAssetOutput struct {

	// The response body from the underlying API tracked by the API asset.
	Body *string

	// The response headers from the underlying API tracked by the API asset.
	//
	// Map keys will be normalized to lower-case.
	ResponseHeaders map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendApiAssetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendApiAsset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendApiAsset{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opSendApiAssetMiddleware(stack); err != nil {
		return err
	}
	if err = addOpSendApiAssetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendApiAsset(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opSendApiAssetMiddleware struct {
}

func (*endpointPrefix_opSendApiAssetMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opSendApiAssetMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api-fulfill." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opSendApiAssetMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opSendApiAssetMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opSendApiAsset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dataexchange",
		OperationName: "SendApiAsset",
	}
}
