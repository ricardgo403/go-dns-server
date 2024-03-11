package main

import (
	"dns-resolver/services"
	"fmt"

	"github.com/miekg/dns"
)

func main() {
	handler := services.NewDefaultDNSHandler()
	server := &dns.Server{
		Addr:      ":53",
		Net:       "udp",
		Handler:   handler,
		UDPSize:   65535,
		ReusePort: true,
	}

	fmt.Println("Starting DNS server on port 53")

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
	}
}
