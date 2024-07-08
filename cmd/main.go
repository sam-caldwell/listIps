package main

import (
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/arg"
	"github.com/sam-caldwell/exit"
	subnetgenerator "github.com/sam-caldwell/lsIp"
)

func main() {
	var (
		err    error
		cidr   *arg.Cidr
		subnet *subnetgenerator.SubnetAddresses
	)

	defer ansi.Reset()

	if cidr, err = arg.NewCidr("cidr", "127.0.0.1/29", "Cidr Address"); err != nil {
		ansi.Red().Printf(err.Error()).Fatal(exit.InvalidInput)
	}

	if err = cidr.Verify(); err != nil {
		ansi.Red().Printf(err.Error()).Fatal(exit.InvalidInput)
	}

	if subnet, err = subnetgenerator.NewSubnet(*cidr.Value()); err != nil {
		ansi.Red().Printf(err.Error()).LF().Fatal(exit.GeneralError).Reset()
	}

	for {
		hasIp, ip := subnet.Next()
		if !hasIp {
			break
		}
		ansi.Println(ip)
	}
}
