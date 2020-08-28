// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This examples serializes a structure in the payload. Note that serializing a
// structure changes the wrapper element name to match the targeted structure.
func (c *Client) HttpPayloadWithStructure(ctx context.Context, params *HttpPayloadWithStructureInput, optFns ...func(*Options)) (*HttpPayloadWithStructureOutput, error) {
	stack := middleware.NewStack("HttpPayloadWithStructure", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpHttpPayloadWithStructureMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPayloadWithStructure(options.Region), middleware.Before)

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
			OperationName: "HttpPayloadWithStructure",
			Err:           err,
		}
	}
	out := result.(*HttpPayloadWithStructureOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadWithStructureInput struct {
	Nested *types.NestedPayload
}

type HttpPayloadWithStructureOutput struct {
	Nested *types.NestedPayload

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpHttpPayloadWithStructureMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpHttpPayloadWithStructure{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpHttpPayloadWithStructure{}, middleware.After)
}

func newServiceMetadataMiddleware_opHttpPayloadWithStructure(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpPayloadWithStructure",
	}
}
