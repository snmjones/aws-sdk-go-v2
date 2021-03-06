// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the active violations for a given Device Defender security profile.
func (c *Client) ListActiveViolations(ctx context.Context, params *ListActiveViolationsInput, optFns ...func(*Options)) (*ListActiveViolationsOutput, error) {
	if params == nil {
		params = &ListActiveViolationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListActiveViolations", params, optFns, addOperationListActiveViolationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListActiveViolationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListActiveViolationsInput struct {

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// The name of the Device Defender security profile for which violations are
	// listed.
	SecurityProfileName *string

	// The name of the thing whose active violations are listed.
	ThingName *string
}

type ListActiveViolationsOutput struct {

	// The list of active violations.
	ActiveViolations []*types.ActiveViolation

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListActiveViolationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListActiveViolations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListActiveViolations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListActiveViolations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListActiveViolations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListActiveViolations",
	}
}
