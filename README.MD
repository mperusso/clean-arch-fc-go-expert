migrate  -path=sql/migrations --database="mysql://root:root@tcp(localhost:3306)/orders" up



cd cmd/ordersystem
go run .





protoc \
--proto_path=internal/infra/grpc/protofiles \
--go_out=internal/infra/grpc/pb \
--go_opt=paths=source_relative \
--go-grpc_out=internal/infra/grpc/pb \
--go-grpc_opt=paths=source_relative \
internal/infra/grpc/protofiles/order.proto
