// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches for system instances in the user's account.
func (c *Client) SearchSystemInstances(ctx context.Context, params *SearchSystemInstancesInput, optFns ...func(*Options)) (*SearchSystemInstancesOutput, error) {
	if params == nil {
		params = &SearchSystemInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchSystemInstances", params, optFns, addOperationSearchSystemInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchSystemInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchSystemInstancesInput struct {

	// Optional filter to apply to the search. Valid filters are SYSTEM_TEMPLATE_ID,
	// STATUS, and GREENGRASS_GROUP_NAME. Multiple filters function as OR criteria in
	// the query. Multiple values passed inside the filter function as AND criteria.
	Filters []*types.SystemInstanceFilter

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string
}

type SearchSystemInstancesOutput struct {

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// An array of objects that contain summary data abour the system instances in the
	// result set.
	Summaries []*types.SystemInstanceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchSystemInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchSystemInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchSystemInstances{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchSystemInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchSystemInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "SearchSystemInstances",
	}
}
