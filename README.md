# docker-machine-driver-greenedge

This file is simply to build the rancher openstack driver.
it's for the latest openstack (train) as we're using at GreenEdge

It is required because the openstack driver is otherwise included
by default in the docker-machine binary.
Simply put it somewhere, and build it with
go build -o docker-machine-driver-greenedge and then use it like
docker-machine create --driver greenedge [etc]
