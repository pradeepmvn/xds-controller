FROM golang:1.15.6 as build
ARG TYPE=client
WORKDIR /src
COPY . .
ENV CGO_ENABLED=0
RUN go build -o main ./${TYPE}

FROM scratch
EXPOSE 8080
COPY --from=build /src/main /main
ENTRYPOINT ["/main"]
