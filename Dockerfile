# Build binary
FROM golang:1.21 as builder

ARG TARGETOS
ARG TARGETARCH

WORKDIR /bin

COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go

RUN go mod download

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o sonic main.go

# Build image
FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /bin/sonic .
USER 65532:65532

ENTRYPOINT ["/bin/sonic"]
