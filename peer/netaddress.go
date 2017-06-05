package peer

import (
	"encoding/binary"
	"io"
	"net"
	"time"
)

const (
	NetAddressTimeVersion uint32 = 31402
)

type ServiceFlag uint64

func maxNetAddressPayload(pver uint32) uint32 {
	plen := uint32(26)

	if pver >= NetAddressTimeVersion {
		plen += 4
	}

	return plen
}

type NetAddress struct {
	Timestamp time.Time
	Services  ServiceFlag
	IP        net.IP
	Port      uint16
}

/*
func (na *NetAddress) HasService(service ServiceFlag) bool {
	return na.Services&service == service
}

func (na *NetAddress) AddService(service ServiceFlag) {
	na.Services |= service
}

func NewNetAddressIPPort(ip net.IP, port uint16, services ServiceFlag) *NetAddress {
	return NewNetAddressTimestamp(time.Now(), services, ip, port)
}

func NewNetAddressTimestamp(
	timestamp time.Time, services ServiceFlag, ip net.IP, port uint16) *NetAddress {
	na := NetAddress{
		Timestamp: time.Unix(timestamp.Unix(), 0),
		Services:  services,
		IP:        ip,
		Port:      port,
	}
	return &na
}

func NewNetAddress(addr *net.TCPAddr, services ServiceFlag) *NetAddress {
	return NewNetAddressIPPort(addr.IP, uint16(addr.Port), services)
}

func readNetAddress(r io.Reader, pver uint32, na *NetAddress, ts bool) error {
	var ip [16]byte

	if ts && pver >= NetAddressTimeVersion {
		err := readElement(r, (*uint32Time)(&na.Timestamp))
		if err != nil {
			return err
		}
	}

	err := readElements(r, &na.Services, &ip)
	if err != nil {
		return err
	}
	port, err := binarySerializer.Uint16(r, bigEndian)
	if err != nil {
		return err
	}

	*na = NetAddress{
		Timestamp: na.Timestamp,
		Services:  na.Services,
		IP:        net.IP(ip[:]),
		Port:      port,
	}
	return nil
}

func writeNetAddress(w io.Writer, pver uint32, na *NetAddress, ts bool) error {
	if ts && pver >= NetAddressTimeVersion {
		err := writeElement(w, uint32(na.Timestamp.Unix()))
		if err != nil {
			return err
		}
	}

	var ip [16]byte
	if na.IP != nil {
		copy(ip[:], na.IP.To16())
	}
	err := writeElements(w, na.Services, ip)
	if err != nil {
		return err
	}

	return binary.Write(w, bigEndian, na.Port)
}*/
