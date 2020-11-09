// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	acceptencodingcust "github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding"
	"github.com/aws/aws-sdk-go-v2/service/internal/s3shared"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/logging"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"net/http"
	"time"
)

const ServiceID = "S3"
const ServiceAPIVersion = "2006-03-01"

// Client provides the API client to make operations call for Amazon Simple Storage
// Service.
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveDefaultLogger(&options)

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveDefaultEndpointConfiguration(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// Configures the events that will be sent to the configured logger.
	ClientLogMode aws.ClientLogMode

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions EndpointResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// The logger writer interface to write logging messages to.
	Logger logging.Logger

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// Allows you to enable arn region support for the service.
	UseARNRegion bool

	// Allows you to enable S3 Accelerate feature. All operations compatible with S3
	// Accelerate will use the accelerate endpoint for requests. Requests not
	// compatible will fall back to normal S3 requests. The bucket must be enabled for
	// accelerate to be used with S3 client with accelerate enabled. If the bucket is
	// not enabled for accelerate an error will be returned. The bucket name must be
	// DNS compatible to work with accelerate.
	UseAccelerate bool

	// Allows you to enable Dualstack endpoint support for the service.
	UseDualstack bool

	// Allows you to enable the client to use path-style addressing, i.e.,
	// https://s3.amazonaws.com/BUCKET/KEY. By default, the S3 client will use virtual
	// hosted bucket addressing when possible(https://BUCKET.s3.amazonaws.com/KEY).
	UsePathStyle bool

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]func(*middleware.Stack) error, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}
func (c *Client) invokeOperation(ctx context.Context, opID string, params interface{}, optFns []func(*Options), stackFns ...func(*middleware.Stack, Options) error) (result interface{}, metadata middleware.Metadata, err error) {
	stack := middleware.NewStack(opID, smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}

	for _, fn := range stackFns {
		if err := fn(stack, options); err != nil {
			return nil, metadata, err
		}
	}

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, metadata, err
		}
	}

	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err = handler.Handle(ctx, params)
	if err != nil {
		err = &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: opID,
			Err:           err,
		}
	}
	return result, metadata, err
}

func resolveDefaultLogger(o *Options) {
	if o.Logger != nil {
		return
	}
	o.Logger = logging.Nop{}
}

func addSetLoggerMiddleware(stack *middleware.Stack, o Options) error {
	return middleware.AddSetLoggerMiddleware(stack, o.Logger)
}

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:        cfg.Region,
		Retryer:       cfg.Retryer,
		HTTPClient:    cfg.HTTPClient,
		Credentials:   cfg.Credentials,
		APIOptions:    cfg.APIOptions,
		Logger:        cfg.Logger,
		ClientLogMode: cfg.ClientLogMode,
	}
	resolveAWSEndpointResolver(cfg, &opts)
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = aws.NewBuildableHTTPClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func resolveAWSEndpointResolver(cfg aws.Config, o *Options) {
	if cfg.EndpointResolver == nil {
		return
	}
	o.EndpointResolver = WithEndpointResolver(cfg.EndpointResolver, NewDefaultEndpointResolver())
}

func addClientUserAgent(stack *middleware.Stack) error {
	return awsmiddleware.AddUserAgentKey("s3")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) error {
	return stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, o.HTTPSignerV4), middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = v4.NewSigner(
		func(s *v4.Signer) {
			s.Logger = o.Logger
			s.LogSigning = o.ClientLogMode.IsSigning()
			s.DisableURIPathEscaping = true
		},
	)
}

func addRetryMiddlewares(stack *middleware.Stack, o Options) error {
	mo := retry.AddRetryMiddlewaresOptions{
		Retryer:          o.Retryer,
		LogRetryAttempts: o.ClientLogMode.IsRetries(),
	}
	return retry.AddRetryMiddlewares(stack, mo)
}

func addMetadataRetrieverMiddleware(stack *middleware.Stack) error {
	return s3shared.AddMetadataRetrieverMiddleware(stack)
}

// getBucketFromInput returns a boolean indicating if the input has a modeled
// bucket name, and a pointer to string denoting a provided bucket member value
func getBucketFromInput(input interface{}) (*string, bool) {
	switch i := input.(type) {
	case *AbortMultipartUploadInput:
		return i.Bucket, true
	case *CompleteMultipartUploadInput:
		return i.Bucket, true
	case *CopyObjectInput:
		return i.Bucket, true
	case *CreateBucketInput:
		return i.Bucket, true
	case *CreateMultipartUploadInput:
		return i.Bucket, true
	case *DeleteBucketInput:
		return i.Bucket, true
	case *DeleteBucketAnalyticsConfigurationInput:
		return i.Bucket, true
	case *DeleteBucketCorsInput:
		return i.Bucket, true
	case *DeleteBucketEncryptionInput:
		return i.Bucket, true
	case *DeleteBucketInventoryConfigurationInput:
		return i.Bucket, true
	case *DeleteBucketLifecycleInput:
		return i.Bucket, true
	case *DeleteBucketMetricsConfigurationInput:
		return i.Bucket, true
	case *DeleteBucketOwnershipControlsInput:
		return i.Bucket, true
	case *DeleteBucketPolicyInput:
		return i.Bucket, true
	case *DeleteBucketReplicationInput:
		return i.Bucket, true
	case *DeleteBucketTaggingInput:
		return i.Bucket, true
	case *DeleteBucketWebsiteInput:
		return i.Bucket, true
	case *DeleteObjectInput:
		return i.Bucket, true
	case *DeleteObjectsInput:
		return i.Bucket, true
	case *DeleteObjectTaggingInput:
		return i.Bucket, true
	case *DeletePublicAccessBlockInput:
		return i.Bucket, true
	case *GetBucketAccelerateConfigurationInput:
		return i.Bucket, true
	case *GetBucketAclInput:
		return i.Bucket, true
	case *GetBucketAnalyticsConfigurationInput:
		return i.Bucket, true
	case *GetBucketCorsInput:
		return i.Bucket, true
	case *GetBucketEncryptionInput:
		return i.Bucket, true
	case *GetBucketInventoryConfigurationInput:
		return i.Bucket, true
	case *GetBucketLifecycleConfigurationInput:
		return i.Bucket, true
	case *GetBucketLoggingInput:
		return i.Bucket, true
	case *GetBucketMetricsConfigurationInput:
		return i.Bucket, true
	case *GetBucketNotificationConfigurationInput:
		return i.Bucket, true
	case *GetBucketOwnershipControlsInput:
		return i.Bucket, true
	case *GetBucketPolicyInput:
		return i.Bucket, true
	case *GetBucketPolicyStatusInput:
		return i.Bucket, true
	case *GetBucketReplicationInput:
		return i.Bucket, true
	case *GetBucketRequestPaymentInput:
		return i.Bucket, true
	case *GetBucketTaggingInput:
		return i.Bucket, true
	case *GetBucketVersioningInput:
		return i.Bucket, true
	case *GetBucketWebsiteInput:
		return i.Bucket, true
	case *GetObjectInput:
		return i.Bucket, true
	case *GetObjectAclInput:
		return i.Bucket, true
	case *GetObjectLegalHoldInput:
		return i.Bucket, true
	case *GetObjectLockConfigurationInput:
		return i.Bucket, true
	case *GetObjectRetentionInput:
		return i.Bucket, true
	case *GetObjectTaggingInput:
		return i.Bucket, true
	case *GetObjectTorrentInput:
		return i.Bucket, true
	case *GetPublicAccessBlockInput:
		return i.Bucket, true
	case *HeadBucketInput:
		return i.Bucket, true
	case *HeadObjectInput:
		return i.Bucket, true
	case *ListBucketAnalyticsConfigurationsInput:
		return i.Bucket, true
	case *ListBucketInventoryConfigurationsInput:
		return i.Bucket, true
	case *ListBucketMetricsConfigurationsInput:
		return i.Bucket, true
	case *ListMultipartUploadsInput:
		return i.Bucket, true
	case *ListObjectsInput:
		return i.Bucket, true
	case *ListObjectsV2Input:
		return i.Bucket, true
	case *ListObjectVersionsInput:
		return i.Bucket, true
	case *ListPartsInput:
		return i.Bucket, true
	case *PutBucketAccelerateConfigurationInput:
		return i.Bucket, true
	case *PutBucketAclInput:
		return i.Bucket, true
	case *PutBucketAnalyticsConfigurationInput:
		return i.Bucket, true
	case *PutBucketCorsInput:
		return i.Bucket, true
	case *PutBucketEncryptionInput:
		return i.Bucket, true
	case *PutBucketInventoryConfigurationInput:
		return i.Bucket, true
	case *PutBucketLifecycleConfigurationInput:
		return i.Bucket, true
	case *PutBucketLoggingInput:
		return i.Bucket, true
	case *PutBucketMetricsConfigurationInput:
		return i.Bucket, true
	case *PutBucketNotificationConfigurationInput:
		return i.Bucket, true
	case *PutBucketOwnershipControlsInput:
		return i.Bucket, true
	case *PutBucketPolicyInput:
		return i.Bucket, true
	case *PutBucketReplicationInput:
		return i.Bucket, true
	case *PutBucketRequestPaymentInput:
		return i.Bucket, true
	case *PutBucketTaggingInput:
		return i.Bucket, true
	case *PutBucketVersioningInput:
		return i.Bucket, true
	case *PutBucketWebsiteInput:
		return i.Bucket, true
	case *PutObjectInput:
		return i.Bucket, true
	case *PutObjectAclInput:
		return i.Bucket, true
	case *PutObjectLegalHoldInput:
		return i.Bucket, true
	case *PutObjectLockConfigurationInput:
		return i.Bucket, true
	case *PutObjectRetentionInput:
		return i.Bucket, true
	case *PutObjectTaggingInput:
		return i.Bucket, true
	case *PutPublicAccessBlockInput:
		return i.Bucket, true
	case *RestoreObjectInput:
		return i.Bucket, true
	case *UploadPartInput:
		return i.Bucket, true
	case *UploadPartCopyInput:
		return i.Bucket, true
	default:
		return nil, false
	}
}

// supportAccelerate returns a boolean indicating if the operation associated with
// the provided input supports S3 Transfer Acceleration
func supportAccelerate(input interface{}) bool {
	switch input.(type) {
	case *CreateBucketInput:
		return false
	case *DeleteBucketInput:
		return false
	case *ListBucketsInput:
		return false
	default:
		return true
	}
}

func addResponseErrorMiddleware(stack *middleware.Stack) error {
	return s3shared.AddResponseErrorMiddleware(stack)
}

func disableAcceptEncodingGzip(stack *middleware.Stack) error {
	return acceptencodingcust.AddAcceptEncodingGzip(stack, acceptencodingcust.AddAcceptEncodingGzipOptions{})
}

func addRequestResponseLogging(stack *middleware.Stack, o Options) error {
	return stack.Deserialize.Add(&smithyhttp.RequestResponseLogger{
		LogRequest:          o.ClientLogMode.IsRequest(),
		LogRequestWithBody:  o.ClientLogMode.IsRequestWithBody(),
		LogResponse:         o.ClientLogMode.IsResponse(),
		LogResponseWithBody: o.ClientLogMode.IsResponseWithBody(),
	}, middleware.After)
}
