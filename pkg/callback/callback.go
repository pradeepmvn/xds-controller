package callback

import (
	"context"
	"sync"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/pradeepmvn/xds-controller/pkg/log"
	"github.com/prometheus/client_golang/prometheus"
)

// go-control-plane defines standard interface for  callback mechanism, which can be used to record and expose metrics out of the xDS requests.
// GRPC SoTW (State of The World) part of XDS server functions below.. (Rest functions are not implemented)
// OnStreamOpen, OnStreamClosed
//OnStreamRequest, OnStreamResponse

type Callbacks struct {
	//Stream counters
	ActiveStrms *prometheus.Gauge
	ReqC        *prometheus.Counter
	ResC        *prometheus.Counter
	// mux for incrementing counters
	mu sync.Mutex
}

// OnStreamOpen is called once an xDS stream is open with a stream ID and the type URL (or "" for ADS).
// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
func (cb *Callbacks) OnStreamOpen(_ context.Context, id int64, typ string) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	(*cb.ActiveStrms).Inc()
	log.Debug.Printf("Callback: Stream open for  id: %d open for type: %s", id, typ)
	return nil
}

// OnStreamClosed is called immediately prior to closing an xDS stream with a stream ID.
func (cb *Callbacks) OnStreamClosed(id int64) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	(*cb.ActiveStrms).Dec()
	log.Debug.Printf("Callback: Stream Closed for  id: %d", id)
}

// OnStreamRequest is called once a request is received on a stream.
// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
func (cb *Callbacks) OnStreamRequest(a int64, d *discovery.DiscoveryRequest) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	(*cb.ReqC).Inc()
	log.Debug.Println("Callback: Stream Request", d)
	return nil
}

// OnStreamResponse is called immediately prior to sending a response on a stream.
func (cb *Callbacks) OnStreamResponse(a int64, req *discovery.DiscoveryRequest, d *discovery.DiscoveryResponse) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	(*cb.ResC).Inc()
	log.Debug.Println("Callback: Stream Response", d)
}

// OnFetchRequest Marker Impl: No expecting Rest Client
func (cb *Callbacks) OnFetchRequest(ctx context.Context, req *discovery.DiscoveryRequest) error {
	return nil
}

// OnFetchResponse Marker Impl: No expecting Rest Client
func (cb *Callbacks) OnFetchResponse(req *discovery.DiscoveryRequest, resp *discovery.DiscoveryResponse) {
}

// OnDeltaStreamOpen is called once an incremental xDS stream is open with a stream ID and the type URL (or "" for ADS).
// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
func (cb *Callbacks) OnDeltaStreamOpen(context.Context, int64, string) error {
	return nil
}

// OnDeltaStreamClosed is called immediately prior to closing an xDS stream with a stream ID.
func (cb *Callbacks) OnDeltaStreamClosed(int64) {

}

// OnStreamDeltaRequest is called once a request is received on a stream.
// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
func (cb *Callbacks) OnStreamDeltaRequest(int64, *discovery.DeltaDiscoveryRequest) error {
	return nil
}

// OnStreamDelatResponse is called immediately prior to sending a response on a stream.
func (cb *Callbacks) OnStreamDeltaResponse(int64, *discovery.DeltaDiscoveryRequest, *discovery.DeltaDiscoveryResponse) {
}
