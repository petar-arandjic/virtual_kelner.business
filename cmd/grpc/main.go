package main

import "github.com/petar-arandjic/virtual_kelner.business/pkg/grpc"

func main() {
	grpc.Start(8083)
}
