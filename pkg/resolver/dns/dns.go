package dns

import (
	"context"
	"net"
	"strings"
	"time"

	"github.com/pradeepmvn/xds-controller/pkg/config"
	"github.com/pradeepmvn/xds-controller/pkg/log"
	"github.com/pradeepmvn/xds-controller/pkg/resolver"
)

const (
	// Type used in config.
	Type = "dns"
)

type Dnsr struct {
	refresh chan bool
	c       *config.Cluster
	latest  []string
	stop    chan struct{}
}

// Entry point to start resolving a new DNS
func New(c *config.Cluster) resolver.Resolver {
	return &Dnsr{c: c, refresh: make(chan bool), stop: make(chan struct{}, 1)}
}

// GetEndPoints Resolves to ip address with a dns_test
func (d *Dnsr) GetEndPoints() []string {
	var netResolver = net.DefaultResolver
	var err error
	d.latest, err = netResolver.LookupHost(context.Background(), d.c.Address)
	if err != nil {
		log.Error.Panic("Unable to fetch Endpoints. Panicing for now..", err)
	}
	return d.latest
}

// `A` records in the dns should be the same order and comparision can be as simple as `==`
// Watch updates latest on dnsr and triggers a bool chan for refresh.
func (d *Dnsr) Watch() {
watcher:
	for {
		select {
		case <-d.stop:
			log.Debug.Println("Stopping the resolver")
			close(d.refresh)
			break watcher
		default:
			var s []string
			var err error
			s, err = net.DefaultResolver.LookupHost(context.Background(), d.c.Address)
			if err != nil {
				// log the error as error and continue the watcher
				log.Error.Printf("Error while resolving DNS for %s", d.c.Address)
				log.Error.Println(err)
				time.Sleep(time.Duration(d.c.RefreshIntervalInSec) * time.Second)
				continue
			}
			// compare records to previous result
			if (len(s) != len(d.latest)) ||
				!strings.EqualFold(strings.Join(s, ""), strings.Join(d.latest, "")) {
				//length is same. join the strings and see if they are equal to latest..
				log.Debug.Println("Found change..Trigger Refresh")
				d.latest = s
				d.refresh <- true
			}
			time.Sleep(time.Duration(d.c.RefreshIntervalInSec) * time.Second)
		}
	}
}

// Refresh tells when to refresh the data
func (d *Dnsr) Refresh() <-chan bool {
	return d.refresh
}

// Close stops the recolver
func (d *Dnsr) Close() {
	d.stop <- struct{}{}
}

// Latest returns the latest when watcher triggers a refresh
func (d *Dnsr) Latest() []string {
	return d.latest
}
