// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an AWS Fargate profile. When you delete a Fargate profile, any pods
// running on Fargate that were created with the profile are deleted. If those pods
// match another Fargate profile, then they are scheduled on Fargate with that
// profile. If they no longer match any Fargate profiles, then they are not
// scheduled on Fargate and they may remain in a pending state. Only one Fargate
// profile in a cluster can be in the DELETING status at a time. You must wait for
// a Fargate profile to finish deleting before you can delete any other profiles in
// that cluster.
func (c *Client) DeleteFargateProfile(ctx context.Context, params *DeleteFargateProfileInput, optFns ...func(*Options)) (*DeleteFargateProfileOutput, error) {
	stack := middleware.NewStack("DeleteFargateProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteFargateProfileMiddlewares(stack)
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
	addOpDeleteFargateProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFargateProfile(options.Region), middleware.Before)

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
			OperationName: "DeleteFargateProfile",
			Err:           err,
		}
	}
	out := result.(*DeleteFargateProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFargateProfileInput struct {
	// The name of the Fargate profile to delete.
	FargateProfileName *string
	// The name of the Amazon EKS cluster associated with the Fargate profile to
	// delete.
	ClusterName *string
}

type DeleteFargateProfileOutput struct {
	// The deleted Fargate profile.
	FargateProfile *types.FargateProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteFargateProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteFargateProfile{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteFargateProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteFargateProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "DeleteFargateProfile",
	}
}