package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	controlplane "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/pradeepmvn/xds-controller/pkg/callback"
	"github.com/pradeepmvn/xds-controller/pkg/config"
	"github.com/pradeepmvn/xds-controller/pkg/log"
	"github.com/pradeepmvn/xds-controller/pkg/server"
	"github.com/pradeepmvn/xds-controller/pkg/snapshot"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Prometheus Metrics plugged into callbac
	activeStreams = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "xds_controller",
		Subsystem: "grpc",
		Name:      "active_streams",
		Help:      "Active grpc streams to xds-controller",
	})

	streamReq = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "xds_controller",
		Subsystem: "grpc",
		Name:      "stream_requests",
		Help:      "No.of requests via grpc streams to xds-controller",
	})

	streamResp = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "xds_controller",
		Subsystem: "grpc",
		Name:      "stream_responses",
		Help:      "No.of Reponses sent to clients by  xds-controller ",
	})
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
	// ADS flag forces a delay in responding to streaming requests until all resources are explicitly named in the request.
	// Making it to false to make it convinient for one control plane to send partial responses to clients.
	cache := cache.NewSnapshotCache(false, cache.IDHash{}, l)

	// Start the snapshot refresher..
	sn := snapshot.NewSnapshot(cfg, cache)
	defer sn.Close()
	go sn.StartRefresher()

	// Register Prometheus metrics handler.
	http.Handle("/metrics", promhttp.Handler())
	pm := fmt.Sprintf(":%d", cfg.PrometheusPort)
	log.Info.Println("Starting Prometheus Metrics Agent at ", pm)
	promServer := &http.Server{
		Addr:              pm,
		ReadHeaderTimeout: 3 * time.Second,
	}
	go promServer.ListenAndServe()
	//-----------------------

	// Configure  the xDS server
	ctx := context.Background()
	cb := &callback.Callbacks{ActiveStrms: &activeStreams, ReqC: &streamReq, ResC: &streamResp}
	srv := controlplane.NewServer(ctx, cache, cb)
	// Run xDS server
	server.RunControlPlaneServer(ctx, srv, cfg)
}
