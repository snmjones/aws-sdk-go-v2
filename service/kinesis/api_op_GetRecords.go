// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets data records from a Kinesis data stream's shard. Specify a shard iterator
// using the ShardIterator parameter. The shard iterator specifies the position in
// the shard from which you want to start reading data records sequentially. If
// there are no records available in the portion of the shard that the iterator
// points to, GetRecords () returns an empty list. It might take multiple calls to
// get to a portion of the shard that contains records. You can scale by
// provisioning multiple shards per stream while considering service limits (for
// more information, see Amazon Kinesis Data Streams Limits
// (https://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html)
// in the Amazon Kinesis Data Streams Developer Guide). Your application should
// have one thread per shard, each reading continuously from its stream. To read
// from a stream continually, call GetRecords () in a loop. Use GetShardIterator ()
// to get the shard iterator to specify in the first GetRecords () call. GetRecords
// () returns a new shard iterator in NextShardIterator. Specify the shard iterator
// returned in NextShardIterator in subsequent calls to GetRecords (). If the shard
// has been closed, the shard iterator can't return more data and GetRecords ()
// returns null in NextShardIterator. You can terminate the loop when the shard is
// closed, or when the shard iterator reaches the record with the sequence number
// or other attribute that marks it as the last record to process. Each data record
// can be up to 1 MiB in size, and each shard can read up to 2 MiB per second. You
// can ensure that your calls don't exceed the maximum supported size or throughput
// by using the Limit parameter to specify the maximum number of records that
// GetRecords () can return. Consider your average record size when determining
// this limit. The maximum number of records that can be returned per call is
// 10,000.  <p>The size of the data returned by <a>GetRecords</a> varies depending
// on the utilization of the shard. The maximum size of data that <a>GetRecords</a>
// can return is 10 MiB. If a call returns this amount of data, subsequent calls
// made within the next 5 seconds throw
// <code>ProvisionedThroughputExceededException</code>. If there is insufficient
// provisioned throughput on the stream, subsequent calls made within the next 1
// second throw <code>ProvisionedThroughputExceededException</code>.
// <a>GetRecords</a> doesn't return any data when it throws an exception. For this
// reason, we recommend that you wait 1 second between calls to <a>GetRecords</a>.
// However, it's possible that the application will get exceptions for longer than
// 1 second.</p> <p>To detect whether the application is falling behind in
// processing, you can use the <code>MillisBehindLatest</code> response attribute.
// You can also monitor the stream using CloudWatch metrics and other mechanisms
// (see <a
// href="https://docs.aws.amazon.com/kinesis/latest/dev/monitoring.html">Monitoring</a>
// in the <i>Amazon Kinesis Data Streams Developer Guide</i>).</p> <p>Each Amazon
// Kinesis record includes a value, <code>ApproximateArrivalTimestamp</code>, that
// is set when a stream successfully receives and stores a record. This is commonly
// referred to as a server-side time stamp, whereas a client-side time stamp is set
// when a data producer creates or sends the record to a stream (a data producer is
// any data source putting data records into a stream, for example with
// <a>PutRecords</a>). The time stamp has millisecond precision. There are no
// guarantees about the time stamp accuracy, or that the time stamp is always
// increasing. For example, records in a shard or across a stream might have time
// stamps that are out of order.</p> <p>This operation has a limit of five
// transactions per second per account.</p>
func (c *Client) GetRecords(ctx context.Context, params *GetRecordsInput, optFns ...func(*Options)) (*GetRecordsOutput, error) {
	stack := middleware.NewStack("GetRecords", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetRecordsMiddlewares(stack)
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
	addOpGetRecordsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecords(options.Region), middleware.Before)

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
			OperationName: "GetRecords",
			Err:           err,
		}
	}
	out := result.(*GetRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for GetRecords ().
type GetRecordsInput struct {
	// The maximum number of records to return. Specify a value of up to 10,000. If you
	// specify a value that is greater than 10,000, GetRecords () throws
	// InvalidArgumentException.
	Limit *int32
	// The position in the shard from which you want to start sequentially reading data
	// records. A shard iterator specifies this position using the sequence number of a
	// data record in the shard.
	ShardIterator *string
}

// Represents the output for GetRecords ().
type GetRecordsOutput struct {
	// The number of milliseconds the GetRecords () response is from the tip of the
	// stream, indicating how far behind current time the consumer is. A value of zero
	// indicates that record processing is caught up, and there are no new records to
	// process at this moment.
	MillisBehindLatest *int64
	// The next position in the shard from which to start sequentially reading data
	// records. If set to null, the shard has been closed and the requested iterator
	// does not return any more data.
	NextShardIterator *string
	// The data records retrieved from the shard.
	Records []*types.Record

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetRecordsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetRecords{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRecords{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRecords(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "GetRecords",
	}
}