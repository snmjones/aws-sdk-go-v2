// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/EnableAWSOrganizationsAccessInput
type EnableAWSOrganizationsAccessInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableAWSOrganizationsAccessInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/EnableAWSOrganizationsAccessOutput
type EnableAWSOrganizationsAccessOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableAWSOrganizationsAccessOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableAWSOrganizationsAccess = "EnableAWSOrganizationsAccess"

// EnableAWSOrganizationsAccessRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Enable portfolio sharing feature through AWS Organizations. This API will
// allow Service Catalog to receive updates on your organization in order to
// sync your shares with the current structure. This API can only be called
// by the master account in the organization.
//
// By calling this API Service Catalog will make a call to organizations:EnableAWSServiceAccess
// on your behalf so that your shares can be in sync with any changes in your
// AWS Organizations structure.
//
//    // Example sending a request using EnableAWSOrganizationsAccessRequest.
//    req := client.EnableAWSOrganizationsAccessRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/EnableAWSOrganizationsAccess
func (c *Client) EnableAWSOrganizationsAccessRequest(input *EnableAWSOrganizationsAccessInput) EnableAWSOrganizationsAccessRequest {
	op := &aws.Operation{
		Name:       opEnableAWSOrganizationsAccess,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableAWSOrganizationsAccessInput{}
	}

	req := c.newRequest(op, input, &EnableAWSOrganizationsAccessOutput{})
	return EnableAWSOrganizationsAccessRequest{Request: req, Input: input, Copy: c.EnableAWSOrganizationsAccessRequest}
}

// EnableAWSOrganizationsAccessRequest is the request type for the
// EnableAWSOrganizationsAccess API operation.
type EnableAWSOrganizationsAccessRequest struct {
	*aws.Request
	Input *EnableAWSOrganizationsAccessInput
	Copy  func(*EnableAWSOrganizationsAccessInput) EnableAWSOrganizationsAccessRequest
}

// Send marshals and sends the EnableAWSOrganizationsAccess API request.
func (r EnableAWSOrganizationsAccessRequest) Send(ctx context.Context) (*EnableAWSOrganizationsAccessResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableAWSOrganizationsAccessResponse{
		EnableAWSOrganizationsAccessOutput: r.Request.Data.(*EnableAWSOrganizationsAccessOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableAWSOrganizationsAccessResponse is the response type for the
// EnableAWSOrganizationsAccess API operation.
type EnableAWSOrganizationsAccessResponse struct {
	*EnableAWSOrganizationsAccessOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableAWSOrganizationsAccess request.
func (r *EnableAWSOrganizationsAccessResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}