package lib

import "net/http"

type LoadBalancer struct {
	Port            string
	RoundRobinCount int
	Servers         []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		Port:            port,
		RoundRobinCount: 0,
		Servers:         servers,
	}
}

func (lb *LoadBalancer) AddServerToLoadBalancer(server Server) {
	lb.Servers = append(lb.Servers, server)
}

func (lb *LoadBalancer) GetNextAvailableServer() Server {
	serverCount := len(lb.Servers)
	next := lb.RoundRobinCount
	server := lb.Servers[next]

	// keep round until server is found
	for !server.IsAlive() {
		lb.RoundRobinCount = (lb.RoundRobinCount + 1) % serverCount
		next = lb.RoundRobinCount
		server = lb.Servers[next]
	}

	// bump count for next server
	lb.RoundRobinCount = (lb.RoundRobinCount + 1) % serverCount

	return lb.Servers[next]
}

func (lb *LoadBalancer) ServerProxy(rw http.ResponseWriter, r *http.Request) {
	server := lb.GetNextAvailableServer()
	server.Serve(rw, r)
}
