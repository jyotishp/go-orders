FROM golang:1.14 as builder

ARG REPO_HOST=github.com
ARG REPO_NAMESPACE=jyotishp/go-orders

ENV CGO_ENABLED 0
ENV REPO_PATH /go/src/${REPO_HOST}/${REPO_NAMESPACE}

COPY . ${REPO_PATH}
WORKDIR ${REPO_PATH}

RUN make install-proto
RUN make build

FROM scratch as server
COPY --from=builder /go/src/github.com/jyotishp/go-orders/server /bin/server
COPY swagger-ui .
CMD ["/bin/server"]
