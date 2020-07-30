all: proto run

compile_proto: build
	protoc \
	-I pkg \
	-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
	--go_out=plugins=grpc:. \
	--grpc-gateway_out=logtostderr=true:. \
	--swagger_out=logtostderr=true,allow_merge=true,merge_file_name=app:build \
	--proto_path pkg/proto auth.proto customer.proto restaurant.proto orders.proto analysis.proto utils.proto

build:
	mkdir -p build

clean:
	rm -rf build

fix_swagger:
	cat build/app.swagger.json | jq -c | sed 's/"title":".*\.proto"/"title":"Store"/g' > build/swagger.json && \
    mv build/swagger.json swagger-ui/app.swagger.json && \
    rm -f build/swagger.json

proto: compile_proto fix_swagger

run:
	go run cmd/server/main.go

process-data:
	go run cmd/dataprocessing/main.go


