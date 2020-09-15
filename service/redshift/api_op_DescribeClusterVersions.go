// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns descriptions of the available Amazon Redshift cluster versions. You can
// call this operation even before creating any clusters to learn more about the
// Amazon Redshift versions. For more information about managing clusters, go to
// Amazon Redshift Clusters
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html) in
// the Amazon Redshift Cluster Management Guide.
func (c *Client) DescribeClusterVersions(ctx context.Context, params *DescribeClusterVersionsInput, optFns ...func(*Options)) (*DescribeClusterVersionsOutput, error) {
	stack := middleware.NewStack("DescribeClusterVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeClusterVersionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusterVersions(options.Region), middleware.Before)

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
			OperationName: "DescribeClusterVersions",
			Err:           err,
		}
	}
	out := result.(*DescribeClusterVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeClusterVersionsInput struct {
	// The name of a specific cluster parameter group family to return details for.
	// Constraints:
	//
	//     * Must be 1 to 255 alphanumeric characters
	//
	//     * First
	// character must be a letter
	//
	//     * Cannot end with a hyphen or contain two
	// consecutive hyphens
	ClusterParameterGroupFamily *string
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32
	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterVersions () request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records by
	// providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string
	// The specific cluster version to return. Example: 1.0
	ClusterVersion *string
}

// Contains the output from the DescribeClusterVersions () action.
type DescribeClusterVersionsOutput struct {
	// A list of Version elements.
	ClusterVersions []*types.ClusterVersion
	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeClusterVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeClusterVersions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeClusterVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeClusterVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeClusterVersions",
	}
}