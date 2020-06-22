// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithytime "github.com/awslabs/smithy-go/time"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestClient_AllQueryStringTypes_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *AllQueryStringTypesInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes query string parameters with all supported types
		"RestJsonAllQueryStringTypes": {
			Params: &AllQueryStringTypesInput{
				QueryString: ptr.String("Hello there"),
				QueryStringList: []*string{
					ptr.String("a"),
					ptr.String("b"),
					ptr.String("c"),
				},
				QueryStringSet: []*string{
					ptr.String("a"),
					ptr.String("b"),
					ptr.String("c"),
				},
				QueryByte:    ptr.Int8(1),
				QueryShort:   ptr.Int16(2),
				QueryInteger: ptr.Int32(3),
				QueryIntegerList: []*int32{
					ptr.Int32(1),
					ptr.Int32(2),
					ptr.Int32(3),
				},
				QueryIntegerSet: []*int32{
					ptr.Int32(1),
					ptr.Int32(2),
					ptr.Int32(3),
				},
				QueryLong:   ptr.Int64(4),
				QueryFloat:  ptr.Float32(1),
				QueryDouble: ptr.Float64(1),
				QueryDoubleList: []*float64{
					ptr.Float64(1.0),
					ptr.Float64(2.0),
					ptr.Float64(3.0),
				},
				QueryBoolean: ptr.Bool(true),
				QueryBooleanList: []*bool{
					ptr.Bool(true),
					ptr.Bool(false),
					ptr.Bool(true),
				},
				QueryTimestamp: ptr.Time(smithytime.ParseEpochSeconds(1)),
				QueryTimestampList: []*time.Time{
					ptr.Time(smithytime.ParseEpochSeconds(1)),
					ptr.Time(smithytime.ParseEpochSeconds(2)),
					ptr.Time(smithytime.ParseEpochSeconds(3)),
				},
				QueryEnum: "Foo",
				QueryEnumList: []types.FooEnum{
					"Foo",
					"Baz",
					"Bar",
				},
			},
			ExpectMethod:  "GET",
			ExpectURIPath: "/AllQueryStringTypesInput",
			ExpectQuery: []smithytesting.QueryItem{
				{Key: "String", Value: "Hello%20there"},
				{Key: "StringList", Value: "a"},
				{Key: "StringList", Value: "b"},
				{Key: "StringList", Value: "c"},
				{Key: "StringSet", Value: "a"},
				{Key: "StringSet", Value: "b"},
				{Key: "StringSet", Value: "c"},
				{Key: "Byte", Value: "1"},
				{Key: "Short", Value: "2"},
				{Key: "Integer", Value: "3"},
				{Key: "IntegerList", Value: "1"},
				{Key: "IntegerList", Value: "2"},
				{Key: "IntegerList", Value: "3"},
				{Key: "IntegerSet", Value: "1"},
				{Key: "IntegerSet", Value: "2"},
				{Key: "IntegerSet", Value: "3"},
				{Key: "Long", Value: "4"},
				{Key: "Float", Value: "1"},
				{Key: "Double", Value: "1"},
				{Key: "DoubleList", Value: "1.0"},
				{Key: "DoubleList", Value: "2.0"},
				{Key: "DoubleList", Value: "3.0"},
				{Key: "Boolean", Value: "true"},
				{Key: "BooleanList", Value: "true"},
				{Key: "BooleanList", Value: "false"},
				{Key: "BooleanList", Value: "true"},
				{Key: "Timestamp", Value: "1"},
				{Key: "TimestampList", Value: "1970-01-01T00%3A00%3A01Z"},
				{Key: "TimestampList", Value: "1970-01-01T00%3A00%3A02Z"},
				{Key: "TimestampList", Value: "1970-01-01T00%3A00%3A03Z"},
				{Key: "Enum", Value: "Foo"},
				{Key: "EnumList", Value: "Foo"},
				{Key: "EnumList", Value: "Baz"},
				{Key: "EnumList", Value: "Bar"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					actualReq = r
					return &http.Response{
						StatusCode: 200,
						Header:     http.Header{},
						Body:       ioutil.NopCloser(strings.NewReader("")),
					}, nil
				}),
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Build.Clear()
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = "https://127.0.0.1"
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2"})
			result, err := client.AllQueryStringTypes(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.Path; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if actualReq.Body != nil {
				defer actualReq.Body.Close()
			}
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}
