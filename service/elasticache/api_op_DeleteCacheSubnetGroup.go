// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a cache subnet group. You cannot delete a cache subnet group if it is
// associated with any clusters.
func (c *Client) DeleteCacheSubnetGroup(ctx context.Context, params *DeleteCacheSubnetGroupInput, optFns ...func(*Options)) (*DeleteCacheSubnetGroupOutput, error) {
	stack := middleware.NewStack("DeleteCacheSubnetGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteCacheSubnetGroupMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteCacheSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCacheSubnetGroup(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DeleteCacheSubnetGroup",
			Err:           err,
		}
	}
	out := result.(*DeleteCacheSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DeleteCacheSubnetGroup operation.
type DeleteCacheSubnetGroupInput struct {
	// The name of the cache subnet group to delete. Constraints: Must contain no more
	// than 255 alphanumeric characters or hyphens.
	CacheSubnetGroupName *string
}

type DeleteCacheSubnetGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteCacheSubnetGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteCacheSubnetGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteCacheSubnetGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteCacheSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DeleteCacheSubnetGroup",
	}
}