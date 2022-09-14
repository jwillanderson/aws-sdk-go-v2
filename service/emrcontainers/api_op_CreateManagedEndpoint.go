// Code generated by smithy-go-codegen DO NOT EDIT.

package emrcontainers

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emrcontainers/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a managed endpoint. A managed endpoint is a gateway that connects EMR
// Studio to Amazon EMR on EKS so that EMR Studio can communicate with your virtual
// cluster.
func (c *Client) CreateManagedEndpoint(ctx context.Context, params *CreateManagedEndpointInput, optFns ...func(*Options)) (*CreateManagedEndpointOutput, error) {
	if params == nil {
		params = &CreateManagedEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateManagedEndpoint", params, optFns, c.addOperationCreateManagedEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateManagedEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateManagedEndpointInput struct {

	// The client idempotency token for this create call.
	//
	// This member is required.
	ClientToken *string

	// The ARN of the execution role.
	//
	// This member is required.
	ExecutionRoleArn *string

	// The name of the managed endpoint.
	//
	// This member is required.
	Name *string

	// The Amazon EMR release version.
	//
	// This member is required.
	ReleaseLabel *string

	// The type of the managed endpoint.
	//
	// This member is required.
	Type *string

	// The ID of the virtual cluster for which a managed endpoint is created.
	//
	// This member is required.
	VirtualClusterId *string

	// The certificate ARN provided by users for the managed endpoint. This field is
	// under deprecation and will be removed in future releases.
	//
	// Deprecated: Customer provided certificate-arn is deprecated and would be removed
	// in future.
	CertificateArn *string

	// The configuration settings that will be used to override existing
	// configurations.
	ConfigurationOverrides *types.ConfigurationOverrides

	// The tags of the managed endpoint.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateManagedEndpointOutput struct {

	// The output contains the ARN of the managed endpoint.
	Arn *string

	// The output contains the ID of the managed endpoint.
	Id *string

	// The output contains the name of the managed endpoint.
	Name *string

	// The output contains the ID of the virtual cluster.
	VirtualClusterId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateManagedEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateManagedEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateManagedEndpoint{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateManagedEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateManagedEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateManagedEndpoint(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateManagedEndpoint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateManagedEndpoint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateManagedEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateManagedEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateManagedEndpointInput ")
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
func addIdempotencyToken_opCreateManagedEndpointMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateManagedEndpoint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateManagedEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "emr-containers",
		OperationName: "CreateManagedEndpoint",
	}
}
