# STEP 1 build optimized executable binary

FROM golang:1.18 AS builder

ARG SERVICE_RELATIVE_PATH
ARG SERVICE_VERSION

COPY . /src

WORKDIR /src/${SERVICE_RELATIVE_PATH}

RUN make build VERSION=${SERVICE_VERSION}

# STEP 2 build a small image

FROM debian:stable-slim

ARG SERVICE_RELATIVE_PATH
ARG GRPC_PORT
ARG HTTP_PORT

RUN apt update && \
    apt install -y --no-install-recommends ca-certificates netbase && \
    rm -rf /var/lib/apt/lists/ && \
    apt autoremove -y && \
    apt autoclean -y

COPY --from=builder /src/${SERVICE_RELATIVE_PATH}/bin /app
COPY --from=builder /src/${SERVICE_RELATIVE_PATH}/configs/ /app/conf/

WORKDIR /app

EXPOSE ${GRPC_PORT} ${HTTP_PORT}
VOLUME /data/conf

#USER nonroot:nonroot

CMD ["./server", "-conf", "/app/conf"]