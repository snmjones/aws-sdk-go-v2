// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This test serializes strings, numbers, and boolean values.
func (c *Client) SimpleInputParams(ctx context.Context, params *SimpleInputParamsInput, optFns ...func(*Options)) (*SimpleInputParamsOutput, error) {
	stack := middleware.NewStack("SimpleInputParams", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSimpleInputParamsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSimpleInputParams(options.Region), middleware.Before)

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
			OperationName: "SimpleInputParams",
			Err:           err,
		}
	}
	out := result.(*SimpleInputParamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SimpleInputParamsInput struct {
	Foo     *string
	Bar     *string
	Baz     *bool
	Bam     *int32
	Boo     *float64
	Qux     []byte
	FooEnum types.FooEnum
}

type SimpleInputParamsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSimpleInputParamsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSimpleInputParams{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSimpleInputParams{}, middleware.After)
}

func newServiceMetadataMiddleware_opSimpleInputParams(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SimpleInputParams",
	}
}
