module github.com/pradeepmvn/xds-controller

go 1.15

require (
	github.com/envoyproxy/go-control-plane v0.9.9-0.20201210154907-fd9021fe5dad
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.4
	github.com/prometheus/client_golang v1.9.0
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
	k8s.io/api v0.20.1
	k8s.io/apimachinery v0.20.1
	k8s.io/client-go v0.20.1
)

replace (
	github.com/nats-io/nats-server/v2 =>github.com/nats-io/nats-server/v2 v2.1.9
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
	github.com/aws/aws-sdk-go => github.com/aws/aws-sdk-go v1.36.31
)
