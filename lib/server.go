package lib

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

type SimpleServer struct {
	Addr  string
	Proxy *httputil.ReverseProxy
}

func (s *SimpleServer) Address() string {
	return s.Addr
}

func (s *SimpleServer) IsAlive() bool {
	return true // !FIXME
}

func (s *SimpleServer) Serve(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving server", s.Addr)
	s.Proxy.ServeHTTP(rw, r)
}

func NewSimpleServer(addr string) (*SimpleServer, error) {
	url, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	return &SimpleServer{
		Addr:  addr,
		Proxy: httputil.NewSingleHostReverseProxy(url),
	}, nil
}
