package peer

import "fmt"
import "net"
import "bytes"

//import "github.com/davecgh/go-spew/spew"

func Hello(ip net.IP) bool {
	//fmt.Println(" ", ip)
	conn, err := net.DialTimeout("tcp", "["+ip.String()+"]:8333", 100000000)
	if err == nil {
		fmt.Println(" ", conn, err)
		return true
	}
	return false
}

type BitcoinNet uint32

const CommandSize = 12
const BIP0031Version uint32 = 60000
const ProtocolVersion uint32 = 70013
const MainNet BitcoinNet = 0xd9b4bef9

func WriteMessage() error {
	WriteMessage(p, "version")
	n, err := wire.WriteMessageN(p.conn, msg, ProtocolVersion, MainNet)
}

type messageHeader struct {
	magic    BitcoinNet // 4 bytes
	command  string     // 12 bytes
	length   uint32     // 4 bytes
	checksum [4]byte    // 4 bytes
}

func DoubleHashB(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

func WriteMessage(w io.Writer, cmd string, pver uint32, btcnet BitcoinNet) (int, error) {
	totalBytes := 0

	var command [CommandSize]byte
	copy(command[:], []byte(cmd))

	var bw bytes.Buffer
	err := msg.BtcEncode(&bw, pver)
	if err != nil {
		return totalBytes, err
	}
	payload := bw.Bytes()
	lenp := len(payload)

	hdr := messageHeader{}
	hdr.magic = btcnet
	hdr.command = cmd
	hdr.length = uint32(lenp)
	copy(hdr.checksum[:], DoubleHashB(payload)[0:4])

	hw := bytes.NewBuffer(make([]byte, 0, MessageHeaderSize))
	writeElements(hw, hdr.magic, command, hdr.length, hdr.checksum)

	n, err := w.Write(hw.Bytes())
	totalBytes += n
	if err != nil {
		return totalBytes, err
	}

	n, err = w.Write(payload)
	totalBytes += n
	return totalBytes, err
}

func (msg *MsgVersion) BtcEncode(w io.Writer, pver uint32) error {
	err := validateUserAgent(msg.UserAgent)
	if err != nil {
		return err
	}

	err = writeElements(w, msg.ProtocolVersion, msg.Services,
		msg.Timestamp.Unix())
	if err != nil {
		return err
	}

	err = writeNetAddress(w, pver, &msg.AddrYou, false)
	if err != nil {
		return err
	}

	err = writeNetAddress(w, pver, &msg.AddrMe, false)
	if err != nil {
		return err
	}

	err = writeElement(w, msg.Nonce)
	if err != nil {
		return err
	}

	err = WriteVarString(w, pver, msg.UserAgent)
	if err != nil {
		return err
	}

	err = writeElement(w, msg.LastBlock)
	if err != nil {
		return err
	}

	if pver >= BIP0037Version {
		err = writeElement(w, !msg.DisableRelayTx)
		if err != nil {
			return err
		}
	}
	return nil
}
