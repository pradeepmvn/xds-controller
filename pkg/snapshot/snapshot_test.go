package snapshot_test

import (
	"testing"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/pradeepmvn/xds-controller/pkg/config"
	"github.com/pradeepmvn/xds-controller/pkg/log"
	"github.com/pradeepmvn/xds-controller/pkg/snapshot"
)

func TestSnapshot(t *testing.T) {
	con := &config.ControllerConfig{
		NodeId: "3243232",
		Clusters: []*config.Cluster{{
			Name:         "FakeK8",
			ResolverType: "dns",
			Address:      "google.com",
		},
		},
	}

	// setup logger
	log.NewLogger(true)
	// Create a cache
	var l log.CLog
	cache := cache.NewSnapshotCache(false, cache.IDHash{}, l)
	sn := snapshot.NewSnapshot(con, cache)
	sn.StartRefresher()
}
