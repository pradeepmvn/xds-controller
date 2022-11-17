# Example usage of xds-controller with a server and client

gRPC client: time bounded caller that calls name service, No ports and no services are exposed. All it does it make a grpc call to server and log it.
gRPC service: Name service that returns random name for every request


## Generate proto
```
protoc -I=example/proto --go_out=example/proto example/proto/person.proto

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    example/proto/person.proto
```

## A Kubernates deployment of microservices
- client: polls for names continuously(every sec) for 3 minutes
- server: responds with random names

## Run
To run example on a local cluster.,

### Build
```bash
docker build -t xds.example/xds-controller .
docker build -t xds.example/client -f example/Dockerfile --build-arg TYPE=client example
docker build -t xds.example/server -f example/Dockerfile --build-arg TYPE=server example
docker pull envoyproxy/envoy:v1.24.0
```

### Deployments
#### Pattern 1: CLient server communication using xds:look aside load balancing
```
kubectl create namespace xds-test
kubectl apply -f example/k8s/xds-controller.deployment.yaml -n xds-test
kubectl apply -f example/k8s/server.deployment.yaml -n xds-test
kubectl apply -f example/k8s/client.deployment.yaml -n xds-test
kubectl apply -f example/k8s/envoy.deployment.yaml -n xds-test
```
#### Pattern 2: Client with 2 servers communication using xds:look aside load balancing
```
kubectl create namespace xds-test
kubectl apply -f example/k8s/xds-controller.deployment.yaml -n xds-test
kubectl apply -f example/k8s/server.deployment.yaml -n xds-test
kubectl apply -f example/k8s/server.b.deployment.yaml -n xds-test
kubectl apply -f example/k8s/client.deployment.yaml -n xds-test
```

#### Pattern 3: External Client proxied via Envoy L7 Load balancer (No xds) using static dns
```
kubectl create namespace xds-test
kubectl apply -f example/k8s/server.deployment.yaml -n xds-test
kubectl apply -f example/k8s/envoy.deployment.yaml -n xds-test
```
#### Pattern 4: External Client proxied via Envoy L7 Load balancer. Evoy using xDS stream to discover server
```
kubectl create namespace xds-test
kubectl apply -f example/k8s/xds-controller.deployment.yaml -n xds-test
kubectl apply -f example/k8s/server.deployment.yaml -n xds-test
kubectl apply -f example/k8s/envoy.deployment.yaml -n xds-test
```


### Metrics
xds-controller exposes metrics via /metrics endpoint by default on 8082 port. Portforwarding can be used to hit localhost and see the metric details

```
// get pods
kubectl get pods -n xds-test
kubectl port-forward {pod-name} 8082:8082 -n xds-test
curl localhost:8082/metrics
 ```

## Test
Scale up and scaledown server pods and client should get notifications on changes and new pods should start getting
the requests
