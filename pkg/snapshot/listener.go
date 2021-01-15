package snapshot

import (
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"

	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/golang/protobuf/ptypes"
	"github.com/pradeepmvn/xds-controller/pkg/config"
	"github.com/pradeepmvn/xds-controller/pkg/log"
)

// Listener resourcetype in envoy configuration.
// Only one listner for multiple cluster configurations. Listener has a route config
func createListener(c *config.Cluster) types.Resource {
	cm := &hcm.HttpConnectionManager{
		CodecType: hcm.HttpConnectionManager_AUTO,
		RouteSpecifier: &hcm.HttpConnectionManager_Rds{
			Rds: &hcm.Rds{
				RouteConfigName: c.Name + routeSuffix,
				ConfigSource: &core.ConfigSource{
					ConfigSourceSpecifier: &core.ConfigSource_Ads{
						Ads: &core.AggregatedConfigSource{},
					},
				},
			},
		},
	}
	pbst, err := ptypes.MarshalAny(cm)
	if err != nil {
		log.Error.Panic("Unable to Marshal Connection Manager ", err)
	}

	return types.Resource(&listener.Listener{
		Name: c.Name + listenerSuffix,
		ApiListener: &listener.ApiListener{
			ApiListener: pbst,
		},
		Address: &core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.SocketAddress_TCP,
					Address:  "0.0.0.0",
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: 10000,
					},
				},
			},
		},
		FilterChains: []*listener.FilterChain{{
			Filters: []*listener.Filter{{
				Name: wellknown.HTTPConnectionManager,
				ConfigType: &listener.Filter_TypedConfig{
					TypedConfig: pbst,
				},
			}},
		}},
	})
}
