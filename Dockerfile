FROM golang:1.19.2 as build
WORKDIR /app/src
COPY . .
ENV CGO_ENABLED=0
RUN go build -o /app/main ./cmd

FROM scratch
# grpc server and prometheus
EXPOSE 8080 8082
COPY --from=build /app/main /main
COPY --from=build /app/src/config/config.yaml /config/config.yaml
WORKDIR /
ENTRYPOINT [ "/main" ]
