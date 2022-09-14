// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an event data store. The required EventDataStore value is an ARN or the
// ID portion of the ARN. Other parameters are optional, but at least one optional
// parameter must be specified, or CloudTrail throws an error. RetentionPeriod is
// in days, and valid values are integers between 90 and 2557. By default,
// TerminationProtection is enabled. AdvancedEventSelectors includes or excludes
// management and data events in your event data store; for more information about
// AdvancedEventSelectors, see PutEventSelectorsRequest$AdvancedEventSelectors.
func (c *Client) UpdateEventDataStore(ctx context.Context, params *UpdateEventDataStoreInput, optFns ...func(*Options)) (*UpdateEventDataStoreOutput, error) {
	if params == nil {
		params = &UpdateEventDataStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEventDataStore", params, optFns, c.addOperationUpdateEventDataStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEventDataStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEventDataStoreInput struct {

	// The ARN (or the ID suffix of the ARN) of the event data store that you want to
	// update.
	//
	// This member is required.
	EventDataStore *string

	// The advanced event selectors used to select events for the event data store.
	AdvancedEventSelectors []types.AdvancedEventSelector

	// Specifies whether an event data store collects events from all regions, or only
	// from the region in which it was created.
	MultiRegionEnabled *bool

	// The event data store name.
	Name *string

	// Specifies whether an event data store collects events logged for an organization
	// in Organizations.
	OrganizationEnabled *bool

	// The retention period, in days.
	RetentionPeriod *int32

	// Indicates that termination protection is enabled and the event data store cannot
	// be automatically deleted.
	TerminationProtectionEnabled *bool

	noSmithyDocumentSerde
}

type UpdateEventDataStoreOutput struct {

	// The advanced event selectors that are applied to the event data store.
	AdvancedEventSelectors []types.AdvancedEventSelector

	// The timestamp that shows when an event data store was first created.
	CreatedTimestamp *time.Time

	// The ARN of the event data store.
	EventDataStoreArn *string

	// Indicates whether the event data store includes events from all regions, or only
	// from the region in which it was created.
	MultiRegionEnabled *bool

	// The name of the event data store.
	Name *string

	// Indicates whether an event data store is collecting logged events for an
	// organization in Organizations.
	OrganizationEnabled *bool

	// The retention period, in days.
	RetentionPeriod *int32

	// The status of an event data store. Values can be ENABLED and PENDING_DELETION.
	Status types.EventDataStoreStatus

	// Indicates whether termination protection is enabled for the event data store.
	TerminationProtectionEnabled *bool

	// The timestamp that shows when the event data store was last updated.
	// UpdatedTimestamp is always either the same or newer than the time shown in
	// CreatedTimestamp.
	UpdatedTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEventDataStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateEventDataStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateEventDataStore{}, middleware.After)
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
	if err = addOpUpdateEventDataStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEventDataStore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEventDataStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "UpdateEventDataStore",
	}
}
