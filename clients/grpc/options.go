package grpc

import (
	"crypto/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Opt func(c *Option)

type Option struct {
	Uri     string
	GrpcOpt []grpc.DialOption
}

func DefaultOptions() *Option {
	opt := &Option{}
	opt.GrpcOpt = append(opt.GrpcOpt, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	return opt
}

func WithURI(target string) Opt {
	return func(o *Option) {
		o.Uri = target
	}
}

func WithCredentials() Opt {
	return func(c *Option) {
		c.GrpcOpt = append(c.GrpcOpt, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	}
}

func WithDisablesCredentials() Opt {
	return func(c *Option) {
		c.GrpcOpt = append(c.GrpcOpt, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
}

func WithPerRPCCredentials(creds credentials.PerRPCCredentials) Opt {
	return func(c *Option) {
		c.GrpcOpt = append(c.GrpcOpt, grpc.WithPerRPCCredentials(creds))
	}
}

func WithTls(creds credentials.TransportCredentials) Opt {
	return func(c *Option) {
		c.GrpcOpt = append(c.GrpcOpt, grpc.WithTransportCredentials(creds))
	}
}

func WithDisableTls() Opt {
	return func(c *Option) {
		c.GrpcOpt = append(c.GrpcOpt, grpc.WithInsecure())
	}
}
