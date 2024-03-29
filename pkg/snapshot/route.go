package snapshot

import (
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/pradeepmvn/xds-controller/pkg/config"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Route Confiuration pointintg to cluster.Contains, VH,
func createRoute(c *config.Cluster, clusterName string) types.Resource {
	routeAction := &route.Route_Route{
		Route: &route.RouteAction{
			ClusterSpecifier: &route.RouteAction_Cluster{
				Cluster: clusterName,
			},
		},
	}
	if c.Retry.Enabled {
		routeAction.Route.RetryPolicy = &route.RetryPolicy{
			RetryOn:    c.Retry.RetryOn,
			NumRetries: &wrapperspb.UInt32Value{Value: c.Retry.NumRetries},
			RetryBackOff: &route.RetryPolicy_RetryBackOff{
				BaseInterval: &durationpb.Duration{Seconds: int64(c.Retry.BackoffIntervalInSec)},
				MaxInterval:  &durationpb.Duration{Seconds: int64(c.Retry.BackoffMaxIntervalInSec)},
			},
		}
	}
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
						Action: routeAction,
					},
				},
			},
		},
	})
}
