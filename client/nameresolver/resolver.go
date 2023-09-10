package nameresolver

import (
	"fmt"

	"google.golang.org/grpc/resolver"
)

/**
 * Name resolving
 * A name resolver can be seen as a map[service-name][]backend-ip. It takes a service name, and returns a list of IPs of the backends.
 */

const (
	scheme = "example"
	name   = "lb.example.grpc.io"
)

// BuildURI returns the URI to be used in grpc Dial.
func BuildURI() string {
	return fmt.Sprintf("%s:///%s", scheme, name)
}

type ResolverBuilder struct {
	addrs []string
}

// NewBuilder creates a Resolver Builder which is used in grpc Dial to create a resolver.
func NewBuilder(addrs []string) resolver.Builder {
	return &ResolverBuilder{addrs: addrs}
}

func (rb *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &myResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			name: rb.addrs,
		},
	}
	r.start()
	return r, nil
}
func (*ResolverBuilder) Scheme() string { return scheme }

type myResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *myResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint()]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
func (*myResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (*myResolver) Close()                                  {}
