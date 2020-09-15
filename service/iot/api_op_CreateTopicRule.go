// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a rule. Creating rules is an administrator-level action. Any user who
// has permission to create rules will be able to access data processed by the
// rule.
func (c *Client) CreateTopicRule(ctx context.Context, params *CreateTopicRuleInput, optFns ...func(*Options)) (*CreateTopicRuleOutput, error) {
	stack := middleware.NewStack("CreateTopicRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateTopicRuleMiddlewares(stack)
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
	addOpCreateTopicRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTopicRule(options.Region), middleware.Before)

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
			OperationName: "CreateTopicRule",
			Err:           err,
		}
	}
	out := result.(*CreateTopicRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CreateTopicRule operation.
type CreateTopicRuleInput struct {
	// The rule payload.
	TopicRulePayload *types.TopicRulePayload
	// The name of the rule.
	RuleName *string
	// Metadata which can be used to manage the topic rule. For URI Request parameters
	// use format: ...key1=value1&key2=value2... For the CLI command-line parameter use
	// format: --tags "key1=value1&key2=value2..." For the cli-input-json file use
	// format: "tags": "key1=value1&key2=value2..."
	Tags *string
}

type CreateTopicRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateTopicRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateTopicRule{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTopicRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTopicRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateTopicRule",
	}
}