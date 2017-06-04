package peer

import "fmt"
import "net"

func Hello(ip net.IP) {
	//fmt.Println(" ", ip)
	conn, err := net.DialTimeout("tcp", "["+ip.String()+"]:8333", 10000)
	if err == nil {
		fmt.Println(" ", conn, err)
	}
}
