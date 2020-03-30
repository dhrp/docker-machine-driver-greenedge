package main

// This file is simply to build the rancher openstack driver.
// it's for the latest openstack (train) as we're using at GreenEdge
// It is required because the openstack driver is otherwise included
// by default in the docker-machine binary.
// Simply put it somewhere, and build it with
// go build -o docker-machine-driver-greenedge and then use it like
// docker-machine create --driver greenedge [etc]

import (
	"github.com/rancher/machine/drivers/openstack"
	"github.com/rancher/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(openstack.NewDriver("", ""))
}
