proto:
	protoc \
	-I pkg \
	-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
	--go_out=plugins=grpc:. \
	--grpc-gateway_out=logtostderr=true:. \
	--swagger_out=logtostderr=true:swagger-ui \
	--proto_path pkg/proto analysis.proto restaurant.proto customer.proto orders.proto auth.proto

run:
	go run cmd/server/main.go

process-data:
	go run cmd/dataprocessing/main.go
