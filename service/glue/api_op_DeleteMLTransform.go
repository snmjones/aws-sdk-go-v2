// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an AWS Glue machine learning transform. Machine learning transforms are
// a special type of transform that use machine learning to learn the details of
// the transformation to be performed by learning from examples provided by humans.
// These transformations are then saved by AWS Glue. If you no longer need a
// transform, you can delete it by calling DeleteMLTransforms. However, any AWS
// Glue jobs that still reference the deleted transform will no longer succeed.
func (c *Client) DeleteMLTransform(ctx context.Context, params *DeleteMLTransformInput, optFns ...func(*Options)) (*DeleteMLTransformOutput, error) {
	stack := middleware.NewStack("DeleteMLTransform", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteMLTransformMiddlewares(stack)
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
	addOpDeleteMLTransformValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMLTransform(options.Region), middleware.Before)

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
			OperationName: "DeleteMLTransform",
			Err:           err,
		}
	}
	out := result.(*DeleteMLTransformOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMLTransformInput struct {
	// The unique identifier of the transform to delete.
	TransformId *string
}

type DeleteMLTransformOutput struct {
	// The unique identifier of the transform that was deleted.
	TransformId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteMLTransformMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteMLTransform{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteMLTransform{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteMLTransform(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeleteMLTransform",
	}
}