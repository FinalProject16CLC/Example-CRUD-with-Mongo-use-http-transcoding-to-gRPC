# Example project using gRPC and http transcoding

## Make sure all the dependencies is in sync

`dep status` && `dep ensure`

## Generate gRPC stub

- Generating client and server code

  protoc -I/usr/local/include -I. \
  -Ivendor \
  -Igoogleapis \
  --go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
  protos/entity.proto

- Generate reverse-proxy for your RESTful API:

  protoc -I/usr/local/include -I. \
   -Igoogleapis --include_imports --include_source_info \
   --descriptor_set_out=protos/proto.pb protos/entity.proto

or you can use make file:
  `make`

## Start project

  `docker-compose up`

## Example API Calls

## List entities

  `curl -X GET 'http://localhost:8080/entities'`

## Create entity

  `curl -X POST 'http://localhost:8080/entities' -d '{"name":"Phuc qua dep trai","description":"Kha la banh","url":"phucdeptrai.com.vn"}'`

## Read entity

  `curl -X GET "http://localhost:8080/entities/5d11e96b9dadaf6eef8599be"`

## Update entity

  `curl -X PUT 'http://localhost:8080/entities' -d '{"id":"5d11e8ee9dadaf6eef8599b9","name":"Phuc qua dep trai","description":"Kha la banh","url":"phucdeptrai.com.vn"}'`

## Delete entity

  `curl -X DELETE "http://localhost:8080/entities/5d11e8ee9dadaf6eef8599b9"`

## USE KONG AS API GATEWAY

- If you want to use Kong (https://konghq.com/kong/) as API gateway. You can checkout to banch kong-api-gw

  `git checkout kong-api-gw`

## USE ENVOY AS API GATEWAY

- If you want to use Kong (https://www.envoyproxy.io/) as API gateway. You can checkout to banch envoy

  `git checkout envoy`
