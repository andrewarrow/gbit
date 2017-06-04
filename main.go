package main

import "fmt"
import "net"

var Seeds = []string{
	"seed.bitcoin.sipa.be",
	"dnsseed.bluematt.me",
	"dnsseed.bitcoin.dashjr.org",
	"seed.bitcoinstats.com",
	"seed.bitnodes.io",
	"seed.bitcoin.jonasschnelli.ch",
}

func main() {
	ips, _ := net.LookupIP(Seeds[0])
	fmt.Println("gbit <?>", ips)
}
