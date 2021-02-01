package main

import (
	"context"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	controlplane "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/pradeepmvn/xds-controller/pkg/callback"
	"github.com/pradeepmvn/xds-controller/pkg/config"
	"github.com/pradeepmvn/xds-controller/pkg/log"
	"github.com/pradeepmvn/xds-controller/pkg/server"
	"github.com/pradeepmvn/xds-controller/pkg/snapshot"
)

// Main Entry point for xds-controller.
// Starts a grpc server, creates snapshot refresher in background and serves it.
func main() {
	// load config
	cfg := config.LoadConfig()

	// setup logger
	log.NewLogger(cfg.LogDebug)

	// Create a cache
	var l log.CLog
	cache := cache.NewSnapshotCache(false, cache.IDHash{}, l)

	// Start the snapshot refresher..
	sn := snapshot.NewSnapshot(cfg, cache)
	defer sn.Close()
	go sn.StartRefresher()

	// Run the xDS server
	ctx := context.Background()
	cb := &callback.Callbacks{}
	srv := controlplane.NewServer(ctx, cache, cb)
	server.RunControlPlaneServer(ctx, srv, cfg)
}