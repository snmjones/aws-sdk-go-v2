// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the configuration required to provision the specified
// product using the specified provisioning artifact. If the output contains a
// TagOption key with an empty list of values, there is a TagOption conflict for
// that key. The end user cannot take action to fix the conflict, and launch is not
// blocked. In subsequent calls to ProvisionProduct (), do not include conflicted
// TagOption keys as tags, or this causes the error "Parameter validation failed:
// Missing required parameter in Tags[N]:Value". Tag the provisioned product with
// the value sc-tagoption-conflict-portfolioId-productId.
func (c *Client) DescribeProvisioningParameters(ctx context.Context, params *DescribeProvisioningParametersInput, optFns ...func(*Options)) (*DescribeProvisioningParametersOutput, error) {
	stack := middleware.NewStack("DescribeProvisioningParameters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeProvisioningParametersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProvisioningParameters(options.Region), middleware.Before)

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
			OperationName: "DescribeProvisioningParameters",
			Err:           err,
		}
	}
	out := result.(*DescribeProvisioningParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProvisioningParametersInput struct {
	// The identifier of the provisioning artifact. You must provide the name or ID,
	// but not both.
	ProvisioningArtifactId *string
	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
	// The name of the path. You must provide the name or ID, but not both.
	PathName *string
	// The name of the provisioning artifact. You must provide the name or ID, but not
	// both.
	ProvisioningArtifactName *string
	// The product identifier. You must provide the product name or ID, but not both.
	ProductId *string
	// The name of the product. You must provide the name or ID, but not both.
	ProductName *string
	// The path identifier of the product. This value is optional if the product has a
	// default path, and required if the product has more than one path. To list the
	// paths for a product, use ListLaunchPaths (). You must provide the name or ID,
	// but not both.
	PathId *string
}

type DescribeProvisioningParametersOutput struct {
	// Information about the constraints used to provision the product.
	ConstraintSummaries []*types.ConstraintSummary
	// Information about the parameters used to provision the product.
	ProvisioningArtifactParameters []*types.ProvisioningArtifactParameter
	// An object that contains information about preferences, such as regions and
	// accounts, for the provisioning artifact.
	ProvisioningArtifactPreferences *types.ProvisioningArtifactPreferences
	// Information about the TagOptions associated with the resource.
	TagOptions []*types.TagOptionSummary
	// The output of the provisioning artifact.
	ProvisioningArtifactOutputs []*types.ProvisioningArtifactOutput
	// Any additional metadata specifically related to the provisioning of the product.
	// For example, see the Version field of the CloudFormation template.
	UsageInstructions []*types.UsageInstruction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeProvisioningParametersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeProvisioningParameters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeProvisioningParameters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeProvisioningParameters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DescribeProvisioningParameters",
	}
}