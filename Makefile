all: proto run

install-proto:
	sudo apt update && sudo apt install -y protobuf-compiler
	go get -u -v github.com/golang/protobuf/protoc-gen-go

compile-proto: pre-build proto-dependencies
	protoc \
	-I pkg \
	-I ./vendor/github.com/grpc-ecosystem/grpc-gateway \
	--go_out=plugins=grpc:. \
	--grpc-gateway_out=logtostderr=true:. \
	--swagger_out=logtostderr=true,allow_merge=true,merge_file_name=app:build \
	--proto_path pkg/proto auth.proto customer.proto restaurant.proto orders.proto analysis.proto utils.proto

pre-build:
	mkdir -p build

proto-dependencies:
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

build: pre-build compile-proto
	go mod vendor
	go build -o server cmd/server/main.go

clean:
	rm -rf build
	rm -f swagger-ui/app.swagger.json

tests:
	go test ./... -v -coverprofile coverage.out
	go tool cover -html=c.out -o coverage.html

fix-swagger:
	cat build/app.swagger.json | jq -c | sed 's/"title":".*\.proto"/"title":"Store"/g' > build/swagger.json && \
    mv build/swagger.json swagger-ui/app.swagger.json && \
    rm -f build/swagger.json

proto: compile-proto fix-swagger

run:
	go run cmd/server/main.go

process-data:
	go run cmd/dataprocessing/main.go


