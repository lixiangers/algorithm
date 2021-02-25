package flag2

import (
	"flag"
	"fmt"
)

var (
	listenAddr = flag.String("listen-addr", "", "listen addr")
	etcdAddrs  = flag.String("etcd-addrs", "", "etcd addr list, such as \"ip1:port1,ip2:port2...\"")
	namespace  = flag.String("namespace", "", "etcd namespace, such as \"lx.p.1.bjdt\" for platform, or \"lx.c.bjdt\" for center")
)

func LxMain() {
	flag.Parse()
	fmt.Printf("listenAddr:%s\netcdAddrs:%s\nnamespace:%s\n", *listenAddr, *etcdAddrs, *namespace)
}
