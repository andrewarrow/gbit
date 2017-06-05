package main

import "fmt"
import "net"
import "github.com/andrewarrow/gbit/peer"

var Seeds = []string{
	"seed.bitcoin.sipa.be",
	"dnsseed.bluematt.me",
	"dnsseed.bitcoin.dashjr.org",
	"seed.bitcoinstats.com",
	"seed.bitnodes.io",
	"seed.bitcoin.jonasschnelli.ch",
}

func main() {
	b := false
	for _, seed := range Seeds {
		ips, _ := net.LookupIP(seed)
		fmt.Println("")

		for _, ip := range ips {
			b = peer.Hello(ip)
			if b {
				break
			}
		}
		if b {
			break
		}
	}
}
