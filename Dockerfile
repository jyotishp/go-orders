FROM golang:1.14 as builder

ARG REPO_HOST=github.com
ARG REPO_NAMESPACE=jyotishp/go-orders

ENV CGO_ENABLED 0
ENV REPO_PATH /go/src/${REPO_HOST}/${REPO_NAMESPACE}

RUN apt update && apt install -y sudo jq

COPY . ${REPO_PATH}
WORKDIR ${REPO_PATH}

RUN make install-proto
RUN make build
RUN make fix-swagger

FROM scratch as server
COPY --from=builder /go/src/github.com/jyotishp/go-orders/server /bin/server
COPY --from=builder /go/src/github.com/jyotishp/go-orders/swagger-ui /opt/swagger-ui
WORKDIR /opt
CMD ["/bin/server"]

