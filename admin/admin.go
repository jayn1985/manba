package main

import (
	"flag"
	"github.com/fagongzi/gateway/admin/pkg/server"
	"runtime"
)

var (
	cpus       = flag.Int("cpus", 1, "use cpu nums")
	addr       = flag.String("addr", ":8080", "listen addr.(e.g. ip:port)")
	etcdAddr   = flag.String("etcd-addr", "http://127.0.0.1:2379", "etcd address, use ',' to splite.")
	etcdPrefix = flag.String("etcd-prefix", "/gateway", "etcd node prefix.")
)

var (
	userName = flag.String("user", "admin", "admin user name")
	pwd      = flag.String("pwd", "admin", "admin user pwd")
)

func main() {
	flag.Parse()

	runtime.GOMAXPROCS(*cpus)

	address := []string{*etcdAddr}
	s := server.NewAdminServer(*addr, address, *etcdPrefix, *userName, *pwd)
	s.Start()
}
