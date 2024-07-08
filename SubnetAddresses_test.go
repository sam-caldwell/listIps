package subnetgenerator

import (
	"bytes"
	"net"
	"testing"
)

func TestNewSubnetAddresses(t *testing.T) {
	var s SubnetAddresses
	if s.ipNet != nil {
		t.Fatalf("expect nil")
	}
	s.ip = net.IP{0x01, 0x02, 0x03, 0x04}
	if !bytes.Equal(s.ip, []byte{0x01, 0x02, 0x03, 0x04}) {
		t.Fatal("mismatch")
	}
}
