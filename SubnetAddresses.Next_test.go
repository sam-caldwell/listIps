package subnetgenerator

import (
	"testing"
)

func TestSubnetAddresses_Next(t *testing.T) {
	cidr := "192.168.1.0/29"
	subnet, err := NewSubnet(cidr)
	if err != nil {
		t.Fatalf("Failed to create subnet: %v", err)
	}

	expectedIPs := []string{
		"192.168.1.0",
		"192.168.1.1",
		"192.168.1.2",
		"192.168.1.3",
		"192.168.1.4",
		"192.168.1.5",
		"192.168.1.6",
		"192.168.1.7",
	}

	var generatedIPs []string
	for {
		hasNext, ip := subnet.Next()
		if !hasNext {
			break
		}
		generatedIPs = append(generatedIPs, ip)
	}

	if !compareStringSlices(generatedIPs, expectedIPs) {
		t.Errorf("mismatch\n"+
			"actual   : %v\n"+
			"expected : %v",
			generatedIPs, expectedIPs)
	}
}

// Utility function to compare two string slices
func compareStringSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
