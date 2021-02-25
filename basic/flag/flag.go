package main

import (
	"flag"
	"fmt"
	"github.com/lixiangers/algorithm/basic/flag/flag2"
)

var (
	test = false
)

func main() {
	if test {
		flag2.LxMain()
	} else {
		listenAddr := flag.String("listen-addr2", "", "listen addr")
		etcdAddrs := flag.String("etcd-addrs2", "", "etcd addr list, such as \"ip1:port1,ip2:port2...\"")
		flag.Parse()
		fmt.Printf("listenAddr:%s\netcdAddrs:%s\n", *listenAddr, *etcdAddrs)
	}
	// painc flag redefined: listen-addr

}
