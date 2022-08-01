FROM golang:1.18.5-buster AS build-env
ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
COPY ./gov-gr-exporter /build
WORKDIR /build
RUN go build -o gov_gr_exporter

FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=build-env /build/gov_gr_exporter /app
ENTRYPOINT ["./gov_gr_exporter"]