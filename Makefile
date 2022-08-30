proto-build:
	protoc -Iapi/v1 -Ipkg/third_party --go_out=pkg/gen/proto/v1	\
		--go_opt paths=source_relative	\
		--go-grpc_out=pkg/gen/proto/v1	\
		--go-grpc_opt paths=source_relative	\
		--grpc-gateway_out=pkg/gen/proto/v1 \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
        --openapiv2_out=api/v1 \
        --openapiv2_opt logtostderr=true \
        --openapiv2_opt generate_unbound_methods=true \
		company.proto


run:
	docker build --tag job .
	docker run -d --name emil -p 9997:9997 -p 8888:8888 job

test-%:
	go run cmd/gRPC-client/main.go --inn=$*


