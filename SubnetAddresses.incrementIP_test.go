package subnetgenerator

import (
	"testing"
)

func TestSubnetAddresses_incrementIP(t *testing.T) {
	t.Run("Test case 1: Increment within the subnet mask", func(t *testing.T) {
		cidr := "192.168.1.0/29"
		expectedIP := "192.168.1.5"
		subnet, err := NewSubnet(cidr)
		if err != nil {
			t.Fatalf("Failed to create subnet: %v", err)
		}
		// Increment IP a few times
		for i := 0; i < 5; i++ {
			subnet.incrementIP()
		}
		// Check the incremented IP
		if subnet.ip.String() != expectedIP {
			t.Errorf("Expected IP %s, but got %s", expectedIP, subnet.ip.String())
		}
	})

	t.Run("Test case 2: Increment at subnet mask boundary", func(t *testing.T) {
		cidr := "10.0.0.0/24"
		subnet, err := NewSubnet(cidr)
		expectedIP := "10.0.1.0"
		if err != nil {
			t.Fatalf("Failed to create subnet: %v", err)
		}

		// Increment IP until overflow to next subnet
		for i := 0; i < 256; i++ {
			subnet.incrementIP()
		}

		// Check the incremented IP

		if subnet.ip.String() != expectedIP {
			t.Errorf("Expected IP %s, but got %s", expectedIP, subnet.ip.String())
		}
	})
}
