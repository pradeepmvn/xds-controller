package snapshot

import (
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/pradeepmvn/xds-controller/pkg/config"
)

// Route Confiuration pointintg to cluster.Contains, VH,
func createRoute(c *config.Cluster, clusterName string) types.Resource {
	return types.Resource(&route.RouteConfiguration{
		Name: c.Name + routeSuffix,
		VirtualHosts: []*route.VirtualHost{
			{
				Name: c.Name + vhostSuffix,
				Domains: []string{
					c.Name + listenerSuffix,
				},
				Routes: []*route.Route{
					{
						Match: &route.RouteMatch{
							PathSpecifier: &route.RouteMatch_Prefix{
								// Todo
								Prefix: "",
							},
						},
						Action: &route.Route_Route{
							Route: &route.RouteAction{
								ClusterSpecifier: &route.RouteAction_Cluster{
									Cluster: clusterName,
								},
							},
						},
					},
				},
			},
		},
	})
}
