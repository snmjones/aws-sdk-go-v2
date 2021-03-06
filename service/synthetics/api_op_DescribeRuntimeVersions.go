// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/synthetics/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of Synthetics canary runtime versions. For more information, see
// Canary Runtime Versions
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html).
func (c *Client) DescribeRuntimeVersions(ctx context.Context, params *DescribeRuntimeVersionsInput, optFns ...func(*Options)) (*DescribeRuntimeVersionsOutput, error) {
	if params == nil {
		params = &DescribeRuntimeVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRuntimeVersions", params, optFns, addOperationDescribeRuntimeVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRuntimeVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRuntimeVersionsInput struct {

	// Specify this parameter to limit how many runs are returned each time you use the
	// DescribeRuntimeVersions operation. If you omit this parameter, the default of
	// 100 is used.
	MaxResults *int32

	// A token that indicates that there is more data available. You can use this token
	// in a subsequent DescribeRuntimeVersions operation to retrieve the next set of
	// results.
	NextToken *string
}

type DescribeRuntimeVersionsOutput struct {

	// A token that indicates that there is more data available. You can use this token
	// in a subsequent DescribeRuntimeVersions operation to retrieve the next set of
	// results.
	NextToken *string

	// An array of objects that display the details about each Synthetics canary
	// runtime version.
	RuntimeVersions []*types.RuntimeVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeRuntimeVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRuntimeVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRuntimeVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRuntimeVersions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeRuntimeVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "DescribeRuntimeVersions",
	}
}
