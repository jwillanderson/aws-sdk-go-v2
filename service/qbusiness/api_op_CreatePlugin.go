// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Q plugin.
func (c *Client) CreatePlugin(ctx context.Context, params *CreatePluginInput, optFns ...func(*Options)) (*CreatePluginOutput, error) {
	if params == nil {
		params = &CreatePluginInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePlugin", params, optFns, c.addOperationCreatePluginMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePluginOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePluginInput struct {

	// The identifier of the application that will contain the plugin.
	//
	// This member is required.
	ApplicationId *string

	// Authentication configuration information for an Amazon Q plugin.
	//
	// This member is required.
	AuthConfiguration types.PluginAuthConfiguration

	// A the name for your plugin.
	//
	// This member is required.
	DisplayName *string

	// The source URL used for plugin configuration.
	//
	// This member is required.
	ServerUrl *string

	// The type of plugin you want to create.
	//
	// This member is required.
	Type types.PluginType

	// A token that you provide to identify the request to create your Amazon Q plugin.
	ClientToken *string

	// A list of key-value pairs that identify or categorize the data source
	// connector. You can also use tags to help control access to the data source
	// connector. Tag keys and values can consist of Unicode letters, digits, white
	// space, and any of the following symbols: _ . : / = + - @.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreatePluginOutput struct {

	// The Amazon Resource Name (ARN) of a plugin.
	PluginArn *string

	// The identifier of the plugin created.
	PluginId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePluginMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePlugin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePlugin{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePlugin"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreatePluginMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePluginValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePlugin(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreatePlugin struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePlugin) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePlugin) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePluginInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePluginInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreatePluginMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePlugin{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePlugin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePlugin",
	}
}
