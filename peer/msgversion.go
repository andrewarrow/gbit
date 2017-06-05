package peer

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"
)

const MaxUserAgentLen = 256
const DefaultUserAgent = "/gbit:0.1.0/"

type MsgVersion struct {
	ProtocolVersion int32
	Services        ServiceFlag
	Timestamp       time.Time
	AddrYou         NetAddress
	AddrMe          NetAddress
	Nonce           uint64
	UserAgent       string
	LastBlock       int32
	DisableRelayTx  bool
}

/*

func (msg *MsgVersion) HasService(service ServiceFlag) bool {
	return msg.Services&service == service
}

func (msg *MsgVersion) AddService(service ServiceFlag) {
	msg.Services |= service
}

func (msg *MsgVersion) BtcDecode(r io.Reader, pver uint32) error {
	buf, ok := r.(*bytes.Buffer)
	if !ok {
		return fmt.Errorf("MsgVersion.BtcDecode reader is not a " +
			"*bytes.Buffer")
	}

	err := readElements(buf, &msg.ProtocolVersion, &msg.Services,
		(*int64Time)(&msg.Timestamp))
	if err != nil {
		return err
	}

	err = readNetAddress(buf, pver, &msg.AddrYou, false)
	if err != nil {
		return err
	}

	if buf.Len() > 0 {
		err = readNetAddress(buf, pver, &msg.AddrMe, false)
		if err != nil {
			return err
		}
	}
	if buf.Len() > 0 {
		err = readElement(buf, &msg.Nonce)
		if err != nil {
			return err
		}
	}
	if buf.Len() > 0 {
		userAgent, err := ReadVarString(buf, pver)
		if err != nil {
			return err
		}
		err = validateUserAgent(userAgent)
		if err != nil {
			return err
		}
		msg.UserAgent = userAgent
	}

	if buf.Len() > 0 {
		err = readElement(buf, &msg.LastBlock)
		if err != nil {
			return err
		}
	}

	if buf.Len() > 0 {
		var relayTx bool
		readElement(r, &relayTx)
		msg.DisableRelayTx = !relayTx
	}

	return nil
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

func (msg *MsgVersion) Command() string {
	return CmdVersion
}

func (msg *MsgVersion) MaxPayloadLength(pver uint32) uint32 {
	return 33 + (maxNetAddressPayload(pver) * 2) + MaxVarIntPayload +
		MaxUserAgentLen
}

func NewMsgVersion(me *NetAddress, you *NetAddress, nonce uint64,
	lastBlock int32) *MsgVersion {

	return &MsgVersion{
		ProtocolVersion: int32(ProtocolVersion),
		Services:        0,
		Timestamp:       time.Unix(time.Now().Unix(), 0),
		AddrYou:         *you,
		AddrMe:          *me,
		Nonce:           nonce,
		UserAgent:       DefaultUserAgent,
		LastBlock:       lastBlock,
		DisableRelayTx:  false,
	}
}

func validateUserAgent(userAgent string) error {
	if len(userAgent) > MaxUserAgentLen {
		str := fmt.Sprintf("user agent too long [len %v, max %v]",
			len(userAgent), MaxUserAgentLen)
		return messageError("MsgVersion", str)
	}
	return nil
}

func (msg *MsgVersion) AddUserAgent(name string, version string,
	comments ...string) error {

	newUserAgent := fmt.Sprintf("%s:%s", name, version)
	if len(comments) != 0 {
		newUserAgent = fmt.Sprintf("%s(%s)", newUserAgent,
			strings.Join(comments, "; "))
	}
	newUserAgent = fmt.Sprintf("%s%s/", msg.UserAgent, newUserAgent)
	err := validateUserAgent(newUserAgent)
	if err != nil {
		return err
	}
	msg.UserAgent = newUserAgent
	return nil
}
*/
