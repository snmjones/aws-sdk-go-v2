// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the analytics configurations for the bucket. You can have up to 1,000
// analytics configurations per bucket. This operation supports list pagination and
// does not return more than 100 configurations at a time. You should always check
// the IsTruncated element in the response. If there are no more configurations to
// list, IsTruncated is set to false. If there are more configurations to list,
// IsTruncated is set to true, and there will be a value in NextContinuationToken.
// You use the NextContinuationToken value to continue the pagination of the list
// by passing the value in continuation-token in the request to GET the next page.
// To use this operation, you must have permissions to perform the
// s3:GetAnalyticsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see Permissions Related to Bucket Subresource
// Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html). For
// information about Amazon S3 analytics feature, see Amazon S3 Analytics – Storage
// Class Analysis
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html).
// The following operations are related to ListBucketAnalyticsConfigurations:
//
// *
// GetBucketAnalyticsConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketAnalyticsConfiguration.html)
//
// *
// DeleteBucketAnalyticsConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketAnalyticsConfiguration.html)
//
// *
// PutBucketAnalyticsConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAnalyticsConfiguration.html)
func (c *Client) ListBucketAnalyticsConfigurations(ctx context.Context, params *ListBucketAnalyticsConfigurationsInput, optFns ...func(*Options)) (*ListBucketAnalyticsConfigurationsOutput, error) {
	if params == nil {
		params = &ListBucketAnalyticsConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBucketAnalyticsConfigurations", params, optFns, addOperationListBucketAnalyticsConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBucketAnalyticsConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBucketAnalyticsConfigurationsInput struct {

	// The name of the bucket from which analytics configurations are retrieved.
	//
	// This member is required.
	Bucket *string

	// The ContinuationToken that represents a placeholder from where this request
	// should begin.
	ContinuationToken *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type ListBucketAnalyticsConfigurationsOutput struct {

	// The list of analytics configurations for a bucket.
	AnalyticsConfigurationList []*types.AnalyticsConfiguration

	// The marker that is used as a starting point for this analytics configuration
	// list response. This value is present if it was sent in the request.
	ContinuationToken *string

	// Indicates whether the returned list of analytics configurations is complete. A
	// value of true indicates that the list is not complete and the
	// NextContinuationToken will be provided for a subsequent request.
	IsTruncated *bool

	// NextContinuationToken is sent when isTruncated is true, which indicates that
	// there are more analytics configurations to list. The next request must include
	// this NextContinuationToken. The token is obfuscated and is not a usable value.
	NextContinuationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBucketAnalyticsConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListBucketAnalyticsConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListBucketAnalyticsConfigurations{}, middleware.After)
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
	if err = addOpListBucketAnalyticsConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBucketAnalyticsConfigurations(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addListBucketAnalyticsConfigurationsUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListBucketAnalyticsConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListBucketAnalyticsConfigurations",
	}
}

func addListBucketAnalyticsConfigurationsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{GetBucketFromInput: getBucketFromInput,
			SupportsAccelerate: supportAccelerate,
		},
		UsePathStyle:            options.UsePathStyle,
		UseAccelerate:           options.UseAccelerate,
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
