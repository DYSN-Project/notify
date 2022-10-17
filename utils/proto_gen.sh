protoc -I ../internal/transport/grpc/proto/ \
 notify.proto \
 --go-grpc_out=../internal/transport --go_out=../internal/transport