// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the ServiceQuotaTemplateAssociationStatus value from the service. Use
// this action to determine if the Service Quota template is associated, or
// enabled.
func (c *Client) GetAssociationForServiceQuotaTemplate(ctx context.Context, params *GetAssociationForServiceQuotaTemplateInput, optFns ...func(*Options)) (*GetAssociationForServiceQuotaTemplateOutput, error) {
	stack := middleware.NewStack("GetAssociationForServiceQuotaTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetAssociationForServiceQuotaTemplateMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssociationForServiceQuotaTemplate(options.Region), middleware.Before)

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
			OperationName: "GetAssociationForServiceQuotaTemplate",
			Err:           err,
		}
	}
	out := result.(*GetAssociationForServiceQuotaTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssociationForServiceQuotaTemplateInput struct {
}

type GetAssociationForServiceQuotaTemplateOutput struct {
	// Specifies whether the template is ASSOCIATED or DISASSOCIATED. If the template
	// is ASSOCIATED, then it requests service quota increases for all new accounts
	// created in your organization.
	ServiceQuotaTemplateAssociationStatus types.ServiceQuotaTemplateAssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetAssociationForServiceQuotaTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetAssociationForServiceQuotaTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAssociationForServiceQuotaTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAssociationForServiceQuotaTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "GetAssociationForServiceQuotaTemplate",
	}
}