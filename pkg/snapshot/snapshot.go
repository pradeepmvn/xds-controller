package snapshot

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/pradeepmvn/xds-controller/pkg/config"
	"github.com/pradeepmvn/xds-controller/pkg/log"
	"github.com/pradeepmvn/xds-controller/pkg/resolver"
	"github.com/pradeepmvn/xds-controller/pkg/resolver/dns"
	"github.com/pradeepmvn/xds-controller/pkg/resolver/k8"
	"google.golang.org/protobuf/types/known/anypb"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

/*
Resources:
Clusters: a name and eds type
listener:
   |- RouteConfigName
   |- Address
   |- Filterchains
endpoints:
   |- clustername
   |- list of ips
routeconfig
   |- clustername
   |- virtualhost
*/

type snapshot struct {
	clientset *kubernetes.Clientset
	cfg       *config.ControllerConfig
	cache     cache.SnapshotCache
	lstate    map[string]resolver.Resolver
	m         sync.Mutex
}
type resources struct {
	// snapshot resoources...all part of 1 snapshot in cache
	endpoints []types.Resource
	clusters  []types.Resource
	routes    []types.Resource
	listeners []types.Resource
}

type Snapshot interface {
	StartRefresher()
	Close()
}

const (
	RouteTypePath   = "path"
	RouteTypePrefix = "prefix"
	listenerSuffix  = "-listener"
	vhostSuffix     = "-vhost"
	routeSuffix     = "-route"
	clusterSuffix   = "-cluster"
)

// NewSnapshot creates an instan of Snapshot that refreshes itself
func NewSnapshot(cfg *config.ControllerConfig, cache cache.SnapshotCache) Snapshot {
	var cls *kubernetes.Clientset
	// check if there is a k8 discovery needed
	for _, clusterConfig := range cfg.Clusters {
		if strings.EqualFold(clusterConfig.ResolverType, k8.Type) {
			// create k8 config
			config, err := rest.InClusterConfig()
			if err != nil {
				log.Error.Panic("Found a K8 resolver type., but unable to  get K8 CLuster Config", err)
			}
			// creates the client set
			cls, err = kubernetes.NewForConfig(config)
			if err != nil {
				log.Error.Panic(err)
			}
			break
		}
	}
	return &snapshot{
		cfg:       cfg,
		cache:     cache,
		clientset: cls,
		m:         sync.Mutex{},
	}
}

// SnapshotRefresher triggers a new shapshot for configured interval time
// Calling this function is an infinite loop that keeps refreshing till the main process is stopped.
func (sn *snapshot) StartRefresher() {
	// prepare data
	sn.prepare()
	// generate snapshot
	sn.generate()
}

// prepares localstate with resolvers and refreshers
func (sn *snapshot) prepare() {
	sn.lstate = make(map[string]resolver.Resolver)
	for _, cl := range sn.cfg.Clusters {
		var r resolver.Resolver
		if strings.EqualFold(cl.ResolverType, dns.Type) {
			r = dns.New(cl)
		} else if strings.EqualFold(cl.ResolverType, k8.Type) {
			r = k8.New(cl, sn.clientset)
		} else {
			log.Error.Fatal("Unknow type of resolver.. Please fix it..")
		}
		sn.lstate[cl.Name] = r
		go func(r resolver.Resolver, sn *snapshot) {
			for range r.Refresh() {
				// got a refresh trigger regenerate
				sn.generate()
			}
		}(r, sn)
		go r.Watch()
	}
	log.Info.Println("Done Preparing structures for resolvers")
}

// xDS Envoy configuration includes cluster, endpoints, route
// generate updates the snapshot with new resource def
func (sn *snapshot) generate() {
	sn.m.Lock()
	defer sn.m.Unlock()
	lr := &resources{
		endpoints: make([]types.Resource, 0),
		clusters:  make([]types.Resource, 0),
		routes:    make([]types.Resource, 0),
		listeners: make([]types.Resource, 0),
	}
	for _, cl := range sn.cfg.Clusters {
		r := sn.lstate[cl.Name]
		clusterName := cl.Name + clusterSuffix
		ips := r.GetEndPoints()
		lr.endpoints = append(lr.endpoints, createEds(cl, ips, clusterName)) // endpoints
		lr.clusters = append(lr.clusters, createCluster(cl, clusterName))    //cds
		lr.routes = append(lr.routes, createRoute(cl, clusterName))          // rds
		lr.listeners = append(lr.listeners, createListener(cl))              // lds
	}
	// set snapshot
	version := uuid.New()
	xdsSn, err := cache.NewSnapshot(
		version.String(),
		map[resource.Type][]types.Resource{
			resource.ClusterType:  lr.clusters,  //cds,
			resource.RouteType:    lr.routes,    //rds,
			resource.ListenerType: lr.listeners, //lds,
			resource.EndpointType: lr.endpoints, // endpoints
		},
	)
	if err != nil {
		log.Error.Panic("Unable To Create New Snapshot.. Panicing !!", err)
	}
	if err := xdsSn.Consistent(); err != nil {
		log.Error.Panic("Snapshot inconsistent.. Panicing !!", err)
	}
	if err := sn.cache.SetSnapshot(context.Background(), sn.cfg.NodeId, xdsSn); err != nil {
		log.Error.Println("Failed to Set Snapshot to cache. ", xdsSn)
		log.Error.Panic("Faile with error", err)
	}
	log.Debug.Printf("Snapshot added to cache %s", xdsSn)
}

// Close resolver and related components
func (sn *snapshot) Close() {
	sn.m.Lock()
	defer sn.m.Unlock()
	for _, r := range sn.lstate {
		r.Close()
	}
}

// MarshalAny is a convenience function to marshal protobuf messages into any
// protos. It will panic if the marshaling fails.
func MarshalAny(m proto.Message) *anypb.Any {
	a, err := ptypes.MarshalAny(m)
	if err != nil {
		panic(fmt.Sprintf("ptypes.MarshalAny(%+v) failed: %v", m, err))
	}
	return a
}
