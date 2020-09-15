// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a route from a Client VPN endpoint. You can only delete routes that you
// manually added using the CreateClientVpnRoute action. You cannot delete routes
// that were automatically added when associating a subnet. To remove routes that
// have been automatically added, disassociate the target subnet from the Client
// VPN endpoint.
func (c *Client) DeleteClientVpnRoute(ctx context.Context, params *DeleteClientVpnRouteInput, optFns ...func(*Options)) (*DeleteClientVpnRouteOutput, error) {
	stack := middleware.NewStack("DeleteClientVpnRoute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDeleteClientVpnRouteMiddlewares(stack)
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
	addOpDeleteClientVpnRouteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteClientVpnRoute(options.Region), middleware.Before)

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
			OperationName: "DeleteClientVpnRoute",
			Err:           err,
		}
	}
	out := result.(*DeleteClientVpnRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteClientVpnRouteInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The IPv4 address range, in CIDR notation, of the route to be deleted.
	DestinationCidrBlock *string
	// The ID of the target subnet used by the route.
	TargetVpcSubnetId *string
	// The ID of the Client VPN endpoint from which the route is to be deleted.
	ClientVpnEndpointId *string
}

type DeleteClientVpnRouteOutput struct {
	// The current state of the route.
	Status *types.ClientVpnRouteStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDeleteClientVpnRouteMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDeleteClientVpnRoute{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteClientVpnRoute{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteClientVpnRoute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteClientVpnRoute",
	}
}