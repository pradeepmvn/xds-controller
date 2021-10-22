package snapshot

import (
	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/pradeepmvn/xds-controller/pkg/config"
	"github.com/pradeepmvn/xds-controller/pkg/log"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Cluster Configuration from xDS control plane.EDS type.
func createCluster(cl *config.Cluster, clusterName string) types.Resource {
	log.Debug.Printf("Creating CLUSTER %s ", clusterName)
	return types.Resource(&cluster.Cluster{
		Name: clusterName,
		CircuitBreakers: &cluster.CircuitBreakers{
			Thresholds: []*cluster.CircuitBreakers_Thresholds{
				{
					Priority:    core.RoutingPriority_DEFAULT,
					MaxRequests: &wrapperspb.UInt32Value{Value: cl.MaxRequests},
				},
			},
		},
		LbPolicy:             cluster.Cluster_LbPolicy(cluster.Cluster_LbPolicy_value[cl.LbPolicy]),
		ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_EDS},
		EdsClusterConfig: &cluster.Cluster_EdsClusterConfig{
			EdsConfig: &core.ConfigSource{
				ConfigSourceSpecifier: &core.ConfigSource_Ads{},
			},
		},
	})
}
