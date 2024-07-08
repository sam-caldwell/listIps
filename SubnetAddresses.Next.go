package subnetgenerator

// Next returns the next IP address in the CIDR block along with a boolean indicating if there are more IPs.
func (sa *SubnetAddresses) Next() (bool, string) {
	if sa.ipNet.Contains(sa.ip) {
		thisIp := sa.ip.String()
		sa.incrementIP()
		return true, thisIp
	}
	return false, ""
}
