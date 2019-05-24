// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/StartReplicationTaskMessage
type StartReplicationTaskInput struct {
	_ struct{} `type:"structure"`

	// Indicates when you want a change data capture (CDC) operation to start. Use
	// either CdcStartPosition or CdcStartTime to specify when you want a CDC operation
	// to start. Specifying both values results in an error.
	//
	// The value can be in date, checkpoint, or LSN/SCN format.
	//
	// Date Example: --cdc-start-position “2018-03-08T12:12:12”
	//
	// Checkpoint Example: --cdc-start-position "checkpoint:V1#27#mysql-bin-changelog.157832:1975:-1:2002:677883278264080:mysql-bin-changelog.157832:1876#0#0#*#0#93"
	//
	// LSN Example: --cdc-start-position “mysql-bin-changelog.000024:373”
	CdcStartPosition *string `type:"string"`

	// Indicates the start time for a change data capture (CDC) operation. Use either
	// CdcStartTime or CdcStartPosition to specify when you want a CDC operation
	// to start. Specifying both values results in an error.
	//
	// Timestamp Example: --cdc-start-time “2018-03-08T12:12:12”
	CdcStartTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Indicates when you want a change data capture (CDC) operation to stop. The
	// value can be either server time or commit time.
	//
	// Server time example: --cdc-stop-position “server_time:3018-02-09T12:12:12”
	//
	// Commit time example: --cdc-stop-position “commit_time: 3018-02-09T12:12:12
	// “
	CdcStopPosition *string `type:"string"`

	// The Amazon Resource Name (ARN) of the replication task to be started.
	//
	// ReplicationTaskArn is a required field
	ReplicationTaskArn *string `type:"string" required:"true"`

	// The type of replication task.
	//
	// StartReplicationTaskType is a required field
	StartReplicationTaskType StartReplicationTaskTypeValue `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s StartReplicationTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartReplicationTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartReplicationTaskInput"}

	if s.ReplicationTaskArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationTaskArn"))
	}
	if len(s.StartReplicationTaskType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("StartReplicationTaskType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/StartReplicationTaskResponse
type StartReplicationTaskOutput struct {
	_ struct{} `type:"structure"`

	// The replication task started.
	ReplicationTask *ReplicationTask `type:"structure"`
}

// String returns the string representation
func (s StartReplicationTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartReplicationTask = "StartReplicationTask"

// StartReplicationTaskRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Starts the replication task.
//
// For more information about AWS DMS tasks, see Working with Migration Tasks
//  (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.html) in the
// AWS Database Migration Service User Guide.
//
//    // Example sending a request using StartReplicationTaskRequest.
//    req := client.StartReplicationTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/StartReplicationTask
func (c *Client) StartReplicationTaskRequest(input *StartReplicationTaskInput) StartReplicationTaskRequest {
	op := &aws.Operation{
		Name:       opStartReplicationTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartReplicationTaskInput{}
	}

	req := c.newRequest(op, input, &StartReplicationTaskOutput{})
	return StartReplicationTaskRequest{Request: req, Input: input, Copy: c.StartReplicationTaskRequest}
}

// StartReplicationTaskRequest is the request type for the
// StartReplicationTask API operation.
type StartReplicationTaskRequest struct {
	*aws.Request
	Input *StartReplicationTaskInput
	Copy  func(*StartReplicationTaskInput) StartReplicationTaskRequest
}

// Send marshals and sends the StartReplicationTask API request.
func (r StartReplicationTaskRequest) Send(ctx context.Context) (*StartReplicationTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartReplicationTaskResponse{
		StartReplicationTaskOutput: r.Request.Data.(*StartReplicationTaskOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartReplicationTaskResponse is the response type for the
// StartReplicationTask API operation.
type StartReplicationTaskResponse struct {
	*StartReplicationTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartReplicationTask request.
func (r *StartReplicationTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}