// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the membership details for specified member accounts for a behavior
// graph.
func (c *Client) GetMembers(ctx context.Context, params *GetMembersInput, optFns ...func(*Options)) (*GetMembersOutput, error) {
	stack := middleware.NewStack("GetMembers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetMembersMiddlewares(stack)
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
	addOpGetMembersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMembers(options.Region), middleware.Before)

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
			OperationName: "GetMembers",
			Err:           err,
		}
	}
	out := result.(*GetMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMembersInput struct {
	// The ARN of the behavior graph for which to request the member details.
	GraphArn *string
	// The list of AWS account identifiers for the member account for which to return
	// member details. You cannot use GetMembers to retrieve information about member
	// accounts that were removed from the behavior graph.
	AccountIds []*string
}

type GetMembersOutput struct {
	// The member account details that Detective is returning in response to the
	// request.
	MemberDetails []*types.MemberDetail
	// The requested member accounts for which Detective was unable to return member
	// details. For each account, provides the reason why the request could not be
	// processed.
	UnprocessedAccounts []*types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetMembersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetMembers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMembers{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetMembers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "detective",
		OperationName: "GetMembers",
	}
}