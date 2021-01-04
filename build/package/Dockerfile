FROM golang:1.14.2-buster@sha256:09b04534495af5148e4cc67c8ac55408307c2d7b9e6ce70f6e05f7f02e427f68 AS tools
ENV GO111MODULE on
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.27.0 && \
    go get -u github.com/swaggo/swag/cmd/swag@v1.6.5 && \
    go get github.com/vektra/mockery/.../@v1.0.0

FROM golang:1.14.2-buster@sha256:09b04534495af5148e4cc67c8ac55408307c2d7b9e6ce70f6e05f7f02e427f68 AS base
COPY --from=tools /go/bin/golangci-lint /go/bin/swag /go/bin/mockery /go/bin/
WORKDIR /opt/app

FROM base AS dependency-cache
COPY go.mod go.sum ./
RUN go mod download
#   go mod verify
ARG VERSION
ENV VERSION=$VERSION

FROM dependency-cache AS development
CMD ["sleep infinity"]

FROM dependency-cache AS build
COPY . .
RUN make shell-go-build

FROM gcr.io/distroless/base-debian10@sha256:e5768afa5429b85ac75de67efc98a4bf53e4ef0f7388667fb34c89d481d82b00 AS production
ARG NAME
USER 2000:2000
WORKDIR /opt/app
COPY assets "/opt/app/$NAME/assets"
COPY --from=build /opt/app/bin/$NAME ./$NAME
CMD ["/opt/app/$NAME"]
