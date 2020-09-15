// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a source from an existing flow. This request can be made only if there
// is more than one source on the flow.
func (c *Client) RemoveFlowSource(ctx context.Context, params *RemoveFlowSourceInput, optFns ...func(*Options)) (*RemoveFlowSourceOutput, error) {
	stack := middleware.NewStack("RemoveFlowSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRemoveFlowSourceMiddlewares(stack)
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
	addOpRemoveFlowSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveFlowSource(options.Region), middleware.Before)

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
			OperationName: "RemoveFlowSource",
			Err:           err,
		}
	}
	out := result.(*RemoveFlowSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveFlowSourceInput struct {
	// The flow that you want to remove a source from.
	FlowArn *string
	// The ARN of the source that you want to remove.
	SourceArn *string
}

type RemoveFlowSourceOutput struct {
	// The ARN of the flow that is associated with the source you removed.
	FlowArn *string
	// The ARN of the source that was removed.
	SourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRemoveFlowSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRemoveFlowSource{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveFlowSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveFlowSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "RemoveFlowSource",
	}
}