// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Records a WorkflowExecutionTerminated event and forces closure of the workflow
// execution identified by the given domain, runId, and workflowId. The child
// policy, registered with the workflow type or specified when starting this
// execution, is applied to any open child workflow executions of this workflow
// execution.  <important> <p>If the identified workflow execution was in progress,
// it is terminated immediately.</p> </important> <note> <p>If a runId isn't
// specified, then the <code>WorkflowExecutionTerminated</code> event is recorded
// in the history of the current open workflow with the matching workflowId in the
// domain.</p> </note> <note> <p>You should consider using
// <a>RequestCancelWorkflowExecution</a> action instead because it allows the
// workflow to gracefully close while <a>TerminateWorkflowExecution</a>
// doesn't.</p> </note> <p> <b>Access Control</b> </p> <p>You can use IAM policies
// to control this action's access to Amazon SWF resources as follows:</p> <ul>
// <li> <p>Use a <code>Resource</code> element with the domain name to limit the
// action to only specified domains.</p> </li> <li> <p>Use an <code>Action</code>
// element to allow or deny permission to call this action.</p> </li> <li> <p>You
// cannot use an IAM policy to constrain this action's parameters.</p> </li> </ul>
// <p>If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's <code>cause</code> parameter is set to
// <code>OPERATION_NOT_PERMITTED</code>. For details and example IAM policies, see
// <a
// href="https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html">Using
// IAM to Manage Access to Amazon SWF Workflows</a> in the <i>Amazon SWF Developer
// Guide</i>.</p>
func (c *Client) TerminateWorkflowExecution(ctx context.Context, params *TerminateWorkflowExecutionInput, optFns ...func(*Options)) (*TerminateWorkflowExecutionOutput, error) {
	stack := middleware.NewStack("TerminateWorkflowExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpTerminateWorkflowExecutionMiddlewares(stack)
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
	addOpTerminateWorkflowExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateWorkflowExecution(options.Region), middleware.Before)

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
			OperationName: "TerminateWorkflowExecution",
			Err:           err,
		}
	}
	out := result.(*TerminateWorkflowExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TerminateWorkflowExecutionInput struct {
	// A descriptive reason for terminating the workflow execution.
	Reason *string
	// Details for terminating the workflow execution.
	Details *string
	// The domain of the workflow execution to terminate.
	Domain *string
	// If set, specifies the policy to use for the child workflow executions of the
	// workflow execution being terminated. This policy overrides the child policy
	// specified for the workflow execution at registration time or when starting the
	// execution. The supported child policies are:
	//
	//     * TERMINATE – The child
	// executions are terminated.
	//
	//     * REQUEST_CANCEL – A request to cancel is
	// attempted for each child execution by recording a
	// WorkflowExecutionCancelRequested event in its history. It is up to the decider
	// to take appropriate actions when it receives an execution history with this
	// event.
	//
	//     * ABANDON – No action is taken. The child executions continue to
	// run.
	//
	// A child policy for this workflow execution must be specified either as a
	// default for the workflow type or through this parameter. If neither this
	// parameter is set nor a default child policy was specified at registration time
	// then a fault is returned.
	ChildPolicy types.ChildPolicy
	// The workflowId of the workflow execution to terminate.
	WorkflowId *string
	// The runId of the workflow execution to terminate.
	RunId *string
}

type TerminateWorkflowExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpTerminateWorkflowExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpTerminateWorkflowExecution{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpTerminateWorkflowExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opTerminateWorkflowExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "TerminateWorkflowExecution",
	}
}