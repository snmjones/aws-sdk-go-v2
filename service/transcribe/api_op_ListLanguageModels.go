// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides more information about the custom language models you've created. You
// can use the information in this list to find a specific custom language model.
// You can then use the operation to get more information about it.
func (c *Client) ListLanguageModels(ctx context.Context, params *ListLanguageModelsInput, optFns ...func(*Options)) (*ListLanguageModelsOutput, error) {
	if params == nil {
		params = &ListLanguageModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLanguageModels", params, optFns, addOperationListLanguageModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLanguageModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLanguageModelsInput struct {

	// The maximum number of language models to return in the response. If there are
	// fewer results in the list, the response contains only the actual results.
	MaxResults *int32

	// When specified, the custom language model names returned contain the substring
	// you've specified.
	NameContains *string

	// When included, fetches the next set of jobs if the result of the previous
	// request was truncated.
	NextToken *string

	// When specified, returns only custom language models with the specified status.
	// Language models are ordered by creation date, with the newest models first. If
	// you don't specify a status, Amazon Transcribe returns all custom language models
	// ordered by date.
	StatusEquals types.ModelStatus
}

type ListLanguageModelsOutput struct {

	// A list of objects containing information about custom language models.
	Models []*types.LanguageModel

	// The operation returns a page of jobs at a time. The maximum size of the list is
	// set by the MaxResults parameter. If there are more language models in the list
	// than the page size, Amazon Transcribe returns the NextPage token. Include the
	// token in the next request to the operation to return the next page of language
	// models.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListLanguageModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLanguageModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLanguageModels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLanguageModels(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListLanguageModels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "ListLanguageModels",
	}
}
