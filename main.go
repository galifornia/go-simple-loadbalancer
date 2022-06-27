package main

import (
	"net/http"

	"github.com/galifornia/go-simple-loadbalancer/lib"
)

var loadbalancer *lib.LoadBalancer

func main() {
	server1, _ := lib.NewSimpleServer("https://martingl.com")
	server2, _ := lib.NewSimpleServer("https://duckduckgo.com")
	server3, _ := lib.NewSimpleServer("https://www.elliberal.com")

	servers := []lib.Server{
		server1,
		server2,
		server3,
	}
	loadbalancer = lib.NewLoadBalancer("8030", servers)
	http.HandleFunc("/", loadbalancer.ServerProxy)

	http.ListenAndServe("localhost:"+loadbalancer.Port, nil)
}
