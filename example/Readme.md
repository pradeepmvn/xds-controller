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

```
### Deploy
```
kubectl apply -f example/k8s/xds-controller.deployment.yaml -n xds-test
kubectl apply -f example/k8s/server.deployment.yaml -n xds-test
kubectl apply -f example/k8s/client.deployment.yaml -n xds-test
```

## Test
Scale up and scaledown server pods and client should get notifications on changes and new pods should start getting
the requests
