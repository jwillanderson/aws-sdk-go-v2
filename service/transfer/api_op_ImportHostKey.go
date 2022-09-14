// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a host key to the server specified by the ServerId parameter.
func (c *Client) ImportHostKey(ctx context.Context, params *ImportHostKeyInput, optFns ...func(*Options)) (*ImportHostKeyOutput, error) {
	if params == nil {
		params = &ImportHostKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportHostKey", params, optFns, c.addOperationImportHostKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportHostKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportHostKeyInput struct {

	// The public key portion of an SSH key pair. Transfer Family accepts RSA, ECDSA,
	// and ED25519 keys.
	//
	// This member is required.
	HostKeyBody *string

	// Provide the ID of the server that contains the host key that you are importing.
	//
	// This member is required.
	ServerId *string

	// Enter a text description to identify this host key.
	Description *string

	// Key-value pairs that can be used to group and search for host keys.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type ImportHostKeyOutput struct {

	// Returns the host key ID for the imported key.
	//
	// This member is required.
	HostKeyId *string

	// Returns the server ID that contains the imported key.
	//
	// This member is required.
	ServerId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportHostKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportHostKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportHostKey{}, middleware.After)
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
	if err = addOpImportHostKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportHostKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportHostKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "ImportHostKey",
	}
}
