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
	for _, seed := range Seeds {
		ips, _ := net.LookupIP(seed)
		fmt.Println("")

		for _, ip := range ips {
			peer.Hello(ip)
		}
	}
}
