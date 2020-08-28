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

// Creates a backup for an existing table. Each time you create an on-demand
// backup, the entire table data is backed up. There is no limit to the number of
// on-demand backups that can be taken. When you create an on-demand backup, a time
// marker of the request is cataloged, and the backup is created asynchronously, by
// applying all changes until the time of the request to the last full table
// snapshot. Backup requests are processed instantaneously and become available for
// restore within minutes. You can call CreateBackup at a maximum rate of 50 times
// per second. All backups in DynamoDB work without consuming any provisioned
// throughput on the table. If you submit a backup request on 2018-12-14 at
// 14:25:00, the backup is guaranteed to contain all data committed to the table up
// to 14:24:00, and data committed after 14:26:00 will not be. The backup might
// contain data modifications made between 14:24:00 and 14:26:00. On-demand backup
// does not support causal consistency. Along with data, the following are also
// included on the backups:
//
//     * Global secondary indexes (GSIs)
//
//     * Local
// secondary indexes (LSIs)
//
//     * Streams
//
//     * Provisioned read and write
// capacity
func (c *Client) CreateBackup(ctx context.Context, params *CreateBackupInput, optFns ...func(*Options)) (*CreateBackupOutput, error) {
	stack := middleware.NewStack("CreateBackup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpCreateBackupMiddlewares(stack)
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
	addOpCreateBackupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBackup(options.Region), middleware.Before)

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
			OperationName: "CreateBackup",
			Err:           err,
		}
	}
	out := result.(*CreateBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBackupInput struct {
	// Specified name for the backup.
	BackupName *string
	// The name of the table.
	TableName *string
}

type CreateBackupOutput struct {
	// Contains the details of the backup created for the table.
	BackupDetails *types.BackupDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpCreateBackupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpCreateBackup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateBackup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateBackup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "CreateBackup",
	}
}
