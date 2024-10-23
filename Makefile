gen:
	@protoc \
	--proto_path=app/proto \
	--go_out=app/generated/swc --go_opt=paths=source_relative \
	--go-grpc_out=app/generated/swc --go-grpc_opt=paths=source_relative \
	app/proto/swc_list.proto
