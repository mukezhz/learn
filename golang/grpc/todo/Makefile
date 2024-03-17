SRC_DIR := proto
DST_DIR := pb


.PHONY: gen

gen:
	protoc -I${SRC_DIR} \
		--go_out=${DST_DIR} --go_opt=paths=source_relative \
		--go-grpc_out=${DST_DIR} --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=${DST_DIR} \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		proto/*.proto
