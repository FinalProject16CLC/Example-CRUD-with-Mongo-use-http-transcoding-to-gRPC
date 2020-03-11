pb:
	protoc -I/usr/local/include -I. \
		-Ivendor \
		-Igoogleapis \
		--go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
		protos/entity.proto
		
	protoc -I/usr/local/include -I. \
   -Igoogleapis --include_imports --include_source_info \
   --descriptor_set_out=protos/proto.pb protos/entity.proto