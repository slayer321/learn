package main

import (
	"net/http"
	"net/http/httputil"
	"reflect"
	"testing"
)

func Test_loadbalancer_getNextAvailableServer(t *testing.T) {
	type fields struct {
		port            string
		roundRobinCount int
		servers         []Server
	}
	tests := []struct {
		name   string
		fields fields
		want   Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadbalancer{
				port:            tt.fields.port,
				roundRobinCount: tt.fields.roundRobinCount,
				servers:         tt.fields.servers,
			}
			if got := lb.getNextAvailableServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadbalancer.getNextAvailableServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleErr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleErr(tt.args.err)
		})
	}
}

func Test_newsimpleServer(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want *simpleServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newsimpleServer(tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newsimpleServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLoadBalancer(t *testing.T) {
	type args struct {
		port    string
		servers []Server
	}
	tests := []struct {
		name string
		args args
		want *loadbalancer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoadBalancer(tt.args.port, tt.args.servers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoadBalancer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simpleServer_Address(t *testing.T) {
	type fields struct {
		addr  string
		proxy *httputil.ReverseProxy
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &simpleServer{
				addr:  tt.fields.addr,
				proxy: tt.fields.proxy,
			}
			if got := s.Address(); got != tt.want {
				t.Errorf("simpleServer.Address() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simpleServer_IsAlive(t *testing.T) {
	type fields struct {
		addr  string
		proxy *httputil.ReverseProxy
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &simpleServer{
				addr:  tt.fields.addr,
				proxy: tt.fields.proxy,
			}
			if got := s.IsAlive(); got != tt.want {
				t.Errorf("simpleServer.IsAlive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simpleServer_Serve(t *testing.T) {
	type fields struct {
		addr  string
		proxy *httputil.ReverseProxy
	}
	type args struct {
		rw  http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &simpleServer{
				addr:  tt.fields.addr,
				proxy: tt.fields.proxy,
			}
			s.Serve(tt.args.rw, tt.args.req)
		})
	}
}

func Test_loadbalancer_serverProxy(t *testing.T) {
	type fields struct {
		port            string
		roundRobinCount int
		servers         []Server
	}
	type args struct {
		rw  http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadbalancer{
				port:            tt.fields.port,
				roundRobinCount: tt.fields.roundRobinCount,
				servers:         tt.fields.servers,
			}
			lb.serverProxy(tt.args.rw, tt.args.req)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
