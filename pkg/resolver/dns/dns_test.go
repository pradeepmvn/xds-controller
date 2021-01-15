package dns_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pradeepmvn/xds-controller/pkg/config"
	"github.com/pradeepmvn/xds-controller/pkg/log"
	"github.com/pradeepmvn/xds-controller/pkg/resolver"
	"github.com/pradeepmvn/xds-controller/pkg/resolver/dns"
)

func TestGetEndPoints(t *testing.T) {
	log.NewLogger(true)
	cc := &config.Cluster{
		Name:         "GoogleTest",
		ResolverType: "dns",
		Address:      "www.google.com",
	}
	r := dns.New(cc)
	d := r.GetEndPoints()
	fmt.Println("Inital List", d)

}

func TestWatcher(t *testing.T) {
	log.NewLogger(true)
	cc := &config.Cluster{
		Name:                 "GoogleTest",
		ResolverType:         "dns",
		Address:              "www.google.com",
		RefreshIntervalInSec: 3,
	}
	r := dns.New(cc)
	d := r.GetEndPoints()

	fmt.Println("Inital List", d)

	go func(r resolver.Resolver) {
		//kill after 5 sec
		time.Sleep(5 * time.Second)s
		r.Close()
		fmt.Println("Stopping")
	}(r)
	// recieve refreshes
	go func(r resolver.Resolver) {
		for range r.Refresh() {
			fmt.Println("New List", r.Latest())
		}
	}(r)

	// start watcher
	r.Watch()
}
