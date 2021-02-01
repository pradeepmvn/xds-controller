gRPC project only supports traces . No metrics are supported by default
opencencus and  opentracing are merged into open telementry

golang: https://github.com/open-telemetry/opentelemetry-go
This repo does not include exporter to stackdriver by default.

To export metrics to stackdriver.. this is provided by Google
https://github.com/GoogleCloudPlatform/opentelemetry-operations-go


MEtrics:
Context: span and correlation id
MEter: used to record a measurement
Raw measurement: 
    Measure
    Measuement
Metric:
    kind:counter, measure, observer
    label: key/value pair, metadata
Aggregation
Time

Semantic conventions

Auto instrumentation

2 are different
collector: Pipelines that joijnes recivers to exporters
exporter

https://help.sumologic.com/Traces/Getting_Started_with_Transaction_Tracing/Instrument_your_application_with_OpenTelemetry/Go_OpenTelemetry_auto-instrumentation