// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about Docker images that are managed by AWS CodeBuild.
func (c *Client) ListCuratedEnvironmentImages(ctx context.Context, params *ListCuratedEnvironmentImagesInput, optFns ...func(*Options)) (*ListCuratedEnvironmentImagesOutput, error) {
	stack := middleware.NewStack("ListCuratedEnvironmentImages", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListCuratedEnvironmentImagesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCuratedEnvironmentImages(options.Region), middleware.Before)

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
			OperationName: "ListCuratedEnvironmentImages",
			Err:           err,
		}
	}
	out := result.(*ListCuratedEnvironmentImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCuratedEnvironmentImagesInput struct {
}

type ListCuratedEnvironmentImagesOutput struct {
	// Information about supported platforms for Docker images that are managed by AWS
	// CodeBuild.
	Platforms []*types.EnvironmentPlatform

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListCuratedEnvironmentImagesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListCuratedEnvironmentImages{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCuratedEnvironmentImages{}, middleware.After)
}

func newServiceMetadataMiddleware_opListCuratedEnvironmentImages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "ListCuratedEnvironmentImages",
	}
}