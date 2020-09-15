// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of the document classifiers that you have created.
func (c *Client) ListDocumentClassifiers(ctx context.Context, params *ListDocumentClassifiersInput, optFns ...func(*Options)) (*ListDocumentClassifiersOutput, error) {
	stack := middleware.NewStack("ListDocumentClassifiers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListDocumentClassifiersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDocumentClassifiers(options.Region), middleware.Before)

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
			OperationName: "ListDocumentClassifiers",
			Err:           err,
		}
	}
	out := result.(*ListDocumentClassifiersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDocumentClassifiersInput struct {
	// Identifies the next page of results to return.
	NextToken *string
	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int32
	// Filters the jobs that are returned. You can filter jobs on their name, status,
	// or the date and time that they were submitted. You can only set one filter at a
	// time.
	Filter *types.DocumentClassifierFilter
}

type ListDocumentClassifiersOutput struct {
	// Identifies the next page of results to return.
	NextToken *string
	// A list containing the properties of each job returned.
	DocumentClassifierPropertiesList []*types.DocumentClassifierProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListDocumentClassifiersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListDocumentClassifiers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDocumentClassifiers{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDocumentClassifiers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListDocumentClassifiers",
	}
}