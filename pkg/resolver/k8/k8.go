// Depends on k8 APIs to resolve for endpoints
// Uses K8 Watcher to detect changes and triggers back to refresh endpoint

package k8

import (
	"context"

	"github.com/pradeepmvn/xds-controller/pkg/config"
	"github.com/pradeepmvn/xds-controller/pkg/log"
	"github.com/pradeepmvn/xds-controller/pkg/resolver"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

const (
	// Type used in config.
	Type = "k8"
)

type K8r struct {
	wi        watch.Interface
	refresh   chan bool
	c         *config.Cluster
	clientset kubernetes.Interface
	latest    []string
}

func New(c *config.Cluster, clientset kubernetes.Interface) resolver.Resolver {
	return &K8r{c: c, clientset: clientset, refresh: make(chan bool)}
}

// GetEndPoints returns an array of Ips for the given service at that point of time.
func (k *K8r) GetEndPoints() []string {
	k.latest = make([]string, 0)
	endPoints, err := k.clientset.CoreV1().Endpoints(k.c.NameSpace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Error.Fatal("Unable to fetch Endpoints", err)
	}
	log.Info.Println("Total Endpoints from k8 cluster", len(endPoints.Items))
	for _, endPoint := range endPoints.Items {
		name := endPoint.GetObjectMeta().GetName()
		if name != k.c.Address {
			continue
		}
		// extract ips
		for _, subset := range endPoint.Subsets {
			for _, address := range subset.Addresses {
				k.latest = append(k.latest, address.IP)
			}
		}
	}
	log.Debug.Println("Resolved Ips: ", k.latest)
	return k.latest
}

// Watch triggers a bool chan on K8r to convey refresh to the caller
// Blocks for updates
func (k *K8r) Watch() {
	var err error
	k.wi, err = k.clientset.CoreV1().Endpoints(k.c.NameSpace).Watch(context.Background(), metav1.ListOptions{Watch: true})
	if err != nil {
		log.Error.Fatal("Unable to Eatch for Endpoint changes", err)
	}
	for range k.wi.ResultChan() {
		log.Info.Printf("Change Detected in k8 for %s ", k.c.Address)
		k.refresh <- true
	}
}

// Refresh tells when to refresh the data
func (k *K8r) Refresh() <-chan bool {
	return k.refresh
}

// Close stops the recolver
func (k *K8r) Close() {
	k.wi.Stop()
}

// Latest returns the latest when watcher triggers a refresh
// for k8 since the watcher doesnt update the list of endpoints.,
// Latest call will get the endpoints and update them
func (k *K8r) Latest() []string {
	k.GetEndPoints()
	return k.latest
}
