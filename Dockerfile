FROM golang:1.15.6 as build
WORKDIR /app
COPY . .
RUN go build -o main ./cmd

FROM scratch
# grpc server and prometheus
EXPOSE 8080 8082
COPY --from=build /app/main /
ENTRYPOINT [ "/app/main" ]
