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
RUN echo "[default]\naws_access_key_id = dummy\naws_secret_access_key = dummy\nregion = us-east-2\n" > credentials

FROM scratch as server
COPY --from=builder /go/src/github.com/jyotishp/go-orders/server /bin/server
COPY --from=builder /go/src/github.com/jyotishp/go-orders/swagger-ui /opt/swagger-ui
COPY --from=builder /go/src/github.com/jyotishp/go-orders/credentials /root/.aws/credentials
WORKDIR /opt
CMD ["/bin/server"]

