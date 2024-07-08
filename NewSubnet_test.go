package subnetgenerator

import (
	"bytes"
	"net"
	"testing"
)

func TestNewSubnet(t *testing.T) {
	// Test case 1: Valid CIDR
	cidr := "192.168.1.0/24"
	expectedIP := net.IPv4(192, 168, 1, 0)
	expectedMask := net.IPv4Mask(255, 255, 255, 0)

	subnet, err := NewSubnet(cidr)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if !subnet.ipNet.IP.Equal(expectedIP) {
		t.Errorf("Expected IP %v, but got %v", expectedIP, subnet.ipNet.IP)
	}
	if !bytes.Equal(subnet.ipNet.Mask, expectedMask) {
		t.Errorf("Expected mask %v, but got %v", expectedMask, subnet.ipNet.Mask)
	}

	// Test case 2: Invalid CIDR
	invalidCIDR := "192.168.1.0/33"
	_, err = NewSubnet(invalidCIDR)
	if err == nil {
		t.Errorf("Expected error for invalid CIDR %s, but got nil", invalidCIDR)
	}
}
