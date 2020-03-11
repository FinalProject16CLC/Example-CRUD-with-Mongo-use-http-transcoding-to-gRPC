pb:
	protoc -I/usr/local/include -I. \
		-Ivendor \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api,plugins=grpc:. \
		protos/entity/entity.proto
	protoc -I/usr/local/include -I. \
		-Ivendor \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		protos/entity/entity.proto
	protoc -I/usr/local/include -I. \
		-Ivendor \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api,plugins=grpc:. \
		protos/user/user.proto
	protoc -I/usr/local/include -I. \
		-Ivendor \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		protos/user/user.proto