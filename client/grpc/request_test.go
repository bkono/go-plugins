package grpc

import (
	"testing"

	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func TestMethodToGRPC(t *testing.T) {
	testData := []struct {
		service string
		method  string
		expect  string
		request interface{}
	}{
		{
			"helloworld",
			"Greeter.SayHello",
			"/helloworld.Greeter/SayHello",
			new(pb.HelloRequest),
		},
		{
			"helloworld",
			"/helloworld.Greeter/SayHello",
			"/helloworld.Greeter/SayHello",
			new(pb.HelloRequest),
		},
		{
			"helloworld",
			"Greeter.SayHello",
			"/helloworld.Greeter/SayHello",
			pb.HelloRequest{},
		},
		{
			"helloworld",
			"/helloworld.Greeter/SayHello",
			"/helloworld.Greeter/SayHello",
			pb.HelloRequest{},
		},
		{
			"",
			"Greeter.SayHello",
			"Greeter.SayHello",
			nil,
		},
	}

	for _, d := range testData {
		method := methodToGRPC(d.service, d.method, d.request)
		if method != d.expect {
			t.Fatalf("expected %s got %s", d.expect, method)
		}
	}
}
