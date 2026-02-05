package network

import (
	"fmt"
	"net"
	"net/url"
)

func RunIPscanner(rawUrl string) {

	url, err := url.Parse(rawUrl)
	hostname := rawUrl
	if err == nil && url.Hostname() != "" {
		hostname = url.Hostname()
	}

	ips, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Printf("[!] Error: Can't find IP for %s\n", hostname)
		return
	}

	for _, ip := range ips {
		fmt.Printf("[+] IP Address: %s\n", ip.String())
	}
}