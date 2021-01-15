// Snapshot from example https://github.com/envoyproxy/go-control-plane/blob/master/internal/example/resource.go

package snapshot

import (
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/pradeepmvn/xds-controller/pkg/config"
	"github.com/pradeepmvn/xds-controller/pkg/log"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// End point discovery service for resource type.
// Creats a list of LB ips and add them to Endpoint
func createEds(c *config.Cluster, ips []string, clusterName string) types.Resource {
	var lbs []*endpoint.LbEndpoint
	// create lb for available service ips
	for _, ip := range ips {
		log.Debug.Printf("Creating ENDPOINT %s ", ip)
		lbs = append(lbs, &endpoint.LbEndpoint{
			HostIdentifier: &endpoint.LbEndpoint_Endpoint{
				Endpoint: &endpoint.Endpoint{
					Address: &core.Address{Address: &core.Address_SocketAddress{
						SocketAddress: &core.SocketAddress{
							Address:  ip,
							Protocol: core.SocketAddress_TCP,
							PortSpecifier: &core.SocketAddress_PortValue{
								PortValue: uint32(c.Port),
							},
						},
					}},
				}},
			// TODO
			HealthStatus: core.HealthStatus_HEALTHY,
		})
	}
	// create an eds cluster with lbs
	return types.Resource(&endpoint.ClusterLoadAssignment{
		ClusterName: clusterName,
		Endpoints: []*endpoint.LocalityLbEndpoints{
			{
				Locality: &core.Locality{
					Region: "region",
					Zone:   "zone",
				},
				Priority:            0,
				LoadBalancingWeight: &wrapperspb.UInt32Value{Value: uint32(1000)},
				LbEndpoints:         lbs,
			},
		},
	})
}
