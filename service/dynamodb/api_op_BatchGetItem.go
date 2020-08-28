// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The BatchGetItem operation returns the attributes of one or more items from one
// or more tables. You identify requested items by primary key. A single operation
// can retrieve up to 16 MB of data, which can contain as many as 100 items.
// BatchGetItem returns a partial result if the response size limit is exceeded,
// the table's provisioned throughput is exceeded, or an internal processing
// failure occurs. If a partial result is returned, the operation returns a value
// for UnprocessedKeys. You can use this value to retry the operation starting with
// the next item to get. If you request more than 100 items, BatchGetItem returns a
// ValidationException with the message "Too many items requested for the
// BatchGetItem call." For example, if you ask to retrieve 100 items, but each
// individual item is 300 KB in size, the system returns 52 items (so as not to
// exceed the 16 MB limit). It also returns an appropriate UnprocessedKeys value so
// you can get the next page of results. If desired, your application can include
// its own logic to assemble the pages of results into one dataset. If none of the
// items can be processed due to insufficient provisioned throughput on all of the
// tables in the request, then BatchGetItem returns a
// ProvisionedThroughputExceededException. If at least one of the items is
// successfully processed, then BatchGetItem completes successfully, while
// returning the keys of the unread items in UnprocessedKeys. If DynamoDB returns
// any unprocessed items, you should retry the batch operation on those items.
// However, we strongly recommend that you use an exponential backoff algorithm. If
// you retry the batch operation immediately, the underlying read or write requests
// can still fail due to throttling on the individual tables. If you delay the
// batch operation using exponential backoff, the individual requests in the batch
// are much more likely to succeed. For more information, see Batch Operations and
// Error Handling
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ErrorHandling.html#BatchOperations)
// in the Amazon DynamoDB Developer Guide. By default, BatchGetItem performs
// eventually consistent reads on every table in the request. If you want strongly
// consistent reads instead, you can set ConsistentRead to true for any or all
// tables. In order to minimize response latency, BatchGetItem retrieves items in
// parallel. When designing your application, keep in mind that DynamoDB does not
// return items in any particular order. To help parse the response by item,
// include the primary key values for the items in your request in the
// ProjectionExpression parameter. If a requested item does not exist, it is not
// returned in the result. Requests for nonexistent items consume the minimum read
// capacity units according to the type of read. For more information, see Working
// with Tables
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#CapacityUnitCalculations)
// in the Amazon DynamoDB Developer Guide.
func (c *Client) BatchGetItem(ctx context.Context, params *BatchGetItemInput, optFns ...func(*Options)) (*BatchGetItemOutput, error) {
	stack := middleware.NewStack("BatchGetItem", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpBatchGetItemMiddlewares(stack)
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
	addOpBatchGetItemValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetItem(options.Region), middleware.Before)

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
			OperationName: "BatchGetItem",
			Err:           err,
		}
	}
	out := result.(*BatchGetItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a BatchGetItem operation.
type BatchGetItemInput struct {
	// A map of one or more table names and, for each table, a map that describes one
	// or more items to retrieve from that table. Each table name can be used only once
	// per BatchGetItem request. Each element in the map of items to retrieve consists
	// of the following:
	//
	//     * ConsistentRead - If true, a strongly consistent read is
	// used; if false (the default), an eventually consistent read is used.
	//
	//     *
	// ExpressionAttributeNames - One or more substitution tokens for attribute names
	// in the ProjectionExpression parameter. The following are some use cases for
	// using ExpressionAttributeNames:
	//
	//         * To access an attribute whose name
	// conflicts with a DynamoDB reserved word.
	//
	//         * To create a placeholder for
	// repeating occurrences of an attribute name in an expression.
	//
	//         * To
	// prevent special characters in an attribute name from being misinterpreted in an
	// expression.
	//
	//     Use the # character in an expression to dereference an
	// attribute name. For example, consider the following attribute name:
	//
	//         *
	// Percentile
	//
	//     The name of this attribute conflicts with a reserved word, so it
	// cannot be used directly in an expression. (For the complete list of reserved
	// words, see Reserved Words
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html)
	// in the Amazon DynamoDB Developer Guide). To work around this, you could specify
	// the following for ExpressionAttributeNames:
	//
	//         * {"#P":"Percentile"}
	//
	//
	// You could then use this substitution in an expression, as in this example:
	//
	//
	// * #P = :val
	//
	//     Tokens that begin with the : character are expression attribute
	// values, which are placeholders for the actual value at runtime. For more
	// information about expression attribute names, see Accessing Item Attributes
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html)
	// in the Amazon DynamoDB Developer Guide.
	//
	//     * Keys - An array of primary key
	// attribute values that define specific items in the table. For each primary key,
	// you must provide all of the key attributes. For example, with a simple primary
	// key, you only need to provide the partition key value. For a composite key, you
	// must provide both the partition key value and the sort key value.
	//
	//     *
	// ProjectionExpression - A string that identifies one or more attributes to
	// retrieve from the table. These attributes can include scalars, sets, or elements
	// of a JSON document. The attributes in the expression must be separated by
	// commas. If no attribute names are specified, then all attributes are returned.
	// If any of the requested attributes are not found, they do not appear in the
	// result. For more information, see Accessing Item Attributes
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html)
	// in the Amazon DynamoDB Developer Guide.
	//
	//     * AttributesToGet - This is a
	// legacy parameter. Use ProjectionExpression instead. For more information, see
	// AttributesToGet
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributesToGet.html)
	// in the Amazon DynamoDB Developer Guide.  </p> </li> </ul>
	RequestItems map[string]*types.KeysAndAttributes
	// Determines the level of detail about provisioned throughput consumption that is
	// returned in the response:
	//
	//     * INDEXES - The response includes the aggregate
	// ConsumedCapacity for the operation, together with ConsumedCapacity for each
	// table and secondary index that was accessed. Note that some operations, such as
	// GetItem and BatchGetItem, do not access any indexes at all. In these cases,
	// specifying INDEXES will only return ConsumedCapacity information for table(s).
	//
	//
	// * TOTAL - The response includes only the aggregate ConsumedCapacity for the
	// operation.
	//
	//     * NONE - No ConsumedCapacity details are included in the
	// response.
	ReturnConsumedCapacity types.ReturnConsumedCapacity
}

// Represents the output of a BatchGetItem operation.
type BatchGetItemOutput struct {
	// The read capacity units consumed by the entire BatchGetItem operation. Each
	// element consists of:
	//
	//     * TableName - The table that consumed the provisioned
	// throughput.
	//
	//     * CapacityUnits - The total number of capacity units consumed.
	ConsumedCapacity []*types.ConsumedCapacity
	// A map of table name to a list of items. Each object in Responses consists of a
	// table name, along with a map of attribute data consisting of the data type and
	// attribute value.
	Responses map[string][]map[string]*types.AttributeValue
	// A map of tables and their respective keys that were not processed with the
	// current response. The UnprocessedKeys value is in the same form as RequestItems,
	// so the value can be provided directly to a subsequent BatchGetItem operation.
	// For more information, see RequestItems in the Request Parameters section. Each
	// element consists of:
	//
	//     * Keys - An array of primary key attribute values that
	// define specific items in the table.
	//
	//     * ProjectionExpression - One or more
	// attributes to be retrieved from the table or index. By default, all attributes
	// are returned. If a requested attribute is not found, it does not appear in the
	// result.
	//
	//     * ConsistentRead - The consistency of a read operation. If set to
	// true, then a strongly consistent read is used; otherwise, an eventually
	// consistent read is used.
	//
	// If there are no unprocessed keys remaining, the
	// response contains an empty UnprocessedKeys map.
	UnprocessedKeys map[string]*types.KeysAndAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpBatchGetItemMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpBatchGetItem{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpBatchGetItem{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchGetItem(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "BatchGetItem",
	}
}
