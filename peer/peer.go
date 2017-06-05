package peer

import "fmt"
import "net"

func Hello(ip net.IP) bool {
	//fmt.Println(" ", ip)
	conn, err := net.DialTimeout("tcp", "["+ip.String()+"]:8333", 100000000)
	if err == nil {
		fmt.Println(" ", conn, err)
		return true
	}
	return false
}
