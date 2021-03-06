// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/worklink/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of website authorization providers associated with a specified
// fleet.
func (c *Client) ListWebsiteAuthorizationProviders(ctx context.Context, params *ListWebsiteAuthorizationProvidersInput, optFns ...func(*Options)) (*ListWebsiteAuthorizationProvidersOutput, error) {
	if params == nil {
		params = &ListWebsiteAuthorizationProvidersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWebsiteAuthorizationProviders", params, optFns, addOperationListWebsiteAuthorizationProvidersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWebsiteAuthorizationProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWebsiteAuthorizationProvidersInput struct {

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string

	// The maximum number of results to be included in the next page.
	MaxResults *int32

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string
}

type ListWebsiteAuthorizationProvidersOutput struct {

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	// The website authorization providers.
	WebsiteAuthorizationProviders []*types.WebsiteAuthorizationProviderSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListWebsiteAuthorizationProvidersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWebsiteAuthorizationProviders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWebsiteAuthorizationProviders{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addOpListWebsiteAuthorizationProvidersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWebsiteAuthorizationProviders(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListWebsiteAuthorizationProviders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "ListWebsiteAuthorizationProviders",
	}
}
