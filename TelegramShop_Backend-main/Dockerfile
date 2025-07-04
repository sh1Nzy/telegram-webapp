ARG GO_BASE_VERSION=1.24.3
ARG GO_BASE_IMAGE=alpine

FROM golang:${GO_BASE_VERSION}-${GO_BASE_IMAGE} AS build
ARG SERVICE_NAME=backend

WORKDIR /src

COPY . .

# Build
RUN go build -o ./${SERVICE_NAME} ./cmd/main.go

# Run service
FROM alpine as main
RUN adduser -D ghost && mkdir -p /etc/sudoers.d \
    && echo "ghost ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/ghost \
    && chmod 0440 /etc/sudoers.d/ghost
ARG SERVICE_NAME=backend

COPY --from=build /src/${SERVICE_NAME} ./app/${SERVICE_NAME}

RUN apk --no-cache add curl
RUN set -ex && apk --no-cache add sudo

USER ghost
WORKDIR /app

CMD ["./backend"]