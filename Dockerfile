FROM golang:1.15.6 as build

WORKDIR /app
COPY . .
RUN go build -o main ./cmd

FROM scratch
EXPOSE 8080
COPY --from=build /app/main /
ENTRYPOINT [ "/app/main" ]
