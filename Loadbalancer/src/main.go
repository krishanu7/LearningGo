package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleError(err)

	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func (s *simpleServer) Address() string {
	return s.addr
}

func (s *simpleServer) IsAlive() bool {	
	return true
}

func (s *simpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func newLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	for i := 0; i < len(lb.servers); i++ {
		server := lb.servers[lb.roundRobinCount]
		lb.roundRobinCount = (lb.roundRobinCount + 1) % len(lb.servers)
		if server.IsAlive() {
			return server
		}
	}

	return nil
}

func (lb *LoadBalancer) serveProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("Redirecting request to %s\n", targetServer.Address())
	if targetServer != nil {
		targetServer.Serve(w, r)
	} else {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	servers := []Server{
		newSimpleServer("http://localhost:8081"),
		newSimpleServer("http://localhost:8082"),
		newSimpleServer("http://localhost:8083"),
	}
	lb := newLoadBalancer("8080", servers)
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.serveProxy(w, r)
	}
	http.HandleFunc("/", handleRedirect);
	fmt.Printf("Load balancer listening on port %s\n", lb.port)

	err := http.ListenAndServe(":"+lb.port, nil)
	handleError(err)
}
