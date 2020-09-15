// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a license configuration. A license configuration is an abstraction of a
// customer license agreement that can be consumed and enforced by License Manager.
// Components include specifications for the license type (licensing by instance,
// socket, CPU, or vCPU), allowed tenancy (shared tenancy, Dedicated Instance,
// Dedicated Host, or all of these), host affinity (how long a VM must be
// associated with a host), and the number of licenses purchased and used.
func (c *Client) CreateLicenseConfiguration(ctx context.Context, params *CreateLicenseConfigurationInput, optFns ...func(*Options)) (*CreateLicenseConfigurationOutput, error) {
	stack := middleware.NewStack("CreateLicenseConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateLicenseConfigurationMiddlewares(stack)
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
	addOpCreateLicenseConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLicenseConfiguration(options.Region), middleware.Before)

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
			OperationName: "CreateLicenseConfiguration",
			Err:           err,
		}
	}
	out := result.(*CreateLicenseConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLicenseConfigurationInput struct {
	// Indicates whether hard or soft license enforcement is used. Exceeding a hard
	// limit blocks the launch of new instances.
	LicenseCountHardLimit *bool
	// License rules. The syntax is #name=value (for example,
	// #allowedTenancy=EC2-DedicatedHost). Available rules vary by dimension.
	//
	//     *
	// Cores dimension: allowedTenancy | maximumCores | minimumCores
	//
	//     * Instances
	// dimension: allowedTenancy | maximumCores | minimumCores | maximumSockets |
	// minimumSockets | maximumVcpus | minimumVcpus
	//
	//     * Sockets dimension:
	// allowedTenancy | maximumSockets | minimumSockets
	//
	//     * vCPUs dimension:
	// allowedTenancy | honorVcpuOptimization | maximumVcpus | minimumVcpus
	LicenseRules []*string
	// Product information.
	ProductInformationList []*types.ProductInformation
	// Description of the license configuration.
	Description *string
	// Number of licenses managed by the license configuration.
	LicenseCount *int64
	// Name of the license configuration.
	Name *string
	// Dimension used to track the license inventory.
	LicenseCountingType types.LicenseCountingType
	// Tags to add to the license configuration.
	Tags []*types.Tag
}

type CreateLicenseConfigurationOutput struct {
	// Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateLicenseConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLicenseConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLicenseConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateLicenseConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "CreateLicenseConfiguration",
	}
}