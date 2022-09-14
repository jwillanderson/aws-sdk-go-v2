// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon EKS add-on. Amazon EKS add-ons help to automate the
// provisioning and lifecycle management of common operational software for Amazon
// EKS clusters. Amazon EKS add-ons require clusters running version 1.18 or later
// because Amazon EKS add-ons rely on the Server-side Apply Kubernetes feature,
// which is only available in Kubernetes 1.18 and later. For more information, see
// Amazon EKS add-ons
// (https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html) in the
// Amazon EKS User Guide.
func (c *Client) CreateAddon(ctx context.Context, params *CreateAddonInput, optFns ...func(*Options)) (*CreateAddonOutput, error) {
	if params == nil {
		params = &CreateAddonInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAddon", params, optFns, c.addOperationCreateAddonMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAddonOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAddonInput struct {

	// The name of the add-on. The name must match one of the names returned by
	// DescribeAddonVersions
	// (https://docs.aws.amazon.com/eks/latest/APIReference/API_DescribeAddonVersions.html).
	//
	// This member is required.
	AddonName *string

	// The name of the cluster to create the add-on for.
	//
	// This member is required.
	ClusterName *string

	// The version of the add-on. The version must match one of the versions returned
	// by DescribeAddonVersions
	// (https://docs.aws.amazon.com/eks/latest/APIReference/API_DescribeAddonVersions.html).
	AddonVersion *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientRequestToken *string

	// How to resolve field value conflicts for an Amazon EKS add-on. Conflicts are
	// handled based on the value you choose:
	//
	// * None – If the self-managed version of
	// the add-on is installed on your cluster, Amazon EKS doesn't change the value.
	// Creation of the add-on might fail.
	//
	// * Overwrite – If the self-managed version of
	// the add-on is installed on your cluster and the Amazon EKS default value is
	// different than the existing value, Amazon EKS changes the value to the Amazon
	// EKS default value.
	//
	// * Preserve – Not supported. You can set this value when
	// updating an add-on though. For more information, see UpdateAddon
	// (https://docs.aws.amazon.com/eks/latest/APIReference/API_UpdateAddon.html).
	//
	// If
	// you don't currently have the self-managed version of the add-on installed on
	// your cluster, the Amazon EKS add-on is installed. Amazon EKS sets all values to
	// default values, regardless of the option that you specify.
	ResolveConflicts types.ResolveConflicts

	// The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's
	// service account. The role must be assigned the IAM permissions required by the
	// add-on. If you don't specify an existing IAM role, then the add-on uses the
	// permissions assigned to the node IAM role. For more information, see Amazon EKS
	// node IAM role
	// (https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the
	// Amazon EKS User Guide. To specify an existing IAM role, you must have an IAM
	// OpenID Connect (OIDC) provider created for your cluster. For more information,
	// see Enabling IAM roles for service accounts on your cluster
	// (https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html)
	// in the Amazon EKS User Guide.
	ServiceAccountRoleArn *string

	// The metadata to apply to the cluster to assist with categorization and
	// organization. Each tag consists of a key and an optional value. You define both.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAddonOutput struct {

	// An Amazon EKS add-on. For more information, see Amazon EKS add-ons
	// (https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html) in the
	// Amazon EKS User Guide.
	Addon *types.Addon

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAddonMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAddon{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAddon{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateAddonMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAddonValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAddon(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAddon struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAddon) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAddon) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAddonInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAddonInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAddonMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAddon{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAddon(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "CreateAddon",
	}
}
