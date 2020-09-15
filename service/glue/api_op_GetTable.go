// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the Table definition in a Data Catalog for a specified table.
func (c *Client) GetTable(ctx context.Context, params *GetTableInput, optFns ...func(*Options)) (*GetTableOutput, error) {
	stack := middleware.NewStack("GetTable", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetTableMiddlewares(stack)
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
	addOpGetTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTable(options.Region), middleware.Before)

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
			OperationName: "GetTable",
			Err:           err,
		}
	}
	out := result.(*GetTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTableInput struct {
	// The ID of the Data Catalog where the table resides. If none is provided, the AWS
	// account ID is used by default.
	CatalogId *string
	// The name of the table for which to retrieve the definition. For Hive
	// compatibility, this name is entirely lowercase.
	Name *string
	// The name of the database in the catalog in which the table resides. For Hive
	// compatibility, this name is entirely lowercase.
	DatabaseName *string
}

type GetTableOutput struct {
	// The Table object that defines the specified table.
	Table *types.Table

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetTableMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetTable{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTable{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetTable",
	}
}