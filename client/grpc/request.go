package grpc

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/micro/go-micro/client"
)

type grpcRequest struct {
	service     string
	method      string
	contentType string
	request     interface{}
	opts        client.RequestOptions
}

func methodToGRPC(service, method string, request interface{}) string {
	// no method or already grpc method
	if len(method) == 0 || method[0] == '/' {
		return method
	}
	// can't operate on nil request
	t := reflect.TypeOf(request)
	if t == nil {
		return method
	}
	// dereference
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// assume method is Foo.Bar
	mParts := strings.Split(method, ".")
	if len(mParts) != 2 {
		return method
	}
	// return /pkg.Foo/Bar
	return fmt.Sprintf("/%s.%s/%s", service, mParts[0], mParts[1])
}

func newGRPCRequest(service, method string, request interface{}, contentType string, reqOpts ...client.RequestOption) client.Request {
	var opts client.RequestOptions
	for _, o := range reqOpts {
		o(&opts)
	}

	return &grpcRequest{
		service:     service,
		method:      method,
		request:     request,
		contentType: contentType,
		opts:        opts,
	}
}

func (g *grpcRequest) ContentType() string {
	return g.contentType
}

func (g *grpcRequest) Service() string {
	return g.service
}

func (g *grpcRequest) Method() string {
	return g.method
}

func (g *grpcRequest) Request() interface{} {
	return g.request
}

func (g *grpcRequest) Stream() bool {
	return g.opts.Stream
}
