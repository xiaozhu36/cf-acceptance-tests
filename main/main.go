package main

import (
	"github.com/cloudfoundry/cf-acceptance-tests/tcp_routing"
	"fmt"
)

func main() {
	resp, err := tcp_routing.SendAndReceive("tcp.cloudfoundry.vip", "1025")
	fmt.Printf("**** resp: %#v \n ***** err: %#v", resp, err)
}

